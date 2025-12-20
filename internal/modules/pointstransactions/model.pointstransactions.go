package pointstransactions

import "time"

type PointsTransactions struct {
	ID        string    `json:"id"`
	UserId    string    `json:"user_id"`
	Type      string    `json:"type"`
	Amount    int       `json:"amount"`
	Reason    string    `json:"reason"`
	CreatedAt time.Time `json:"created_at"`
}
