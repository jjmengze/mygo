package app

import (
	"context"
	"fmt"
	"github.com/jjmengze/mygo/pkg/signal"
	"github.com/jjmengze/mygo/pkg/telemetry"
	telemetryClient "github.com/jjmengze/mygo/pkg/telemetry/http_client"
	"github.com/spf13/cobra"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/trace"
	"k8s.io/klog"
	"net/http"
	"os"
	"time"
)

const (
	// component name
	component = "client"
)

func NewClientCommand() *cobra.Command {
	opts := NewOptions()
	cmd := &cobra.Command{
		Use: component,
		Run: func(cmd *cobra.Command, args []string) {
			if err := opts.Complete(); err != nil {
				klog.Fatalf("failed complete: %v", err)
			}
			fmt.Println(opts)
			fmt.Println(opts.ComponentConfig)

			if err := runCommand(cmd, opts); err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
		},
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		},
	}
	fs := cmd.Flags()
	opts.Flags(fs)
	//namedFlagSets :=
	//for _, f := range namedFlagSets.FlagSets {
	//	fs.AddFlagSet(f)
	//}
	//usageFmt := "Usage:\n  %s\n"
	//cmd.SetUsageFunc(func(cmd *cobra.Command) error {
	//	fmt.Fprintf(cmd.OutOrStderr(), usageFmt, cmd.UseLine())
	//	return nil
	//})
	//cmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
	//	fmt.Fprintf(cmd.OutOrStdout(), "%s\n\n"+usageFmt, cmd.Long, cmd.UseLine())
	//})
	cmd.MarkFlagFilename("config", "yaml", "yml", "json")
	return cmd
}

// runCommand runs the server.
func runCommand(cmd *cobra.Command, options *Options) error {
	//verflag.PrintAndExitIfRequested()
	//cliflag.PrintFlags(cmd.Flags())

	ctx, cancel := context.WithCancel(signal.SetupSignalContext())
	defer cancel()

	flushTracer, err := telemetry.NewTelemetryProvider(ctx, options.ComponentConfig.Telemetry)
	if err != nil {
		klog.Warning("tracing config error:", err)
	}
	defer flushTracer()

	tracer := otel.Tracer("client")
	meter := global.Meter("test-client")
	commonLabels := []attribute.KeyValue{
		attribute.String("labelA", "chocolate"),
		attribute.String("labelB", "raspberry"),
		attribute.String("labelC", "vanilla"),
	}

	tracerCtx, span := tracer.Start(
		ctx,
		"Client-example-request",
		trace.WithAttributes(commonLabels...))
	defer span.End()

	childCtx, iSpan := tracer.Start(tracerCtx, fmt.Sprintf("start"))

	workTime := metric.Must(meter).
		NewInt64Counter(
			"test time",
			metric.WithDescription("The worker tested time"),
		).Bind(commonLabels...)
	defer workTime.Unbind()

	requestLatency := metric.Must(meter).
		NewFloat64ValueRecorder(
			"test request_latency",
			metric.WithDescription("The latency of requests processed"),
		).Bind(commonLabels...)
	defer requestLatency.Unbind()

	latencyMs := float64(time.Since(time.Now())) / 1e6
	httpClient := telemetryClient.HttpClientWithTransport(http.DefaultTransport)

	for i := 0; i < 10; i++ {
		func(c context.Context) {
			var sp trace.Span

			childCtx, sp = tracer.Start(childCtx, fmt.Sprintf("Sample-%d", i))
			defer sp.End()

			req, err := http.NewRequestWithContext(childCtx, "GET", options.ComponentConfig.ExampleServer+"/happy", nil)
			if err != nil {
				panic(err)
			}

			req, err = http.NewRequestWithContext(childCtx, "GET", "http://0.0.0.0:8080/happy", nil)
			if err != nil {
				panic(err)
			}
			fmt.Printf("Sending request...\n")
			_, err = httpClient.Do(req)
			if err != nil {
				panic(err)
			}

			latencyMs = float64(time.Since(time.Now())) / 1e6
			<-time.After(time.Millisecond)
			requestLatency.Record(ctx, latencyMs)
		}(childCtx)
	}
	defer iSpan.End()

	//cc, sched, err := Setup(ctx, opts, registryOptions...)
	//if err != nil {
	//	return err
	//}

	return Run(ctx)
}

// Run executes the Server based on the given configuration. It only returns on error or when context is done.
func Run(ctx context.Context) error {
	// To help debugging, immediately log version
	//klog.Infof("Version: %+v", version.Get())
	return nil
}
