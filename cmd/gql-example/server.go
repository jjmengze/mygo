package main

import (
	"context"
	"github.com/jjmengze/mygo/internal/delivery/graph"
	"github.com/jjmengze/mygo/internal/delivery/graph/generated"
	"github.com/jjmengze/mygo/internal/repo"
	"github.com/jjmengze/mygo/internal/usecase"
	infraRepo "github.com/jjmengze/mygo/pkg/repo"
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
	repoConfig := &infraRepo.Config{
		RetryTime: 10,
		Debug:     true,
		Driver:    infraRepo.MySQL,
		Host:      "127.0.0.1",
		Port:      3306,
		Database:  "blog",
		//InstanceName:   "",//for cloud sql
		User:     "root",
		Password: "123456",
		//SearchPath:     "",//for pg

		ConnectTimeout: time.Second,      //todo fix to times
		ReadTimeout:    10 * time.Second, //todo fix to times
		WriteTimeout:   10 * time.Second, //todo fix to times

		//DialTimeout:    nil,//default setting
		//MaxLifetime:    nil,//default setting
		//MaxIdleTime:    nil,//default setting
		//MaxIdleConn:    nil,//default setting
		//MaxOpenConn:    nil,//default setting
		//SSLMode:        false, //for pg
	}
	read, err := repo.NewGORM(repoConfig)
	if err != nil {
		panic(err)
	}
	write, err := repo.NewGORM(repoConfig)
	if err != nil {
		panic(err)
	}
	repository := repo.NewRepository(read, write)
	userUsecase := usecase.NewUserService(repository)
	resolver := graph.NewResolver(userUsecase)

	//ctx, cancel := context.WithCancel(signal.SetupSignalContext())
	//defer cancel()
	flushTracer, err := telemetry.NewTelemetryProvider(
		context.Background(),
		telemetry.Config{
			Name:     "GQL_Example",
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

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))
	srvHandler := http_server.NewHttpHandler(srv)
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srvHandler)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
