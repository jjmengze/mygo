package http_server

import (
	"github.com/jjmengze/mygo/pkg/telemetry/instrumentation"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"io"
	"net/http"
)

var _ http.Handler = &handler{}

type handler struct {
	config  *config
	handler http.Handler
}

func NewHttpHandler(h http.Handler) *handler {
	return &handler{
		config:  newConfig(),
		handler: h,
	}
}

func (h *handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	for _, filterFunc := range h.config.filterFunc {
		if !filterFunc(request) {
			h.handler.ServeHTTP(writer, request)
			return
		}
	}
	ctx := otel.GetTextMapPropagator().Extract(request.Context(), propagation.HeaderCarrier(request.Header))
	ctx, sp := h.config.tracerProvider.Tracer("").Start(ctx, request.RequestURI)
	defer sp.End()

	w := instrumentation.NewRecordingResponseWriter(writer)
	h.handler.ServeHTTP(w, request.WithContext(ctx))

	attributes := []attribute.KeyValue{}
	if w.Error != nil && w.Error != io.EOF {
		attributes = append(attributes, attribute.Key("http.read_error").String(w.Error.Error()))
	}
	attributes = append(attributes, attribute.Key("http.wrote_bytes").Int64(w.Written))
	sp.SetAttributes(attributes...)
}
