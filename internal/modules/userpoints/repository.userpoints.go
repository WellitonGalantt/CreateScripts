package userpoints

import "database/sql"

type UserPointsRepository interface {
	GetById(id int) (*GetByIdDTOOutput, error)
	Credit(quantity int) error
	Debit(quantity int) error
}

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) UserPointsRepository {
	return &postgresRepository{
		db,
	}
}

func (r *postgresRepository) GetById(id int) (*GetByIdDTOOutput, error) {
	return nil, nil
}

func (r *postgresRepository) Credit(quantity int) error {
	return nil
}

func (r *postgresRepository) Debit(quantity int) error {
	return nil
}
