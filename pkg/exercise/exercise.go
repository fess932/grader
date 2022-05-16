package exercise

import (
	"grader/pkg/user"
)

// type Try struct {
//	TryID  string `json:"tryID"`
//	NumTry int    `json:"numTry"`
//
//	ID     string `json:"id"`
//	UserID string `json:"userID"`
//	Files  []File `json:"files"`
// }

type Exercise struct {
	TryID  string `json:"tryID"`
	NumTry int    `json:"numTry"`

	ID     string `json:"id"`
	UserID string `json:"userID"`
	Files  []File `json:"files"`
}

type File struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type Publisher interface {
	Publish(exercise Exercise) error
}

type ExersiceRepository interface {
	Create(exercise Exercise) error
	GetByUserID(userID string) ([]Exercise, error)
}

type ExersiceUsecase struct {
	publisher Publisher
	repo      ExersiceRepository
}

func NewExersiceUsecase(p Publisher, r ExersiceRepository) *ExersiceUsecase {
	return &ExersiceUsecase{
		publisher: p,
		repo:      r,
	}
}

func (e *ExersiceUsecase) GetExercise(user user.User, id string) (Exercise, error) {
	return Exercise{}, nil
}

func (e *ExersiceUsecase) CheckExercise(user user.User, exercise Exercise) error {
	//e.repo.Create(exercise)

	return e.publisher.Publish(exercise)
}
