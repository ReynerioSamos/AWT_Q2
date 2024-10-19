package middleware

import(
	"log"
	"net/http"
)

func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Got request with path %s", r.URL.Path)

		next.ServeHTTP(w, r)
	})
}