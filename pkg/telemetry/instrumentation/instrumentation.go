package instrumentation

import (
	"net/http"
)

type RecordingResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

func NewRecordingResponseWriter(writer http.ResponseWriter) *RecordingResponseWriter {
	return &RecordingResponseWriter{
		ResponseWriter: writer,
	}
}

func (w *RecordingResponseWriter) WriteHeader(status int) {
	w.StatusCode = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *RecordingResponseWriter) Header() http.Header {
	return w.ResponseWriter.Header()
}

func (w *RecordingResponseWriter) Write(b []byte) (int, error) {
	return w.ResponseWriter.Write(b)
}
