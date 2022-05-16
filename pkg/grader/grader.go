package grader

import "grader/configs"

type Grader struct {
}

func (g Grader) Grade() {
	//TODO implement me
	panic("implement me")
}

func NewGrader(c configs.GraderConfig) *Grader {
	return &Grader{}
}
