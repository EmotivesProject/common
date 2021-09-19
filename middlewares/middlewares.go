package middlewares

import (
	"net/http"
	"strings"
)

type Config struct {
	AllowedOrigins string
	AllowedHeaders string
	AllowedMethods string
}

var (
	config              Config
	splitAllowedOrigins []string
)

func Init(middlewareConfig Config) {
	config = middlewareConfig
	splitAllowedOrigins = strings.Split(config.AllowedOrigins, ",")
}

func SimpleMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Possibly set Access-Control-Allow-Origin to something other than the first
			origin := r.Header.Get("Origin")
			if index := stringInSlice(origin, splitAllowedOrigins); index != -1 {
				w.Header().Set("Access-Control-Allow-Origin", splitAllowedOrigins[index])
			} else if origin != "" {
				w.Header().Set("Access-Control-Allow-Origin", splitAllowedOrigins[0])
			}

			w.Header().Add("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Methods", config.AllowedMethods)
			w.Header().Set("Access-Control-Allow-Headers", config.AllowedHeaders)

			if r.Method == "OPTIONS" {
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func stringInSlice(a string, list []string) int {
	for idx, b := range list {
		if b == a {
			return idx
		}
	}

	return -1
}
