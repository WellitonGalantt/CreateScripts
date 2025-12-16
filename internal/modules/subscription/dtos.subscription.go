package subscription

import "time"

type CreateSubscriptionDtoInput struct {
	UserId int `json:"user_id"`
	PlanId int `json:"plan_id"`
}

type EditSubscriptionDtoInput struct {
	ID                   int        `json:"id"`
	UserId               int        `json:"user_id"`
	PlanId               int        `json:"plan_id"`
	Status               *string    `json:"status,omitempty"`
	Current_period_start *time.Time `json:"current_period_start,omitempty"`
	Current_period_end   *time.Time `json:"current_period_end,omitempty"`
}
