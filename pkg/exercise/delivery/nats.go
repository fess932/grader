package delivery

import (
	"grader/pkg/exercise"
)

type NatsDelivery struct {
}

func (n NatsDelivery) Publish(exercise exercise.Exercise) error {
	//TODO implement me
	panic("implement me")
}

func NewNatsDelivery() *NatsDelivery {
	return &NatsDelivery{}
}
