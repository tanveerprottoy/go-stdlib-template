package dto

type CreateUpdateContentDTO struct {
	Name string `json:"name" validate:"required"`
}
