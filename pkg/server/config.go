package server

import (
	"github.com/spf13/pflag"
)

type Config struct {
	// bindAddress is the IP address for the proxy server to serve on (set to 0.0.0.0
	// for all interfaces)
	BindAddress string
}

// AddUniversalFlags adds flags for a specific APIServer to the specified FlagSet
func (c *Config) AddFlags(fs *pflag.FlagSet) {
	if c == nil {
		return
	}
	fs.StringVar(&c.BindAddress, "advertise-address", c.BindAddress, ""+
		"The IP address on which to advertise the apiserver to members of the cluster. This "+
		"address must be reachable by the rest of the cluster. If blank, the --bind-address "+
		"will be used. If --bind-address is unspecified, the host's default interface will "+
		"be used.")
}
