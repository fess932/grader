package grader

import (
	"github.com/rs/zerolog/log"
	"grader/configs"
)

type Grader struct {
}

func NewGrader(c configs.GraderConfig) *Grader {
	return &Grader{}
}

// получает задачу
// проверяет ее на правильность
// возвращает результат проверки

func (g Grader) Grade() {
	log.Debug().Msg("run grade")
}
