package instrumentation

import "net/http"

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
	w.ResponseWriter.WriteHeader(status)
}

func (w *RecordingResponseWriter) Write(b []byte) (int, error) {
	return w.ResponseWriter.Write(b)
}
