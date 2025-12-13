package users

import (
	"database/sql"
)

type UserRepository interface {
	CreateUser(u *RegisterUserDTO) error
	GetUserByEmail(u *LoginUsersDTO) error
	GetUserById(id int) error
}

type postgresRepository struct {
	db *sql.DB
}

func (r *postgresRepository) NewPostgresRepository(db *sql.DB) UserRepository {
	return &postgresRepository{
		db,
	}
}

func (r *postgresRepository) CreateUser(u *RegisterUserDTO) error {
	return nil
}

func (r *postgresRepository) GetUserByEmail(u *LoginUsersDTO) error {
	return nil
}

func (r *postgresRepository) GetUserById(id int) error {
	return nil
}
