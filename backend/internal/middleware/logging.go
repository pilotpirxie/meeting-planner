package middleware

import (
	"log"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	status int
	size   int
}

func (w *responseWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *responseWriter) Write(dataBytes []byte) (int, error) {
	bytesWritten, writeError := w.ResponseWriter.Write(dataBytes)
	w.size += bytesWritten
	return bytesWritten, writeError
}

func Logging(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		wrappedWriter := &responseWriter{ResponseWriter: w, status: http.StatusOK}

		nextHandler.ServeHTTP(wrappedWriter, r)
		log.Printf("%s %s %d %dB %s",
			r.Method,
			r.URL.Path,
			wrappedWriter.status,
			wrappedWriter.size,
			time.Since(startTime),
		)
	})
}
