package middlewares

import (
	"log"
	"net/http"
	"time"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// ДО handler
		start := time.Now()

		// ВЫЗВАТЬ handler
		next.ServeHTTP(w, r)

		// ПОСЛЕ handler
		duration := time.Since(start)
		log.Println(r.Method, r.URL.Path, duration)
	})
}
