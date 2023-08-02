package dto

type CreateUpdatePresignedDTO struct {
	Key string `json:"key" validate:"required"`
}
