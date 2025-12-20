package userpoints

import "time"

type GetByIdDTOOutput struct {
	Points    int       `json:"points"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreditValuesDTOInput struct {
	Quantity int               `json:"quantity"`
	UserId   string            `json:"user_id"`
	Reason   TransactionReason `json:"reason"`
}
