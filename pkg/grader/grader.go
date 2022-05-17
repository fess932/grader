package grader

import (
	"github.com/rs/zerolog/log"
	"grader/configs"
	"os/exec"
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

	cmd := exec.Command("echo", "hello")

	v, err := cmd.Output()
	if err != nil {
		log.Error().Err(err).Msg("error")
	}

	log.Info().Msg(string(v))
}
