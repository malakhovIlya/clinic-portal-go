package config

import (
	"net/http"
	"strings"

	"github.com/go-chi/cors"
)

const (
	corsEnvVar         = "CORS_ALLOWED_ORIGINS"
	defaultCORSOrigins = "http://localhost:5173,http://localhost:5174,http://localhost:3000,http://localhost:8080,http://127.0.0.1:8080"
	allowAllWildcard   = "*"
	defaultMaxAge      = 300
)

// NewCORS returns a Chi middleware that enables cross-origin requests using
// the origin list defined in CORS_ALLOWED_ORIGINS (comma separated).
// Use "*" to allow any origin (e.g. for staging).
func NewCORS() func(http.Handler) http.Handler {
	opts := cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           defaultMaxAge,
	}

	origins := readCORSOrigins()
	if len(origins) == 1 && origins[0] == allowAllWildcard {
		opts.AllowOriginFunc = func(_ *http.Request, _ string) bool {
			return true
		}
	} else {
		opts.AllowedOrigins = origins
	}

	return cors.Handler(opts)
}

func readCORSOrigins() []string {
	raw := getEnv(corsEnvVar, defaultCORSOrigins)
	parts := strings.Split(raw, ",")

	var origins []string
	for _, part := range parts {
		if trimmed := strings.TrimSpace(part); trimmed != "" {
			origins = append(origins, trimmed)
		}
	}

	if len(origins) == 0 {
		return []string{allowAllWildcard}
	}

	return origins
}
