package mapper

import (
	"expense_tracker/dto"
	"expense_tracker/model"
)

func ToUser(reqDto *dto.SignupRequest) *model.User {
	return &model.User{
		Name:  reqDto.Name,
		Email: reqDto.Email,
	}
}

func ToUserResponse(user *model.User) *dto.UserResponse {
	return &dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
