package instrumentation

import (
	"net/http"
)

type RecordingResponseWriter struct {
	http.ResponseWriter
	Written    int64
	StatusCode int
	Error      error
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
	n, err := w.ResponseWriter.Write(b)
	n1 := int64(n)
	w.Written += n1
	w.Error = err
	return n, err
}
