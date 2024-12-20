package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/edaywalid/error-sentinel/internal/utils"
	"github.com/rs/zerolog/log"
)

func RecoveryMiddleware(env string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Info().Msgf("Recovered from panic: %v\n", err)
				stack := string(debug.Stack())
				log.Trace().Msgf("Stack trace: %s", stack)

				if env == "production" {
					http.Error(w, "Something went wrong.", http.StatusInternalServerError)
				} else {
					w.Header().Set("Content-Type", "text/html")
					w.WriteHeader(http.StatusInternalServerError)
					fmt.Fprintf(w, "<h1>Panic occurred: %v</h1><pre>%s</pre>", err, utils.FormatStackTrace(stack))
				}
			}
		}()
		next.ServeHTTP(w, r)
	})
}
