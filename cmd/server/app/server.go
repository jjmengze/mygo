package app

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"k8s.io/klog"
	"mygo/pkg/signal"
	"os"
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
