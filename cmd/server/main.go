package main

import (
	"database/sql"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"grader/configs"
	"grader/pkg/exercise"
	exerciseDelivery "grader/pkg/exercise/delivery"
	"grader/pkg/grader/delivery"
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
	log.Debug().Msgf("start nats: %v:%v", config.NatsHost, config.NatsPort)
	queue.NewNatsQueue(config).Run()

	client, err := delivery.NewGraderClient(config.GraderAddr)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to grader")
	}

	edelivery, err := exerciseDelivery.NewNatsDelivery(config, client)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create exercise delivery")
	}

	if err = edelivery.StartWorkers(config.QueueWorkers); err != nil {
		log.Fatal().Err(err).Msg("failed to start exercise workers")
	}

	erepo := exercise.NewPostgresRepository(&sql.DB{})

	exerUcase := exercise.NewExersiceUsecase(edelivery, erepo)

	rest.NewRest(exerUcase, config).Serve() // blocking
}
