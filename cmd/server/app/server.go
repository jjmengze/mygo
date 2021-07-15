package app

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"k8s.io/klog"
	"mygo/pkg/signal"
	"mygo/pkg/telemetry"
	"os"
	"time"
)

const (
	// component name
	component = "server"
)

func NewServerCommand() *cobra.Command {
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

	flushTracer, err := telemetry.NewTracerProvider(ctx, options.ComponentConfig.Telemetry)
	if err != nil {
		klog.Warning("tracing config error:", err)
	}
	defer flushTracer()
	tracer := otel.Tracer("test-tracer")
	commonLabels := []attribute.KeyValue{
		attribute.String("labelA", "chocolate"),
		attribute.String("labelB", "raspberry"),
		attribute.String("labelC", "vanilla"),
	}

	tracerCtx, span := tracer.Start(
		ctx,
		"CollectorExporter-Example",
		trace.WithAttributes(commonLabels...))
	defer span.End()

	childCtx, iSpan := tracer.Start(tracerCtx, fmt.Sprintf("start-%d", -1))
	for i := 0; i < 10; i++ {
		func(c context.Context) {
			var sp trace.Span
			childCtx, sp = tracer.Start(childCtx, fmt.Sprintf("Sample-%d", i))
			<-time.After(time.Second)
			sp.End()
		}(childCtx)

	}
	iSpan.End()

	//cc, sched, err := Setup(ctx, opts, registryOptions...)
	//if err != nil {
	//	return err
	//}

	return Run(ctx)
}

// Run executes the Server based on the given configuration. It only returns on error or when context is done.
func Run(stopCh context.Context) error {
	// To help debugging, immediately log version
	//klog.Infof("Version: %+v", version.Get())
	return nil
}
