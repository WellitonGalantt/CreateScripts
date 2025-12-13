package plans

import "database/sql"

type PlansRepository interface {
	GetAllPlans() ([]Plans, error)
}

type postgresRepository struct {
	db *sql.DB
}

func (r *postgresRepository) NewPostgresRepository(db *sql.DB) PlansRepository {
	return &postgresRepository{
		db,
	}
}

func (r *postgresRepository) GetAllPlans() ([]Plans, error) {
	return nil, nil
}
