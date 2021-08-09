package telemetry

type Config struct {
	Name               string       `json:"name"`
	ResourceAttributes []Attributes `json:"resourceAttributes"`
	EndPoint           string       `json:"endPoint"`
	Jaeger             *Jaeger      `json:"jaeger"`
	Prometheus         *Prometheus  `json:"prometheus"`
}
type Prometheus struct {
	Name string `json:"name"`
	// MetricsBindAddress is the IP address and port for the metrics server to
	// serve on, defaulting to 0.0.0.0:10251.
	MetricsBindAddress string `json:"metricsBindAddress"`
}

type Jaeger struct {
	Mode     Mode
	Password string `json:"password"`
	UserName string `json:"userName"`
}

type Mode string

var Agent Mode = "Agent"
var Collector Mode = "Collector"

type Attributes struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
