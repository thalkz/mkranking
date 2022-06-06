package middlewares

import (
	"log"
	"net/http"
	"time"
)

func LoggerHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		end := time.Now()
		if r.Method != "OPTIONS" {
			duration := end.UnixMilli() - start.UnixMilli()
			log.Printf("%v %v (%vms)\n", r.RemoteAddr, r.URL, duration)
		}
	})
}
