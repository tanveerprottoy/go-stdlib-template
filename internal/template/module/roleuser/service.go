package comprojroleuser

import (
	"context"
	"net/http"

	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/data"
	"github.com/tanveerprottoy/stdlib-go-template/internal/workersinsights/module/junction/comprojroleuser/dto"
	"github.com/tanveerprottoy/stdlib-go-template/internal/workersinsights/module/junction/comprojroleuser/entity"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/errorext"
)

// Service contains the business logic as well as calls to the
// repository to perform db operations
type Service struct {
	repository *Repository
}

// NewService initializes a new Service
func NewService(r *Repository) *Service {
	return &Service{repository: r}
}

func (s *Service) Create(userID string, d dto.CreateCompanyProjectRoleUserDTO, ctx context.Context) (entity.CompanyProjectUserRole, errorext.HTTPError) {
	var validCompanyID, validProjectID bool
	if d.CompanyID != "" {
		validCompanyID = true
	}
	if d.ProjectID != "" {
		validProjectID = true
	}
	// build entity
	e := entity.CompanyProjectUserRole{
		CompanyID: data.MakeNullString(d.CompanyID, validCompanyID),
		ProjectID: data.MakeNullString(d.ProjectID, validProjectID),
		RoleID:    d.RoleID,
		UserID:    userID,
	}
	q := "INSERT INTO " + tableName + " (company_id, project_id, role_id, user_id) VALUES ($1, $2, $3, $4) RETURNING role_id"
	l, err := s.repository.create(q, e.CompanyID, e.ProjectID, e.RoleID, e.UserID)
	if err != nil {
		return e, errorext.BuildDBError(err)
	}
	e.RoleID = l
	return e, errorext.HTTPError{}
}

func (s *Service) ReadOneRoleForCompanyProjectUser(userID, companyID, projectID string, ctx context.Context) (entity.CompanyProjectUserRole, errorext.HTTPError) {
	var e entity.CompanyProjectUserRole
	row := s.repository.readOneRoleForCompanyProjectUser(userID, companyID, projectID)
	httpErr := data.ScanRow[entity.CompanyProjectUserRole](row, &e, &e.CompanyID, &e.ProjectID, &e.RoleID, &e.UserID)
	if httpErr.Err != nil {
		return e, httpErr
	}
	return e, errorext.HTTPError{}
}

func (s *Service) ReadManyForUser(userID string, ctx context.Context) ([]entity.CompanyProjectUserRole, errorext.HTTPError) {
	d := make([]entity.CompanyProjectUserRole, 0)
	rows, err := s.repository.readManyForUser(userID)
	if err != nil {
		return d, errorext.BuildDBError(err)
	}
	var e entity.CompanyProjectUserRole
	d, httpErr := data.ScanRows(rows, &e, &e.CompanyID, &e.ProjectID, &e.RoleID, &e.UserID)
	if httpErr.Err != nil {
		return d, errorext.HTTPError{Code: http.StatusBadRequest, Err: err}
	}
	return d, errorext.HTTPError{}
}
