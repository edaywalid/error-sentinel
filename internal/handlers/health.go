package handlers

import (
	"net/http"
	"time"
)

var startTime = time.Now()

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	uptime := time.Since(startTime)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK - Uptime: " + uptime.String()))
}
