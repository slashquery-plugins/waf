package waf

import (
	"log"
	"net/http"
)

func WAF(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("WAF init")
		w.Header().Set("sq-waf-version", "1.0.0")
		r.Header.Set("sq-waf-version", "1.0.0")
		next.ServeHTTP(w, r)
	})
}
