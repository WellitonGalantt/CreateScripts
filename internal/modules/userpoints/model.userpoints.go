package userpoints

import "time"

//Saldo atual de pontos do usuário (controle de consumo e limites).

type Userpoints struct {
	UserID    string    `json:"user_id"`
	Points    int       `json:"points"`
	UpdatedAt time.Time `json:"updated_at"`
}
