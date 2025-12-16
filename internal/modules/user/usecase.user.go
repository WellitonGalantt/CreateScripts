package user

import (
	"scriptmake/internal/apperror"
	"scriptmake/internal/auth"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	Register(input RegisterUserDTOInput) error
	Login(input LoginUserDTOInput) (string, error)
	Viewprofile(userid int) (ViewProfileDTOOutput, error)
}

type service struct {
	repo       UserRepository
	jwtService *auth.Service
}

func NewService(repo UserRepository, jwt *auth.Service) UserUseCase {
	return &service{
		repo:       repo,
		jwtService: jwt,
	}
}

func (s *service) Register(input RegisterUserDTOInput) error {
	name := strings.TrimSpace(input.Name)
	email := strings.TrimSpace(input.Email)

	existUser, err := s.repo.GetByEmail(email)
	if err != nil {
		return err
	}

	// Se existe ja um usuario com esse email
	if existUser != nil {
		return apperror.ErrEmailAlreadyExist
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Passord), 12)
	if err != nil {
		return err
	}

	user := &User{
		Name:        name,
		Email:       email,
		PassordHash: string(hash),
	}

	err = s.repo.Create(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Login(input LoginUserDTOInput) (string, error) {
	email := strings.TrimSpace(input.Email)

	existUser, err := s.repo.GetByEmail(email)
	if err != nil {
		return "", err
	}

	if existUser == nil {
		return "", apperror.ErrInvalidCredentials
	}

	//Verificando senha:
	err = bcrypt.CompareHashAndPassword(
		[]byte(existUser.PassordHash),
		[]byte(input.Passord),
	)

	if err != nil {
		return "", apperror.ErrInvalidCredentials
	}

	token, err := s.jwtService.GenerateToken(existUser.ID, existUser.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *service) Viewprofile(userID int) (ViewProfileDTOOutput, error) {
	user, err := s.repo.GetById(userID)
	if err != nil {
		return ViewProfileDTOOutput{}, nil
	}

	viewUser := ViewProfileDTOOutput{
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}

	return viewUser, nil
}
