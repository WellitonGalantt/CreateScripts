package plans

type GetPlanOutput struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	MonthlyPoints int    `json:"monthly_points"`
	IsUnlimited   bool   `json:"is_unlimited"`
	PriceCents    int    `json:"price_cents"`
}
