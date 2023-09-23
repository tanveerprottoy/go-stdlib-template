package dto

type CreateCompanyProjectRoleUserDTO struct {
	CompanyID string `json:"companyId" validate:"omitempty"`
	ProjectID string `json:"projectId" validate:"omitempty"`
	RoleID    string `json:"roleId" validate:"required"`
}

type UpdateCompanyProjectRoleUserDTO struct {
	CompanyID string `json:"companyId" validate:"omitempty"`
	ProjectID string `json:"projectId" validate:"omitempty"`
	RoleID    string `json:"roleId" validate:"required"`
}
