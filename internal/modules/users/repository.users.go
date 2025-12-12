package users

import (
	"database/sql"
	"scriptmake/internal/modules/users"
)

type Repository interface {
	CreateUser(u *users.RegisterUserDTO)
	FindUserByEmail(u *users.LoginUsersDTO)
}

type postgresRepository struct {
	db *sql.DB
}

func (r *postgresRepository) NewPostgresRepository(db *sql.DB) *postgresRepository {
	return &postgresRepository{
		db,
	}
}
