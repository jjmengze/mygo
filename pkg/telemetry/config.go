package telemetry

type Config struct {
	Name               string
	ResourceAttributes []Attributes
	EndPoint           string
	Jaeger             *Jaeger
}

type Jaeger struct {
	Password string
	UserName string
}

type Attributes struct {
	Key   string
	Value string
}
