package users

// DTOS data trasfer obj, objetos especificos para cada operação
type LoginUsersDTO struct {
	Email            string `json:"email"`
	Passord_hash     string `json:"password_hash"`
	Confirm_password string `json:"confirm_password"`
}

type RegisterUserDTO struct {
	Name             string `json:"name"`
	Email            string `json:"email"`
	Passord          string `json:"password"`
	Confirm_password string `json:"confirm_password"`
}
