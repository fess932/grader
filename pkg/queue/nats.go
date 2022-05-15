package queue

import (
	"github.com/nats-io/nats-server/v2/server"
	"github.com/rs/zerolog/log"
	"grader/configs"
)

type NatsQueue struct {
	s *server.Server
}

func NewNatsQueue(config configs.ServerConfig) *NatsQueue {
	// Setup NATS exercise
	nsrv, err := server.NewServer(&server.Options{
		Host: config.NatsHost,
		Port: config.NatsPort,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create NATS exercise")
	}

	return &NatsQueue{
		s: nsrv,
	}
}

func (n *NatsQueue) Run() {
	log.Debug().Msg("Starting queue NATS exercise")

	go n.s.Start()
}
