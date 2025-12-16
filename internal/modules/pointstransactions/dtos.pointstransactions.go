package pointstransactions

type CreatePointsTransactionsDtoInput struct {
	UserId int    `json:"user_id"`
	Type   string `json:"type" validate:"required,oneof=credit debit"`
	Amount int    `json:"amount"`
	Reason string `json:"reason"`
}

type CreatePointsTransactionsDtoOutput struct {
	ID     int    `json:"id"`
	UserId int    `json:"user_id"`
	Type   string `json:"type"`
	Amount int    `json:"amount"`
}
