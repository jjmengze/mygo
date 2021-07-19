package server

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"k8s.io/klog"
	"mygo/pkg/valid"
	"net"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
	listener net.Listener
}

// New creates a new server listening on the provided address that responds to
// the http.Handler. It starts the listener, but does not start the server. If
// an empty port is given, the server randomly chooses one.
func NewServer(config Config) (*Server, error) {
	// Create the net listener first, so the connection ready when we return. This
	// guarantees that it can accept requests.
	ip, port, err := net.SplitHostPort(config.BindAddress)
	if err != nil {
		err := errors.New("must be a valid socket address format, (e.g. 0.0.0.0:10254 or [::]:10254)")
		return nil, err
	}
	portInt, _ := strconv.Atoi(port)
	if err := valid.IsValidPortNum(portInt); err != nil {
		port = "8080"
	}

	if err := valid.IsValidIP(ip); err != nil {
		ip = "0.0.0.0"
	}
	klog.V(3).Infof("server running on %s:%s", ip, port)

	listener, err := net.Listen("tcp", net.JoinHostPort(ip, port))
	if err != nil {
		return nil, fmt.Errorf("failed to create listener on %s: %w", net.JoinHostPort(ip, port), err)
	}

	return &Server{
		listener: listener,
	}, nil
}

// NewFromListener creates a new server on the given listener. This is useful if
// you want to customize the listener type (e.g. udp or tcp) or bind network
// more than `New` allows.
func NewFromListener(listener net.Listener) (*Server, error) {
	return &Server{
		listener: listener,
	}, nil
}

// ServeHTTPHandler is a convenience wrapper around ServeHTTP. It creates an
// HTTP server using the provided handler, wrapped in OpenCensus for
// observability.
func (s *Server) ServeHTTPHandler(ctx context.Context, handler http.Handler) error {
	return s.ServeHTTP(ctx, &http.Server{
		Handler: handler,
	})
}

// ServeHTTP starts the server and blocks until the provided context is closed.
// Once a server has been stopped, it is NOT safe for reuse.
func (s *Server) ServeHTTP(ctx context.Context, srv *http.Server) error {
	btc, _ := context.WithCancel(ctx)
	// Spawn a goroutine that listens for context closure. When the context is
	// closed, the server is stopped.
	errCh := make(chan error, 1)
	go func() {
		<-btc.Done()
		klog.V(3).Infof("server.Serve: receive context closed")
		shutdownCtx, done := context.WithTimeout(context.Background(), 5*time.Second)
		defer done()

		if err := srv.Shutdown(shutdownCtx); err != nil {
			select {
			case errCh <- err:
			default:
			}
		}
	}()

	// Run the server. This will block until the provided context is closed.
	if err := srv.Serve(s.listener); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return errors.New(fmt.Sprintf("failed to serve: %w", err))
	}
	klog.V(3).Infof("server.Serve: serving stopped")

	// Return any errors that happened during shutdown.
	select {
	case err := <-errCh:
		return fmt.Errorf("failed to shutdown: %w", err)
	default:
		return nil
	}
}
