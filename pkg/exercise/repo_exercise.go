package exercise

import "database/sql"

type PostgresRepository struct {
	db *sql.DB
}

func (p PostgresRepository) Create(exercise Exercise) error {
	//TODO implement me
	panic("implement me")
}

func (p PostgresRepository) GetByUserID(userID string) ([]Exercise, error) {
	//TODO implement me
	panic("implement me")
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}
