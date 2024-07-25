package web

import (
	"net/http"
)

type CORS struct {
	AllowedOrigins string
}

func (c *CORS) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", c.AllowedOrigins)
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func NewCORS(allowedOrigins string) *CORS {
	return &CORS{
		AllowedOrigins: allowedOrigins,
	}
}
