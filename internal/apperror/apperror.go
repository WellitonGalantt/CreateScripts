package apperror

import "errors"

var (
	ErrEmailAlreadyExist   = errors.New("email ja esta cadastrado")
	ErrInvalidCredentialas = errors.New("email ou senha invalido")
	ErrInvalidToken        = errors.New("token invalido")
	ErrInvalidInputValues  = errors.New("email ja esta cadastrado")
)
