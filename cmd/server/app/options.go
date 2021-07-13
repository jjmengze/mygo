package app

type Options struct {
	// ConfigFile is the location of the proxy server's configuration file.
	ConfigFile string
	// WriteConfigTo is the path where the default configuration will be written.
	WriteConfigTo string
	// CleanupAndExit, when true, makes the proxy server clean up iptables and ipvs rules, then exit.
	CleanupAndExit bool
	// WindowsService should be set to true if kube-proxy is running as a service on Windows.
	// Its corresponding flag only gets registered in Windows builds
	WindowsService bool
	// config is the proxy server's configuration object.
	config *Configuration
	// errCh is the channel that errors will be sent
	errCh chan error

	// The fields below here are placeholders for flags that can't be directly mapped into
	// config.KubeProxyConfiguration.
	//
	// TODO remove these fields once the deprecated flags are removed.

	// master is used to override the kubeconfig's URL to the apiserver.
	master string
	// healthzPort is the port to be used by the healthz server.
	healthzPort int32
	// metricsPort is the port to be used by the metrics server.
	metricsPort int32

	// hostnameOverride, if set from the command line flag, takes precedence over the `HostnameOverride` value from the config file
	hostnameOverride string
}

// NewOptions returns initialized Options
func NewOptions() *Options {
	return &Options{
		config: new(Configuration),
		//healthzPort: ports.ProxyHealthzPort,
		//metricsPort: ports.ProxyStatusPort,
		errCh: make(chan error),
	}
}
