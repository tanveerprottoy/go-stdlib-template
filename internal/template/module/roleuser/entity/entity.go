package entity

import "github.com/tanveerprottoy/stdlib-go-template/internal/pkg/data"

type CompanyProjectUserRole struct {
	CompanyID data.NullString `db:"company_id" json:"companyId"`
	ProjectID data.NullString `db:"project_id" json:"projectId"`
	RoleID    string          `db:"role_id" json:"roleId"`
	UserID    string          `db:"user_id" json:"userId"`
}
