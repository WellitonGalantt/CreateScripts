package subscription

import "time"

type Subscription struct {
	ID                   int       `json:"id"`
	UserId               int       `json:"user_id"`
	PlanId               int       `json:"plan_id"`
	Status               string    `json:"status"`
	Current_period_start time.Time `json:"current_period_start"`
	Current_period_end   time.Time `json:"current_period_end"`
}
