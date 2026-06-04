package service

import (
	"expense_tracker/dto"
	"expense_tracker/model"
	"expense_tracker/util"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	userService *UserService
	jwtSecret   string
}

func NewAuthService(userService *UserService, jwtSecret string) *AuthService {
	return &AuthService{userService, jwtSecret}
}

func (s *AuthService) Login(reqDto *dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := s.userService.FindByEmail(reqDto.Email)
	if err != nil {
		return nil, err
	}

	if err := util.CompareHashAndPassword(user.PasswordHash, reqDto.Password); err != nil {
		return nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.ID,
		"exp":    time.Now().Add(30 * time.Minute).Unix(),
	})

	t, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{Token: t}, nil
}

func (s *AuthService) Signup(reqDto *dto.SignupRequest) (*model.User, error) { //todo: return auth response instead
	return s.userService.Create(reqDto)
}
