package handlers

import (
	"net/http"

	"github.com/edaywalid/error-sentinel/internal/services"
)

func RetryHandler(env string, maxRetries int) http.HandlerFunc {
	return services.RetryService(env, maxRetries)
}
