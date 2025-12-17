package user

// DTOS data trasfer obj, objetos especificos para cada operação
type LoginUserDTOInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type RegisterUserDTOInput struct {
	Name            string `json:"name" validate:"required,min=3"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=8"`
}

type UpdateUserDTO struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type ViewProfileDTOOutput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
