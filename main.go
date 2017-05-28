package waf

import (
	"log"
	"net/http"
)

func WAF(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("WAF init")
		w.Header().Set("X-WAF-Version", "1.0")
		next.ServeHTTP(w, r)
	})
}
