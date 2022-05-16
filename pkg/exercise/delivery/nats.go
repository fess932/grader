package delivery

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog/log"
	"grader/configs"
	"grader/pkg/exercise"
	"time"
)

const timeout = time.Second * 3

func NewNatsDelivery(config configs.ServerConfig) (*NatsDelivery, error) {
	host := fmt.Sprintf("nats://%v:%v", config.NatsHost, config.NatsPort)
	log.Debug().Msg("Connecting to NATS: " + host)
	time.Sleep(timeout)

	nc, err := nats.Connect(host, nats.Timeout(timeout))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to NATS: %w", err)
	}

	return &NatsDelivery{
		nc: nc,
	}, nil
}

type NatsDelivery struct {
	nc *nats.Conn
}

func (n NatsDelivery) Publish(exercise exercise.Exercise) error {
	msg, err := json.Marshal(exercise)
	if err != nil {
		return fmt.Errorf("failed to marshal exercise: %w", err)
	}

	if err = n.nc.Publish("exercise", msg); err != nil {
		return fmt.Errorf("failed to publish exercise: %w", err)
	}

	return nil
}
