package pointstransactions

type PointsTransactions struct {
	ID     int    `json:"id"`
	UserId int    `json:"user_id"`
	Type   string `json:"type"`
	Amount int    `json:"amount"`
	Reason string `json:"reason"`
}
