package middleware

import "net/http"

// Middleware allows middleware layers to be applied around an HTTP handler.
type Middleware func(http.Handler) http.Handler

// CreateStack returns a single http.Handler that represents the composition of
// all provided middleware layers. The first middleware wraps the second, and so on.
func CreateStack(middlewares ...Middleware) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}

		return next
	}
}
