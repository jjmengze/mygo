package server

import (
	"github.com/spf13/pflag"
	"net"
)

type Config struct {
	// bindAddress is the IP address for the proxy server to serve on (set to 0.0.0.0
	// for all interfaces)
	BindAddress net.IP
	// metricsBindAddress is the IP address and port for the metrics server to serve on,
	// defaulting to 127.0.0.1:10249 (set to 0.0.0.0 for all interfaces)
	MetricsBindAddress string
}

// AddUniversalFlags adds flags for a specific APIServer to the specified FlagSet
func (c *Config) AddUniversalFlags(fs *pflag.FlagSet) {
	fs.IPVar(&c.BindAddress, "advertise-address", c.BindAddress, ""+
		"The IP address on which to advertise the apiserver to members of the cluster. This "+
		"address must be reachable by the rest of the cluster. If blank, the --bind-address "+
		"will be used. If --bind-address is unspecified, the host's default interface will "+
		"be used.")
}
