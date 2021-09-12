package simple_factory

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

//DoRequest is http RoundTrip interface
type DoRequest interface {
	RoundTrip(r *http.Request) (*http.Response, error)
}

type Filter int

const (
	Default = iota
	FilterPing
)

//NewDoRequest return DoRequest instance by type
func NewDoRequest(f Filter) DoRequest {
	switch f {
	case Default:
		return &DefaultRoundTripper{
			base: http.DefaultTransport,
		}
	case FilterPing:
		return &FilterPingRoundTripper{}
	default:
		return nil
	}
}

//DefaultRoundTripper is one of DefaultRoundTripper implement
type DefaultRoundTripper struct {
	base http.RoundTripper
}

//RoundTrip implement http DefaultRoundTripper
func (d *DefaultRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	return d.base.RoundTrip(r)
}

//FilterPingRoundTripper implement filter ping roundTrip
type FilterPingRoundTripper struct {
	base http.RoundTripper
}

//RoundTrip filter ping request
func (f *FilterPingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	if strings.Contains(r.RequestURI, "ping") {
		return nil, errors.New(fmt.Sprintf("have been filter ping request %v", r.RequestURI))
	}

	return f.base.RoundTrip(r)
}
