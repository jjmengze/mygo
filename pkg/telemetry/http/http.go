package http

import (
	"fmt"
	"go.opentelemetry.io/otel/propagation"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"net/http"
)

type Transport struct {
	config *config
}

var _ http.RoundTripper = &Transport{}

func NewTransport(opts ...Option) *Transport {

	t := &Transport{}

	c := newConfig(opts...)
	t.config = c

	return t
}

func (t *Transport) RoundTrip(request *http.Request) (*http.Response, error) {
	for i := 0; i < len(t.config.filterFunc); i++ {
		filterFunc := t.config.filterFunc[i]
		if filterFunc(request) {
			return t.config.base.RoundTrip(request)
		}
	}
	tracer := t.config.TracerProvider.Tracer("")
	ctx, span := tracer.Start(request.Context(), fmt.Sprintf("Send to : %s ", request.URL))

	defer span.End()
	req := request.WithContext(ctx)
	span.SetAttributes(semconv.HTTPClientAttributesFromHTTPRequest(req)...)
	t.config.Propagators.Inject(ctx, propagation.HeaderCarrier(request.Header))

	res, err := t.config.base.RoundTrip(req)
	if err != nil {
		span.RecordError(err)
		return res, err
	}
	span.SetAttributes(semconv.HTTPAttributesFromHTTPStatusCode(res.StatusCode)...)
	span.SetStatus(semconv.SpanStatusFromHTTPStatusCode(res.StatusCode))
	return res, err
}
