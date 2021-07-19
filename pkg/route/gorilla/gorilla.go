package gorilla

import (
	"github.com/gorilla/mux"
)

func NewGorillaServer() (*mux.Router, error) {
	// Create the router
	r := mux.NewRouter()
	return r, nil
}
