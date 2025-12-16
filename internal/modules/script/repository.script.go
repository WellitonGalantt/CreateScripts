package script

import "database/sql"

type ScriptRepository interface {
	Create(input ScriptDTOInput) (*Script, error)
	GetAll(userId int) ([]Script, error)
	GetById(userId int, scriptId int) (*Script, error)
	DeleteById(userId int, scriptId int) error
}

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) ScriptRepository {
	return &postgresRepository{
		db,
	}
}

func (r *postgresRepository) Create(input ScriptDTOInput) (*Script, error) {
	//	  ^^^^^^^^^^^^^^^^^^^^^ É tipo o this de outras linguagens
	//   Isso é o RECEPTOR (receiver)
	return nil, nil
}

func (r *postgresRepository) GetAll(userId int) ([]Script, error) {
	return nil, nil
}

func (r *postgresRepository) GetById(userId int, scriptId int) (*Script, error) {
	return nil, nil
}

func (r *postgresRepository) DeleteById(userId int, scriptId int) error {
	return nil
}
