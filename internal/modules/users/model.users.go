package users

type Users struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Passord_hash string `json:"password_hash"`
	Role         string `json:"role"`
}
