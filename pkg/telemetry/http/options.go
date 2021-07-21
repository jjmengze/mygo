package http

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"net/http"
)

// config represents the configuration options available for the http.Handler
// and http.Transport types.
type config struct {
	Propagators    propagation.TextMapPropagator
	TracerProvider trace.TracerProvider
	MeterProvider  metric.MeterProvider

	base       http.RoundTripper
	filterFunc []func(*http.Request) bool
}

// Option interface used for setting optional config properties.
type Option interface {
	apply(*config)
}

func newConfig(options ...Option) *config {
	c := &config{
		Propagators:    otel.GetTextMapPropagator(),
		TracerProvider: otel.GetTracerProvider(),
		MeterProvider:  global.GetMeterProvider(),
	}

	for _, opt := range options {
		opt.apply(c)
	}
	return c
}

type optionFunc func(*config)

func (o optionFunc) apply(c *config) {
	o(c)
}

func WithRoundTripper(base http.RoundTripper) Option {
	return optionFunc(func(c *config) {
		c.base = base
	})
}

func WithFilterFunc(filterFunc ...func(*http.Request) bool) Option {
	return optionFunc(func(c *config) {
		c.filterFunc = filterFunc
	})
}
