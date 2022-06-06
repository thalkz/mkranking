package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func LoggerHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		end := time.Now()
		logRequest(r, start, end)
	})
}

func logRequest(r *http.Request, start, end time.Time) {
	ipAddr := strings.Split(r.RemoteAddr, ":")[0]
	duration := end.UnixMilli() - start.UnixMilli()
	fmt.Printf("[%v] %v %v %v (%vms)\n", start.Format("2006-01-02 15:04:05"), ipAddr, r.Method, r.URL, duration)
}
