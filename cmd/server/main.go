package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"grader/configs"
	"grader/pkg/exercise"
	"grader/pkg/queue"
	"grader/pkg/rest"
	"os"
	"time"
)

func main() {
	if _, ok := os.LookupEnv("PRODUCTION"); !ok {
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMilli}).
			With().Caller().Logger()
	}

	config := configs.NewServerConfig()
	queue.NewNatsQueue(config).Run()

	app := exercise.New()

	rest.NewRest(app, config).Serve() // blocking
}
