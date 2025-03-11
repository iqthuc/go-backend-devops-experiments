package middleware

import (
	"log"
	"net/http"
	"time"
)

type responseWriterWrapper struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriterWrapper) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rw := &responseWriterWrapper{w, http.StatusOK}

		next.ServeHTTP(rw, r)
		log.Printf("[%s] | %s | %s | %d | %s", r.Method, r.RequestURI, r.RemoteAddr, rw.statusCode, time.Since(start))
	})
}
