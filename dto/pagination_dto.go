package dto

type PaginationInput struct {
	Page  int `query:"page"`
	Limit int `query:"limit"`
}

type PaginatedResponse[T any] struct {
	Data  []T   `json:"data"`
	Page  int   `json:"page"`
	Limit int   `json:"limit"`
	Total int64 `json:"total"`
}
