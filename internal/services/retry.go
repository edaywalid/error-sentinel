package services

import (
	"math/rand"
	"net/http"
	"time"
)

func RetryService(env string, maxRetries int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		retries := 0
		for retries < maxRetries {
			sucess := simulateProcessing()
			if sucess {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("Request succeeded!"))
				return
			}
			retries++
		}
		backoffDur := time.Duration(rand.Intn(1000))
		time.Sleep(backoffDur)
		http.Error(w, "Retry limit reached , Failed to process the request", http.StatusInternalServerError)
	}
}

// 30% chance to succeed
func simulateProcessing() bool {
	return rand.Float32() < 0.3
}
