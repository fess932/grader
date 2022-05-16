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

func (n *NatsDelivery) Publish(exercise exercise.Exercise) error {
	msg, err := json.Marshal(exercise)
	if err != nil {
		return fmt.Errorf("failed to marshal exercise: %w", err)
	}

	for i := 0; i < 30; i++ {
		if err = n.nc.Publish("exercise", msg); err != nil {
			return fmt.Errorf("failed to publish exercise: %w", err)
		}

		time.Sleep(time.Second)
	}

	return nil
}

func (n *NatsDelivery) StartWorkers(workers int) error {
	log.Debug().Msg("Starting NATS workers")

	const chsize = 64
	ch := make(chan *nats.Msg, chsize)

	_, err := n.nc.ChanSubscribe("exercise", ch)
	if err != nil {
		return fmt.Errorf("failed to subscribe to exercise channel: %w", err)
	}

	for ; workers > 0; workers-- {
		go n.worker(ch)
	}

	return nil
}

func (n *NatsDelivery) worker(ch <-chan *nats.Msg) {
	for msg := range ch {
		log.Info().Msgf("input: %s, %s", msg.Data, msg.Header)
		time.Sleep(time.Second * 10)
	}
}
