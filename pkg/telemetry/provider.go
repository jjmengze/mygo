package telemetry

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"google.golang.org/grpc"
)

func NewTracerProvider(ctx context.Context, config Config) (func() error, error) {
	//var tracerProvider oteltest.TracerProvider
	var tracerProvider *sdktrace.TracerProvider
	var err error
	if config.Jaeger != nil {
		tracerProvider, err = initJaegerTracerProvider(ctx, config)
	} else {
		tracerProvider, err = initCollectorProvider(ctx, config)
	}
	if err != nil {
		return nil, err
	}

	otel.SetTracerProvider(tracerProvider)

	// set global propagator to tracecontext (the default is no-op).
	otel.SetTextMapPropagator(propagation.TraceContext{})

	return func() error {
		// Shutdown will flush any remaining spans and shut down the exporter.
		return tracerProvider.Shutdown(ctx)
	}, nil
}

//	 initCollectorProvider offers a vendor-agnostic implementation on how to receive, process and export telemetry data.
func initCollectorProvider(ctx context.Context, config Config) (*sdktrace.TracerProvider, error) {
	res := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(config.Name),
	)

	// Set up a trace exporter
	traceExporter, err := otlptracegrpc.New(ctx,
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint("localhost:30080"),
		otlptracegrpc.WithDialOption(grpc.WithBlock()),
	)
	if err != nil {
		return nil, err
	}

	// Register the trace exporter with a TracerProvider, using a batch
	// span processor to aggregate spans before export.
	bsp := sdktrace.NewBatchSpanProcessor(traceExporter)
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(bsp),
	)

	return tracerProvider, nil
}

func initJaegerTracerProvider(ctx context.Context, config Config) (*sdktrace.TracerProvider, error) {
	res := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(config.Name),
	)

	//todo maybe jaeeger server should login wih username and password
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(config.EndPoint)))
	if err != nil {
		return nil, err
	}
	tracerProvider := sdktrace.NewTracerProvider(
		// Always be sure to batch in production.
		sdktrace.WithBatcher(exp),
		// Record information about this application in an Resource.
		sdktrace.WithResource(res),
	)

	return tracerProvider, nil
}
