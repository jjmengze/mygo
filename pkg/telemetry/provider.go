package telemetry

import (
	"context"
	"errors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/propagation"
	export "go.opentelemetry.io/otel/sdk/export/metric"
	"go.opentelemetry.io/otel/sdk/metric/aggregator/histogram"
	controller "go.opentelemetry.io/otel/sdk/metric/controller/basic"
	processor "go.opentelemetry.io/otel/sdk/metric/processor/basic"
	selector "go.opentelemetry.io/otel/sdk/metric/selector/simple"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"google.golang.org/grpc"
	"k8s.io/klog"
	"mygo/pkg/valid"
	"net"
	"net/http"
	"strconv"
	"time"
)

var defaultTimeout = 1 * time.Second

func NewTelemetryProvider(ctx context.Context, config Config) (func() error, error) {
	var tracerProvider *sdktrace.TracerProvider
	//Prometheus exporter
	var metricProvider *prometheus.Exporter
	var err error
	//initBackoffTime := 10 * time.Millisecond
	//maxBackoffTime := 1 * time.Second
	//backoffFactor := 2.0
	//jitter := 1.1
	//
	//backoff := backoffmanager.NewExponentialBackoffManager(initBackoffTime, maxBackoffTime, time.Second, backoffFactor, jitter, time.Now)
	//retryTime := 0
	//select {
	//case <-backoff.Backoff().C:
	if config.Jaeger != nil {
		tracerProvider, err = initJaegerTracerProvider(ctx, config)
	} else {
		tracerProvider, err = initCollectorProvider(ctx, config)
	}
	//if err != nil || retryTime == 5 {
	if err != nil {

		return nil, err
	}
	//retryTime++
	//}

	if config.Prometheus != nil {
		metricProvider, err = initPrometheusProvider(ctx, config.Prometheus)
	}

	if err != nil {
		return nil, err
	}

	otel.SetTracerProvider(tracerProvider)

	// set global propagator to tracecontext (the default is no-op).
	otel.SetTextMapPropagator(propagation.TraceContext{})

	global.SetMeterProvider(metricProvider.MeterProvider())
	return func() error {
		// Shutdown will flush any remaining spans and shut down the exporter.
		return tracerProvider.Shutdown(ctx)
	}, nil
}

//	initCollectorProvider offers a vendor-agnostic implementation on how to receive, process and export telemetry data.
func initCollectorProvider(ctx context.Context, config Config) (*sdktrace.TracerProvider, error) {
	res := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(config.Name),
	)

	client := otlptracegrpc.NewClient(otlptracegrpc.WithTimeout(time.Duration(2*time.Second)),
		otlptracegrpc.WithRetry(otlptracegrpc.RetrySettings{
			Enabled:         true,
			MaxElapsedTime:  time.Minute,
			InitialInterval: time.Nanosecond,
			MaxInterval:     time.Nanosecond,
		}),
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithTimeout(10*time.Millisecond),
		otlptracegrpc.WithEndpoint(config.EndPoint),
		otlptracegrpc.WithReconnectionPeriod(50*time.Millisecond),
		otlptracegrpc.WithDialOption(grpc.WithBlock()), // useful for testing
	)

	// Set up a trace exporter
	traceExporter, err := otlptrace.New(ctx, client)

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

func initPrometheusProvider(ctx context.Context, config *Prometheus) (*prometheus.Exporter, error) {
	conf := prometheus.Config{}
	controller := controller.New(
		processor.New(
			selector.NewWithHistogramDistribution(
				histogram.WithExplicitBoundaries(conf.DefaultHistogramBoundaries),
			),
			export.CumulativeExportKindSelector(),
			processor.WithMemory(true),
		),
	)
	exporter, err := prometheus.New(conf, controller)
	if err != nil {
		return nil, err
	}

	ip, port, err := net.SplitHostPort(config.MetricsBindAddress)
	if err != nil {
		err := errors.New("must be a valid socket address format, (e.g. 0.0.0.0:10254 or [::]:10254)")
		return nil, err
	}
	portInt, _ := strconv.Atoi(port)
	if err := valid.IsValidPortNum(portInt); err != nil {
		port = "10251"
	}

	if err := valid.IsValidIP(ip); err != nil {
		ip = "0.0.0.0"
	}
	http.HandleFunc("/metric", exporter.ServeHTTP)
	go func() {
		_ = http.ListenAndServe(net.JoinHostPort(ip, port), nil)
	}()
	klog.Infof("Metric server running on %s:%s", ip, port)
	return exporter, err
}
