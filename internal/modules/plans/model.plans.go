package plans

type Plans struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Monthly_points int    `json:"monthly_points"`
	Is_unlimited   bool   `json:"is_unlimited"`
	Price_cents    int    `json:"price_cents"`
}
