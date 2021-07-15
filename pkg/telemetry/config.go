package telemetry

type Config struct {
	Name               string       `json:"Name"`
	ResourceAttributes []Attributes `json:"ResourceAttributes"`
	EndPoint           string       `json:"EndPoint"`
	Jaeger             *Jaeger      `json:"Jaeger"`
}

type Jaeger struct {
	Password string `json:"Password"`
	UserName string `json:"UserName"`
}

type Attributes struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}
