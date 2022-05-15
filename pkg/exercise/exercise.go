package exercise

import "grader/pkg/user"

type Exercise struct {
	Name  string
	Files []string
}

type ExersiceUsecase struct {
}

func NewExersiceUsecase() *ExersiceUsecase {
	return &ExersiceUsecase{}
}

func (e ExersiceUsecase) GetExercise(user user.User, id string) (Exercise, error) {
	return Exercise{}, nil
}

func (e ExersiceUsecase) CheckExercise(user user.User, exercise Exercise) error {
	return nil
}
