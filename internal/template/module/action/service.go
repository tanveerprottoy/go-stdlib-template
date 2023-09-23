package action

import (
	"context"
	"net/http"

	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/constant"
	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/data/postgres"
	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/global"
	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/response"
	"github.com/tanveerprottoy/stdlib-go-template/internal/template/module/action/dto"
	"github.com/tanveerprottoy/stdlib-go-template/internal/template/module/action/entity"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/errorext"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/timeext"
)

// Service contains the business logic as well as calls to the
// repository to perform db operations
type Service struct {
	repository *Repository[entity.Action]
}

// NewService initializes a new Service
func NewService(r *Repository[entity.Action]) *Service {
	return &Service{repository: r}
}

// readOneInternal fetches one entity from db
func (s *Service) readOneInternal(id string) (entity.Action, errorext.HTTPError) {
	var e entity.Action
	row := s.repository.ReadOne(id)
	httpErr := postgres.ScanRow[entity.Action](row, &e, &e.ID, &e.Name, &e.Key, &e.IsDeleted, &e.CreatedAt, &e.UpdatedAt)
	return e, httpErr
}

// create defines the business logic for create post request
func (s *Service) create(d dto.CreateActionDTO, ctx context.Context) (entity.Action, errorext.HTTPError) {
	// build entity
	n := timeext.NowUnix()
	e := entity.Action{
		Name:      d.Name,
		Key:       d.Key,
		CreatedAt: n,
		UpdatedAt: n,
	}
	l, err := s.repository.Create(&e)
	if err != nil {
		return e, errorext.BuildDBError(err)
	}
	e.ID = l
	return e, errorext.HTTPError{}
}

func (s *Service) readMany(limit, page int, ctx context.Context) (response.ReadManyResponse[entity.Action], errorext.HTTPError) {
	res := response.ReadManyResponse[entity.Action]{
		Items: make([]entity.Action, 0),
		Limit: limit,
		Page:  page,
	}
	offset := global.CalculateOffset(limit, page)
	rows, err := s.repository.ReadMany(limit, offset)
	if err != nil {
		return res, errorext.BuildDBError(err)
	}
	var e entity.Action
	d, httpErr := postgres.GetEntities(rows, &e, &e.ID, &e.Name, &e.Key, &e.IsDeleted, &e.CreatedAt, &e.UpdatedAt)
	if httpErr.Err != nil {
		return res, httpErr
	}
	res.Items = d
	// fetch total count
	c, httpErr := global.FetchAndScanTotalCount(tableName, "id", "WHERE is_deleted = FALSE", s.repository.DB(), ctx)
	if httpErr.Err != nil {
		return res, httpErr
	}
	res.Total = c.Count
	return res, errorext.HTTPError{}
}

func (s *Service) ReadOne(id string, ctx context.Context) (entity.Action, errorext.HTTPError) {
	b, httpErr := s.readOneInternal(id)
	if httpErr.Err != nil {
		return b, httpErr
	}
	return b, errorext.HTTPError{}
}

func (s *Service) Update(id string, d *dto.UpdateActionDTO, ctx context.Context) (entity.Action, errorext.HTTPError) {
	e, httpErr := s.readOneInternal(id)
	if httpErr.Err != nil {
		return e, httpErr
	}
	e.Name = d.Name
	e.Key = d.Key
	e.IsDeleted = d.IsDeleted
	e.UpdatedAt = timeext.NowUnix()
	rows, err := s.repository.Update(id, &e)
	if err != nil {
		return e, errorext.BuildDBError(err)
	}
	if rows > 0 {
		return e, errorext.HTTPError{}
	}
	return e, errorext.HTTPError{Code: http.StatusBadRequest, Err: errorext.NewError(constant.OperationNotSuccess)}
}

func (s *Service) Delete(id string, ctx context.Context) (entity.Action, errorext.HTTPError) {
	e, httpErr := s.readOneInternal(id)
	if httpErr.Err != nil {
		return e, httpErr
	}
	n := timeext.NowUnix()
	ctx = context.WithValue(ctx, constant.KeyNowMilli, n)
	rows, err := s.repository.Delete(id, ctx)
	if err != nil {
		return e, errorext.BuildDBError(err)
	}
	if rows > 0 {
		e.IsDeleted = true
		e.UpdatedAt = n
		return e, errorext.HTTPError{}
	}
	return e, errorext.HTTPError{Code: http.StatusBadRequest, Err: errorext.NewError(constant.OperationNotSuccess)}
}

func (s *Service) SearchByKey(key string, ctx context.Context) (entity.Action, errorext.HTTPError) {
	var e entity.Action
	row := s.repository.readOneByKey(key)
	httpErr := postgres.ScanRow[entity.Action](row, &e, &e.ID, &e.Name, &e.Key, &e.IsDeleted, &e.CreatedAt, &e.UpdatedAt)
	return e, httpErr
}
