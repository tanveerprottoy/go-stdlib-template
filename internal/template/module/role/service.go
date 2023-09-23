package role

import (
	"context"
	"net/http"

	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/constant"
	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/data"
	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/global"
	"github.com/tanveerprottoy/stdlib-go-template/internal/workersinsights/module/junction/rolemodaction"
	rolemodactiondto "github.com/tanveerprottoy/stdlib-go-template/internal/workersinsights/module/junction/rolemodaction/dto"
	rolemodactionentity "github.com/tanveerprottoy/stdlib-go-template/internal/workersinsights/module/junction/rolemodaction/entity"
	"github.com/tanveerprottoy/stdlib-go-template/internal/workersinsights/module/role/dto"
	"github.com/tanveerprottoy/stdlib-go-template/internal/workersinsights/module/role/entity"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/errorext"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/timeext"
)

// Service contains the business logic as well as calls to the
// repository to perform db operations
type Service struct {
	repository data.Repository[entity.Role]
	// needs rolemodaction roleModActionService
	roleModActionService *rolemodaction.Service
}

// NewService initializes a new Service
func NewService(r data.Repository[entity.Role], s *rolemodaction.Service) *Service {
	return &Service{repository: r, roleModActionService: s}
}

// readOneInternal fetches one entity from db
func (s *Service) readOneInternal(id string) (entity.Role, errorext.HTTPError) {
	var e entity.Role
	row := s.repository.ReadOne(id)
	return data.GetEntity1[entity.Role](row, &e, &e.ID, &e.Name, &e.Key, &e.IsDeleted, &e.CreatedAt, &e.UpdatedAt)
}

// create defines the business logic for create post request
func (s *Service) create(d dto.CreateRoleDTO, ctx context.Context) (entity.Role, errorext.HTTPError) {
	// build entity
	n := timeext.NowUnix()
	e := entity.Role{Name: d.Name, Key: d.Key, CreatedAt: n, UpdatedAt: n}
	l, err := s.repository.Create(&e)
	if err != nil {
		return e, errorext.BuildDBError(err)
	}
	e.ID = l
	return e, errorext.HTTPError{}
}

func (s *Service) readMany(limit, page int, ctx context.Context) (global.ReadManyResponse[entity.Role], errorext.HTTPError) {
	res := global.ReadManyResponse[entity.Role]{
		Items: make([]entity.Role, 0),
		Limit: limit,
		Page:  page,
	}
	offset := global.CalculateOffset(limit, page)
	rows, err := s.repository.ReadMany(limit, offset)
	if err != nil {
		return res, errorext.BuildDBError(err)
	}
	var e entity.Role
	d, httpErr := data.GetEntities(rows, &e, &e.ID, &e.Name, &e.Key, &e.IsDeleted, &e.CreatedAt, &e.UpdatedAt)
	if httpErr.Err != nil {
		return res, errorext.HTTPError{Code: http.StatusBadRequest, Err: err}
	}
	res.Items = d
	// fetch total count
	c, httpErr := global.FetchAndScanTotalCount(tableName, constant.ColID, constant.IsDeletedClause, s.repository.DB(), ctx)
	if httpErr.Err != nil {
		return res, httpErr
	}
	res.Total = c.Count
	return res, errorext.HTTPError{}
}

func (s *Service) ReadOne(id string, ctx context.Context) (entity.Role, errorext.HTTPError) {
	b, httpErr := s.readOneInternal(id)
	if httpErr.Err != nil {
		return b, httpErr
	}
	return b, errorext.HTTPError{}
}

func (s *Service) update(id string, d *dto.UpdateRoleDTO, ctx context.Context) (entity.Role, errorext.HTTPError) {
	b, httpErr := s.readOneInternal(id)
	if httpErr.Err != nil {
		return b, httpErr
	}
	b.Name = d.Name
	b.Key = d.Key
	b.UpdatedAt = timeext.NowUnix()
	rows, err := s.repository.Update(id, &b)
	if err != nil {
		return b, errorext.BuildDBError(err)
	}
	if rows > 0 {
		return b, errorext.HTTPError{}
	}
	return b, errorext.HTTPError{Code: http.StatusBadRequest, Err: errorext.NewError(constant.OperationNotSuccess)}
}

func (s *Service) delete(id string, ctx context.Context) (entity.Role, errorext.HTTPError) {
	b, httpErr := s.readOneInternal(id)
	if httpErr.Err != nil {
		return b, httpErr
	}
	rows, err := s.repository.Delete(id, nil)
	if err != nil {
		return b, errorext.BuildDBError(err)
	}
	if rows > 0 {
		return b, errorext.HTTPError{}
	}
	return b, errorext.HTTPError{Code: http.StatusBadRequest, Err: errorext.NewError(constant.OperationNotSuccess)}
}

// Related Data Methods
func (s *Service) CreateRoleModuleAction(roleID string, d rolemodactiondto.CreateRoleModuleActionDTO, ctx context.Context) (rolemodactionentity.RoleModuleAction, errorext.HTTPError) {
	return s.roleModActionService.Create(roleID, d, ctx)
}

// Related Data Methods
func (s *Service) ReadManyModulesActionsForRole(roleID string, limit, page int, ctx context.Context) (global.ReadManyResponse[rolemodactionentity.RoleModuleAction], errorext.HTTPError) {
	return s.roleModActionService.ReadManyForRole(roleID, limit, page, ctx)
}

// Related Data Methods
func (s *Service) ReadManyActionsForRole(roleID string, ctx context.Context) ([]string, errorext.HTTPError) {
	return s.roleModActionService.ReadManyActionsForRole(roleID, ctx)
}

// Related Data Methods
func (s *Service) ReadManyActionsForRole1(roleID string, ctx context.Context) (data.JsonStringArray, errorext.HTTPError) {
	return s.roleModActionService.ReadManyActionsForRole1(roleID, ctx)
}
