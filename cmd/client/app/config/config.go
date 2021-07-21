package config

import "mygo/pkg/telemetry"

type Config struct {
	// ExampleServer is the IP address and port to request the example server,
	// defaulting to 0.0.0.0:10251
	ExampleServer string `json:"exampleServer"`

	// Telemetry is the tracing config ,should setup collection server endpoint and
	// support jaeger and OpenTelemetry Collector
	Telemetry telemetry.Config `json:"telemetry"`
}
