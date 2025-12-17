package user

import (
	"database/sql"
	"errors"
)

type UserRepository interface {
	Create(u *User) error
	GetByEmail(email string) (*User, error)
	GetById(id string) (*User, error)
}

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) UserRepository {
	return &postgresRepository{
		db,
	}
}

func (r *postgresRepository) Create(u *User) error {
	query := `INSERT INTO users (name, email, password_hash) VALUES ($1, $2, $3) RETURNING id`

	err := r.db.QueryRow(query, u.Name, u.Email, u.PasswordHash).Scan(&u.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *postgresRepository) GetByEmail(email string) (*User, error) {
	query := `SELECT id, name, email, role, password_hash FROM users WHERE email = $1`

	var u User

	err := r.db.QueryRow(query, email).Scan(&u.ID, &u.Name, &u.Email, &u.Role, &u.PasswordHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &u, nil
}

func (r *postgresRepository) GetById(userID string) (*User, error) {

	query := `SELECT id, name, email, role, password_hash FROM users WHERE id=$1`

	var u User

	err := r.db.QueryRow(query, userID).Scan(&u.ID, &u.Name, &u.Email, &u.Role, &u.PasswordHash)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
