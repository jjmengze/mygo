package gorilla

import (
	"fmt"
	"github.com/gorilla/mux"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	oteltrace "go.opentelemetry.io/otel/trace"
	"mygo/pkg/telemetry/instrumentation"
	"net/http"
)

var tracerName string = "gorilla"

type GorillaTrace struct {
	service     string
	propagators propagation.TextMapPropagator
	handler     http.Handler
}

func Middleware() mux.MiddlewareFunc {
	return func(handler http.Handler) http.Handler {
		return &GorillaTrace{
			service:     "",
			propagators: otel.GetTextMapPropagator(),
			handler:     handler,
		}
	}
}

func (gt *GorillaTrace) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := gt.propagators.Extract(r.Context(), propagation.HeaderCarrier(r.Header))
	spanName := ""
	route := mux.CurrentRoute(r)
	if route != nil {
		var err error
		spanName, err = route.GetPathTemplate()
		if err != nil {
			spanName, _ = route.GetPathRegexp()
		}
	}
	routeStr := spanName
	if spanName == "" {
		spanName = fmt.Sprintf("HTTP %s route not found", r.Method)
	}
	opts := []oteltrace.SpanStartOption{
		oteltrace.WithAttributes(semconv.NetAttributesFromHTTPRequest("tcp", r)...),
		oteltrace.WithAttributes(semconv.EndUserAttributesFromHTTPRequest(r)...),
		oteltrace.WithAttributes(semconv.HTTPServerAttributesFromHTTPRequest(gt.service, routeStr, r)...),
		oteltrace.WithSpanKind(oteltrace.SpanKindServer),
	}

	spanCtx, span := otel.Tracer(tracerName).Start(ctx, spanName, opts...)
	defer span.End()

	resp := r.WithContext(spanCtx)
	writer := instrumentation.NewRecordingResponseWriter(w)
	gt.handler.ServeHTTP(writer, resp)
	attrs := semconv.HTTPAttributesFromHTTPStatusCode(writer.StatusCode)
	spanStatus, spanMessage := semconv.SpanStatusFromHTTPStatusCode(writer.StatusCode)
	span.SetAttributes(attrs...)
	span.SetStatus(spanStatus, spanMessage)
}
