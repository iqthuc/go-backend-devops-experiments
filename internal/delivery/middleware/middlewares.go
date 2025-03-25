package middleware

import "net/http"

func Apply(middlewares ...func(http.Handler) http.Handler) func(http.Handler) http.Handler {
	return func(finalHandler http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			finalHandler = middlewares[i](finalHandler)
		}
		return finalHandler
	}
}

func Wrap(fn http.HandlerFunc, middlewares ...func(http.Handler) http.Handler) http.HandlerFunc {
	handler := http.HandlerFunc(fn)
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler).ServeHTTP
	}
	return handler
}
