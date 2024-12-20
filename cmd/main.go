package main

import (
	"net/http"
	"os"

	"github.com/edaywalid/error-sentinel/config"
	"github.com/edaywalid/error-sentinel/internal/handlers"
	"github.com/edaywalid/error-sentinel/internal/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.TraceLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	log.Info().Msg("Starting the server")

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Error().Msg("Couldn't load the configs")
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/health", handlers.HealthHandler)
	mux.HandleFunc("/panic", handlers.PanicHandler)
	mux.HandleFunc("/error", handlers.ErrorHandler)
	mux.HandleFunc("/retry", handlers.RetryHandler(cfg.ENV, 3))
	mux.HandleFunc("/source", handlers.SourceHandler)
	handler := middleware.RecoveryMiddleware(cfg.ENV, mux)

	log.Info().Msgf("Starting server on :%s", cfg.PORT)
	err = http.ListenAndServe(":"+cfg.PORT, handler)
	if err != nil {
		log.Fatal().Msgf("Server failed to start: %v", err)
	}
}
