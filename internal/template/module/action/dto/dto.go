package dto

type CreateActionDTO struct {
	Name string `json:"name" validate:"required"`
	Key  string `json:"key" validate:"required"`
}

type UpdateActionDTO struct {
	Name      string `json:"name" validate:"required"`
	Key       string `json:"key" validate:"required"`
	IsDeleted bool   `json:"isDeleted" validate:"boolean"`
}
