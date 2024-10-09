package middleware

import (
	"log"
	"net/http"
	"time"
)

// wrappedWriter is a custom http.ResponseWriter that captures the HTTP status code
// in addition to providing the standard ResponseWriter functionality.
type wrappedWriter struct {
	http.ResponseWriter
	status int
}

func (w *wrappedWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &wrappedWriter{ResponseWriter: w}

		next.ServeHTTP(wrapped, r)

		log.Printf("%d %s %s %s", wrapped.status, r.Method, r.URL.Path, time.Since(start))
	})
}
