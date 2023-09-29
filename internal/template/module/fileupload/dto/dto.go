package dto

type CreatePresignedDTO struct {
	Key string `json:"key" validate:"required"`
}
