package middleware

import (
	"chatgpt-challenge/delivery/http_server/http_io"
	"log"
	"net/http"
)

func Recovering(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("recovering panic: %v\n", err)
				headers := make(http.Header)
				headers.Set("Connection", "close")
				http_io.WriteJSON(w, http.StatusInternalServerError, http_io.Envelope{Data: http.StatusText(http.StatusInternalServerError)}, headers)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
