package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"grader/configs"
	"grader/pkg/grader"
	"grader/pkg/grader/delivery"
	"os"
	"time"
)

func main() {
	if _, ok := os.LookupEnv("PRODUCTION"); !ok {
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMilli}).
			With().Caller().Logger()
	}

	config := configs.NewGraderConfig()

	log.Debug().Msgf("start grader server: %v", config.Addr)

	delivery.NewGraderService(
		grader.NewGrader(*config),
		*config,
	).Run()
}
