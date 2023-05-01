package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sr := &statusRecorder{
			ResponseWriter: w,
			Status:         http.StatusOK,
		}
		start := time.Now()
		next.ServeHTTP(sr, r)
		log.Printf("%d %s %s%s %s", sr.Status, r.Method, r.Host, r.RequestURI, time.Since(start))
	})
}

type AuthValidator func(string, string) (bool, error)

func BasicAuth(next http.Handler, av AuthValidator) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		login, password, ok := r.BasicAuth()

		if ok {
			valid, _ := av(login, password)
			if valid {
				next.ServeHTTP(w, r)
				return
			}
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "unauthorized",
		})
	})
}
