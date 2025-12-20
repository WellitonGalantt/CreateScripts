package userpoints

import (
	"scriptmake/internal/apperror"
	"scriptmake/internal/auth"
	"scriptmake/internal/modules/pointstransactions"
	"scriptmake/internal/modules/user"
)

type UserPointsUseCase interface {
	GetById(userId string) (GetByIdDTOOutput, error)
	Debit(body CreditValuesDTOInput) error
	Credit(body CreditValuesDTOInput) error
	GetBalance(userId string) (int, error)
	GetTransactions(userId string) ([]pointstransactions.PointsTransactions, error)
}

type service struct {
	repo       UserPointsRepository
	jwtService *auth.Service
	userRepo   user.UserRepository
}

func NewService(repo UserPointsRepository, jwtService *auth.Service, userRepo user.UserRepository) UserPointsUseCase {
	return &service{
		repo:       repo,
		jwtService: jwtService,
		userRepo:   userRepo,
	}
}

func (s *service) GetById(userId string) (GetByIdDTOOutput, error) {

	existUser, err := s.userRepo.GetById(userId)
	if err != nil {
		return GetByIdDTOOutput{}, err
	}

	if existUser == nil {
		return GetByIdDTOOutput{}, apperror.ErrUserDoesNotExist
	}

	userPoints, err := s.repo.GetById(userId)
	if err != nil {
		return GetByIdDTOOutput{}, err
	}

	pointsOutput := GetByIdDTOOutput{
		Points:    userPoints.Points,
		UpdatedAt: userPoints.UpdatedAt,
	}

	return pointsOutput, nil

}

func (s *service) Debit(body CreditValuesDTOInput) error {

	existUser, err := s.userRepo.GetById(body.UserId)
	if err != nil {
		return err
	}

	if existUser == nil {
		return apperror.ErrUserDoesNotExist
	}

	if !body.Reason.IsValid() || body.Quantity <= 0 {
		return apperror.ErrInvalidInputValues
	}

	err = s.repo.Debit(body.Quantity, body.UserId, body.Reason)
	if err != nil {
		return err
	}

	return nil

}

func (s *service) Credit(body CreditValuesDTOInput) error {

	existUser, err := s.userRepo.GetById(body.UserId)
	if err != nil {
		return err
	}

	if existUser == nil {
		return apperror.ErrUserDoesNotExist
	}

	if !body.Reason.IsValid() || body.Quantity <= 0 {
		return apperror.ErrInvalidInputValues
	}

	err = s.repo.Credit(body.Quantity, body.UserId, body.Reason)
	if err != nil {
		return err
	}

	return nil

}

func (s *service) GetBalance(userId string) (int, error) {

	existUser, err := s.userRepo.GetById(userId)
	if err != nil {
		return 0, err
	}

	if existUser == nil {
		return 0, apperror.ErrUserDoesNotExist
	}

	balance, err := s.repo.GetBalance(userId)
	if err != nil {
		return 0, err
	}

	return balance, nil

}

func (s *service) GetTransactions(userId string) ([]pointstransactions.PointsTransactions, error) {
	existUser, err := s.userRepo.GetById(userId)
	if err != nil {
		return nil, err
	}

	if existUser == nil {
		return nil, apperror.ErrUserDoesNotExist
	}

	pointsTransactions, err := s.repo.GetTransactions(userId)
	if err != nil {
		return nil, err
	}

	return pointsTransactions, nil
}
