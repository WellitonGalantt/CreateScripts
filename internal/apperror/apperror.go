package apperror

import "errors"

var (
	ErrEmailAlreadyExist  = errors.New("email ja esta cadastrado")
	ErrInvalidCredentials = errors.New("email ou senha invalido")
	ErrInvalidToken       = errors.New("token invalido")
	ErrInvalidInputValues = errors.New("formato email invalido")
)
