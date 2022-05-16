package exercise

import "grader/pkg/user"

type Exercise struct {
	ID     string
	UserID string
	Files  []File
}

type File struct {
	Name string
	Path string
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
