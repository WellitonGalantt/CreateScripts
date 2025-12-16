package plans

import "database/sql"

type PlansRepository interface {
	GetAll() ([]Plans, error)
	// implementacao posterior
	EditById(id int) (*Plans, error)
}

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) PlansRepository {
	return &postgresRepository{
		db,
	}
}

func (r *postgresRepository) GetAll() ([]Plans, error) {
	return nil, nil
}

func (r *postgresRepository) EditById(id int) (*Plans, error) {
	return nil, nil
}
