package dto

type CreateRoleDTO struct {
	Name string `json:"name" validate:"required"`
	Key  string `json:"key" validate:"required"`
}

type UpdateRoleDTO struct {
	Name      string `json:"name" validate:"required"`
	Key       string `json:"key" validate:"required"`
	IsDeleted bool   `json:"isDeleted" validate:"boolean"`
}
