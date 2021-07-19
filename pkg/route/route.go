package route

import "net/http"

// APIServer defines routes for the apiserver service.
func NewAPIServer(r http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/", r)
}
