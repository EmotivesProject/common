package middlewares

import (
	"net/http"
	"strings"

	"github.com/TomBowyerResearchProject/common/response"
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
			origin := r.Header.Get("Origin")
			if in := stringInSlice(origin, splitAllowedOrigins); !in {
				response.MessageResponseJSON(w, false, http.StatusMethodNotAllowed, response.Message{Message: "no allowed"})

				return
			}

			w.Header().Add("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", config.AllowedMethods)
			w.Header().Set("Access-Control-Allow-Headers", config.AllowedHeaders)

			if r.Method == "OPTIONS" {
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}

	return false
}
