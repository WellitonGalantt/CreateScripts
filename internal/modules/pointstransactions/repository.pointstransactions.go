package pointstransactions

import "database/sql"

type PointsTransactionsRepository interface {
	GetAll() ([]CreatePointsTransactionsDtoOutput, error)
	Create(input CreatePointsTransactionsDtoInput) (*CreatePointsTransactionsDtoOutput, error)
}

type postgresRepository struct {
	db *sql.DB
}

func NewPointsTransactionsRepository(db *sql.DB) PointsTransactionsRepository {
	return &postgresRepository{
		db,
	}
}

func (r *postgresRepository) GetAll() ([]CreatePointsTransactionsDtoOutput, error) {
	return nil, nil
}

func (r *postgresRepository) Create(input CreatePointsTransactionsDtoInput) (*CreatePointsTransactionsDtoOutput, error) {
	return nil, nil
}
