package main

import (
	"context"
	"github.com/jjmengze/mygo/internal/delivery/graph"
	"github.com/jjmengze/mygo/internal/delivery/graph/generated"
	"github.com/jjmengze/mygo/pkg/telemetry"
	"github.com/jjmengze/mygo/pkg/telemetry/http_server"
	"go.opentelemetry.io/otel"
	"k8s.io/klog"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	//ctx, cancel := context.WithCancel(signal.SetupSignalContext())
	//defer cancel()
	flushTracer, err := telemetry.NewTelemetryProvider(
		context.Background(),
		telemetry.Config{
			Name:     "GQL",
			EndPoint: "http://0.0.0.0:49931/api/traces",
			Jaeger:   &telemetry.Jaeger{},
		})

	if err != nil {
		klog.Warning("tracing config error:", err)
	}
	defer flushTracer()

	_, sp := otel.GetTracerProvider().Tracer("test").Start(context.Background(), "happy")
	time.Sleep(time.Second)
	sp.End()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	srvHandler := http_server.NewHttpHandler(srv)
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srvHandler)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
