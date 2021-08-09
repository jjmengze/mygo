package http_server

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"net/http"
)

type config struct {
	tracerProvider trace.TracerProvider
	filterFunc     []func(*http.Request) bool
}

// Option interface used for setting optional config properties.
type Option interface {
	apply(*config)
}

func newConfig(options ...Option) *config {
	c := &config{
		tracerProvider: otel.GetTracerProvider(),
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

func WithFilterFunc(filterFunc ...func(*http.Request) bool) Option {
	return optionFunc(func(c *config) {
		c.filterFunc = filterFunc
	})
}
