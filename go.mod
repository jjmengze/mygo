module mygo

go 1.16

require (
	github.com/google/wire v0.5.0
	github.com/labstack/echo/v4 v4.4.0
	github.com/lib/pq v1.10.0 // indirect
	github.com/mattn/go-sqlite3 v1.14.6 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/spf13/cobra v1.2.1
	github.com/spf13/pflag v1.0.5
	go.opentelemetry.io/otel v1.0.0-RC1
	go.opentelemetry.io/otel/exporters/jaeger v1.0.0-RC1
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.0.0-RC1
	go.opentelemetry.io/otel/metric v0.21.0
	go.opentelemetry.io/otel/sdk v1.0.0-RC1
	go.opentelemetry.io/otel/sdk/metric v0.21.0
	go.opentelemetry.io/otel/trace v1.0.0-RC1
	google.golang.org/grpc v1.38.0
	gorm.io/driver/mysql v1.1.1
	gorm.io/driver/postgres v1.1.0
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.21.12
	k8s.io/klog v1.0.0
	sigs.k8s.io/yaml v1.2.0
)
