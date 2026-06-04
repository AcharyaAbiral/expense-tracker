package mapper

import (
	"expense_tracker/dto"
	"expense_tracker/model"
)

func ToCategory(reqDto *dto.CategoryRequest) *model.Category {
	return &model.Category{
		Name: reqDto.Name,
	}
}

func ToCategoryResponse(category *model.Category) *dto.CategoryResponse {
	return &dto.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}
}
