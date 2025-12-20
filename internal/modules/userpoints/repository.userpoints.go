package userpoints

import (
	"database/sql"
	"errors"
	"scriptmake/internal/apperror"
	"scriptmake/internal/modules/pointstransactions"
)

type UserPointsRepository interface {
	GetById(userId string) (*Userpoints, error)
	Credit(quantity int, userId string, reason TransactionReason) error
	Debit(quantity int, userId string, reason TransactionReason) error
	GetBalance(userId string) (int, error)
	GetTransactions(userId string) ([]pointstransactions.PointsTransactions, error)
}

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) UserPointsRepository {
	return &postgresRepository{
		db,
	}
}

func (r *postgresRepository) GetById(userId string) (*Userpoints, error) {
	query := `SELECT user_id, points, updated_at FROM user_points WHERE user_id = $1`

	var userPoints Userpoints
	err := r.db.QueryRow(query, userId).Scan(&userPoints.UserID, &userPoints.Points, &userPoints.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &userPoints, nil
}

func (r *postgresRepository) Credit(quantity int, userId string, reason TransactionReason) error {
	// usando transaction para agrupar varias operacoes sql
	// Transactions voce pode ir agrupando operacoes sql e aplica todas de uma vez so no final
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	// Caso nao tenha feito um return tx.Commit() ele desfaz as operacoes
	defer tx.Rollback()

	// OON CONFLICT (user_id) serve para que se ja existir registro com o mesmo user_id ele fazer o update
	updateQuery := `INSERT INTO user_points (user_id, points) 
	VALUES ($1, $2)
	ON CONFLICT (user_id)
	DO UPDATE SET points = user_points.points + $2, updated_at = NOW()`

	_, err = tx.Exec(updateQuery, userId, quantity)
	if err != nil {
		return err
	}

	// Registrar a transacao
	transactionQuery := `
	INSERT INTO points_transactions (user_id, type, amount, reason)
	VALUES ($1, $2, $3, $4)`

	_, err = tx.Exec(transactionQuery, userId, Credit, quantity, reason)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *postgresRepository) Debit(quantity int, userId string, reason TransactionReason) error {

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	var currentPoints int
	checkQuery := `SELECT points FROM user_points WHERE user_id = $1`
	err = tx.QueryRow(checkQuery, userId).Scan(&currentPoints)
	if err != nil {
		return err
	}

	if currentPoints < quantity {
		return apperror.ErrInsufficientPoints
	}

	debitQuery := `
	UPDATE user_points 
	SET points = points - $1, updated_at = NOW()
	WHERE user_id = $2
	`
	_, err = tx.Exec(debitQuery, quantity, userId)
	if err != nil {
		return err
	}

	transactionQuery := `
	INSERT INTO points_transactions (user_id, type, amount, reason)
	VALUES ($1, $2, $3, $4)`

	_, err = tx.Exec(transactionQuery, userId, Debit, quantity, reason)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *postgresRepository) GetBalance(userId string) (int, error) {
	query := `SELECT COALESCE(points, 0) FROM user_points WHERE user_id = $1`

	var points int
	err := r.db.QueryRow(query, userId).Scan(&points)
	if err != nil {
		//Se nao encontrar o registro é poruqe o usuario nao possui pontos ainda, nao é um erro exatamente
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}
		return 0, err
	}

	return points, nil

}

func (r *postgresRepository) GetTransactions(userId string) ([]pointstransactions.PointsTransactions, error) {
	query := `SELECT id, user_id, type, amount, reason, created_at FROM points_transactions WHERE user_id = $1`

	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var transactions []pointstransactions.PointsTransactions
	for rows.Next() {
		var t pointstransactions.PointsTransactions
		err := rows.Scan(&t.ID, &t.UserId, &t.Type, &t.Amount, &t.Reason, &t.CreatedAt)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}

	return transactions, nil
}
