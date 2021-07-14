package app

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"kubernetes/cmd/kube-scheduler/app/options"
	"kubernetes/staging/src/k8s.io/component-base/version/verflag"
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
			if err := runCommand(cmd, opts, registryOptions...); err != nil {
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
	//fs := cmd.Flags()
	return cmd
}

// runCommand runs the server.
func runCommand(cmd *cobra.Command, opts *options.Options, registryOptions ...Option) error {
	verflag.PrintAndExitIfRequested()
	cliflag.PrintFlags(cmd.Flags())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cc, sched, err := Setup(ctx, opts, registryOptions...)
	if err != nil {
		return err
	}

	return Run(ctx, cc, sched)
}

// Run runs the Server Options.  This should never exit.
//func Run(c *config.CompletedConfig, stopCh <-chan struct{}) error {
func Run(stopCh <-chan struct{}) error {
	// To help debugging, immediately log version
	//klog.Infof("Version: %+v", version.Get())
}
