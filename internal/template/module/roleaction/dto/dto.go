package dto

type CreateRoleActionDTO struct {
	ActionIDs []string `json:"actionIds" validate:"required"`
}

type UpdateRoleActionDTO struct {
	ActionIDs []string `json:"actionIds" validate:"required"`
}
