package config

import "mygo/pkg/telemetry"

type Config struct {
	// HealthzBindAddress is the IP address and port for the health check server to serve on,
	// defaulting to 0.0.0.0:10251
	HealthzBindAddress string `json:"healthzBindAddress"`

	// Telemetry is the tracing config ,should setup collection server endpoint and
	// support jaeger and OpenTelemetry Collector
	Telemetry telemetry.Config `json:"telemetry"`
}
