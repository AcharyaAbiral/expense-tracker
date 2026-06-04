package service

import (
	"expense_tracker/dto"
	"expense_tracker/mapper"
	"expense_tracker/model"
	"expense_tracker/repository"
	"expense_tracker/util"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) Create(reqDto *dto.SignupRequest) (*model.User, error) {
	passwordHash, err := util.HashPassword(reqDto.Password)
	if err != nil {
		return nil, err
	}
	user := mapper.ToUser(reqDto)
	user.PasswordHash = passwordHash
	if err := s.repo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) FindByID(id uint) (*model.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) FindByEmail(email string) (*model.User, error) {
	return s.repo.FindByEmail(email)
}
