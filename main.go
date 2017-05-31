package waf

import (
	"log"
	"net/http"
	"time"
)

func WAF(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time := time.Now().UTC().String()
		log.Println("WAF init")
		w.Header().Set("sq-waf-version", time)
		r.Header.Set("sq-waf-version", time)
		next.ServeHTTP(w, r)
	})
}
