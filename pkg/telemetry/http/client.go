package http

import "net/http"

func HttpClientWithTransport(transport http.RoundTripper) *http.Client {
	tp := NewTransport(WithRoundTripper(transport))
	return &http.Client{
		Transport: tp,
	}
}

func HttpClient() *http.Client {
	tp := NewTransport(WithRoundTripper(http.DefaultTransport))
	return &http.Client{
		Transport: tp,
	}
}
