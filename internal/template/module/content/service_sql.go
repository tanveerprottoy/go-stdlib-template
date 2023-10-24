package content

import (
	"context"
	"errors"
	"net/http"

	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/constant"
	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/data/postgres"
	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/errorext"
	"github.com/tanveerprottoy/stdlib-go-template/internal/template/module/content/dto"
	"github.com/tanveerprottoy/stdlib-go-template/internal/template/module/content/entity"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/timeext"
)

// ServiceSQL contains the business logic as well as calls to the
// repository to perform db operations
type ServiceSQL struct {
	repository postgres.Repository[entity.Content]
}

// NewService initializes a new ServiceSQL
func NewServiceSQL(r postgres.Repository[entity.Content]) *ServiceSQL {
	return &ServiceSQL{repository: r}
}

func (s *ServiceSQL) readOneInternal(id string, ctx context.Context) (entity.Content, errorext.HTTPError) {
	var e entity.Content
	row := s.repository.ReadOne(id, ctx)
	err := row.Err()
	if err != nil {
		return e, errorext.HTTPError{Code: http.StatusInternalServerError, Err: err}
	}
	httpErr := postgres.ScanRow[entity.Content](row, &e, &e.ID, &e.Name, &e.CreatedAt, &e.UpdatedAt)
	return e, httpErr
}

// Create defines the business logic for create post request
func (s *ServiceSQL) Create(d dto.CreateUpdateContentDTO, ctx context.Context) (entity.Content, errorext.HTTPError) {
	// build entity
	n := timeext.NowUnixMilli()
	e := entity.Content{
		Name:      d.Name,
		CreatedAt: n,
		UpdatedAt: n,
	}
	l, err := s.repository.Create(e, ctx)
	if err != nil {
		return e, errorext.BuildDBError(err)
	}
	e.ID = l
	return e, errorext.HTTPError{}
}

func (s *ServiceSQL) ReadMany(limit, page int, ctx context.Context) (map[string]any, errorext.HTTPError) {
	m := make(map[string]any)
	m["items"] = make([]entity.Content, 0)
	m["limit"] = limit
	m["page"] = page
	offset := limit * (page - 1)
	rows, err := s.repository.ReadMany(limit, offset, ctx)
	if err != nil {
		return m, errorext.BuildDBError(err)
	}
	defer rows.Close()
	var e entity.Content
	d, httpErr := postgres.ScanRows(rows, &e, &e.ID, &e.Name, &e.CreatedAt, &e.UpdatedAt)
	if httpErr.Err != nil {
		return m, errorext.BuildDBError(err)
	}
	m["items"] = d
	return m, errorext.HTTPError{}
}

func (s *ServiceSQL) ReadOne(id string, ctx context.Context) (entity.Content, errorext.HTTPError) {
	b, httpErr := s.readOneInternal(id, ctx)
	if httpErr.Err != nil {
		return b, errorext.BuildDBError(httpErr.Err)
	}
	return b, errorext.HTTPError{}
}

func (s *ServiceSQL) Update(id string, d *dto.CreateUpdateContentDTO, ctx context.Context) (entity.Content, errorext.HTTPError) {
	b, httpErr := s.readOneInternal(id, ctx)
	if httpErr.Err != nil {
		return b, errorext.BuildDBError(httpErr.Err)
	}
	b.Name = d.Name
	b.UpdatedAt = timeext.NowUnixMilli()
	rows, err := s.repository.Update(id, b, ctx)
	if err != nil {
		return b, errorext.BuildDBError(err)
	}
	if rows > 0 {
		return b, errorext.HTTPError{}
	}
	return b, errorext.HTTPError{Code: http.StatusBadRequest, Err: errors.New(constant.OperationNotSuccess)}
}

func (s *ServiceSQL) Delete(id string, ctx context.Context) (entity.Content, errorext.HTTPError) {
	b, httpErr := s.readOneInternal(id, ctx)
	if httpErr.Err != nil {
		return b, errorext.BuildDBError(httpErr.Err)
	}
	rows, err := s.repository.Delete(id, ctx)
	if err != nil {
		return b, errorext.BuildDBError(err)
	}
	if rows > 0 {
		return b, errorext.HTTPError{}
	}
	return b, errorext.HTTPError{Code: http.StatusBadRequest, Err: errors.New(constant.OperationNotSuccess)}
}
