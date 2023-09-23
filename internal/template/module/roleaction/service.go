package roleaction

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/data/postgres"
	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/global"
	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/response"
	"github.com/tanveerprottoy/stdlib-go-template/internal/template/module/roleaction/dto"
	"github.com/tanveerprottoy/stdlib-go-template/internal/template/module/roleaction/entity"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/errorext"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/sliceext"
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

func (s *Service) Create(roleID string, d dto.CreateRoleActionDTO, ctx context.Context) (entity.RoleAction, errorext.HTTPError) {
	// build entity
	e := entity.RoleAction{
		RoleID:    roleID,
		ActionIDs: d.ActionIDs,
	}
	jsonArr, err := json.Marshal(e.ActionIDs)
	if err != nil {
		return e, errorext.HTTPError{Code: http.StatusInternalServerError, Err: err}
	}
	q := "INSERT INTO " + tableName + " (role_id, action_ids) VALUES ($1, $2, $3) RETURNING role_id"
	l, err := s.repository.Create(q, e.RoleID, jsonArr)
	if err != nil {
		return e, errorext.BuildDBError(err)
	}
	e.RoleID = l
	return e, errorext.HTTPError{}
}

func (s *Service) ReadManyForRole(roleID string, limit, page int, ctx context.Context) (response.ReadManyResponse[entity.RoleAction], errorext.HTTPError) {
	res := response.ReadManyResponse[entity.RoleAction]{
		Items: make([]entity.RoleAction, 0),
		Limit: limit,
		Page:  page,
	}
	offset := global.CalculateOffset(limit, page)
	rows, err := s.repository.ReadManyForRole(roleID, limit, offset)
	if err != nil {
		return res, errorext.BuildDBError(err)
	}
	var e entity.RoleAction
	d, httpErr := postgres.ScanRows(rows, &e, &e.RoleID, &e.ActionIDs)
	if httpErr.Err != nil {
		return res, errorext.HTTPError{Code: http.StatusBadRequest, Err: err}
	}
	res.Items = d
	// fetch total count
	c, httpErr := global.FetchAndScanTotalCount(tableName, "id", "WHERE is_delete = FALSE", s.repository.db, ctx)
	if httpErr.Err != nil {
		return res, httpErr
	}
	res.Total = c.Count
	return res, errorext.HTTPError{}
}

func (s *Service) ReadManyActionsForRole(roleID string, ctx context.Context) ([]string, errorext.HTTPError) {
	d := make([]string, 0)
	rows, err := s.repository.ReadManyActionsForRole(roleID)
	if err != nil {
		return d, errorext.BuildDBError(err)
	}
	var e postgres.Json2dArray
	httpErr := postgres.ScanRowsBasic(rows, &e)
	if httpErr.Err != nil {
		return d, httpErr
	}
	// flatten the 2d array to 1d
	d = sliceext.Flatten[string](e)
	return d, errorext.HTTPError{}
}

func (s *Service) ReadManyActionsForRole1(roleID string, ctx context.Context) (postgres.JsonStringArray, errorext.HTTPError) {
	var e postgres.JsonStringArray
	rows, err := s.repository.ReadManyActionsForRole1(roleID)
	if err != nil {
		return e, errorext.BuildDBError(err)
	}
	httpErr := postgres.ScanRowsBasic(rows, &e)
	if httpErr.Err != nil {
		return e, httpErr
	}
	return e, errorext.HTTPError{}
}
