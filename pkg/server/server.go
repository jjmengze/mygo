package server

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Set is a Wire provider set that produces a *Server given the fields of
// Options.
var Set = wire.NewSet(
	NewServer,
	wire.Struct(new(Options), "RequestLogger", "HealthChecks", "TraceExporter", "DefaultSamplingPolicy", "Driver"),
	//wire.Value(&DefaultDriver{}),
	wire.Bind(new(http.Handler), echo.New()),
)

// Options is the set of optional parameters.
type Options struct {
	// Driver serves HTTP requests.
	router http.Handler
}

type Server struct {
	handler http.Handler
}

func NewServer(h http.Handler) *Server {
	srv := &Server{handler: h}
	return srv
}
