package server

type Config struct {
	// bindAddress is the IP address for the proxy server to serve on (set to 0.0.0.0
	// for all interfaces)
	BindAddress string
	// metricsBindAddress is the IP address and port for the metrics server to serve on,
	// defaulting to 127.0.0.1:10249 (set to 0.0.0.0 for all interfaces)
	MetricsBindAddress string
}
