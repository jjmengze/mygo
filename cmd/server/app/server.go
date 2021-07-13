package app

import (
	"github.com/spf13/cobra"
)

const (
	// component name
	component = "server"
)

func NewServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: component,
	}
	return cmd
}
