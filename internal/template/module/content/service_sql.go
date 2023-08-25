package content

import (
	"context"
	"net/http"

	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/constant"
	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/data/postgres"
	"github.com/tanveerprottoy/stdlib-go-template/internal/template/module/content/dto"
	"github.com/tanveerprottoy/stdlib-go-template/internal/template/module/content/entity"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/errorext"
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

func (s *ServiceSQL) readOneInternal(id string) (entity.Content, error) {
	var e entity.Content
	row := s.repository.ReadOne(id)
	return postgres.GetEntity[entity.Content](row, &e, &e.Id, &e.Name, &e.CreatedAt, &e.UpdatedAt)
}

// Create defines the business logic for create post request
func (s *ServiceSQL) Create(d dto.CreateUpdateContentDTO, ctx context.Context) (entity.Content, *errorext.HTTPError) {
	// build entity
	n := timeext.NowUnixMilli()
	e := entity.Content{
		Name:      d.Name,
		CreatedAt: n,
		UpdatedAt: n,
	}
	l, err := s.repository.Create(&e)
	if err != nil {
		return e, errorext.BuildDBError(err)
	}
	e.Id = l
	return e, nil
}

func (s *ServiceSQL) ReadMany(limit, page int, ctx context.Context) (map[string]any, *errorext.HTTPError) {
	m := make(map[string]any)
	m["items"] = make([]entity.Content, 0)
	m["limit"] = limit
	m["page"] = page
	offset := limit * (page - 1)
	d, err := s.repository.ReadMany(limit, offset)
	if err != nil {
		return m, errorext.BuildDBError(err)
	}
	m["items"] = d
	return m, nil
}

func (s *ServiceSQL) ReadOne(id string, ctx context.Context) (entity.Content, *errorext.HTTPError) {
	b, err := s.readOneInternal(id)
	if err != nil {
		return b, errorext.BuildDBError(err)
	}
	return b, nil
}

func (s *ServiceSQL) Update(id string, d *dto.CreateUpdateContentDTO, ctx context.Context) (entity.Content, *errorext.HTTPError) {
	b, err := s.readOneInternal(id)
	if err != nil {
		return b, errorext.BuildDBError(err)
	}
	b.Name = d.Name
	b.UpdatedAt = timeext.NowUnixMilli()
	rows, err := s.repository.Update(id, &b)
	if err != nil {
		return b, errorext.BuildDBError(err)
	}
	if rows > 0 {
		return b, nil
	}
	return b, &errorext.HTTPError{Code: http.StatusBadRequest, Err: errorext.NewError(constant.OperationNotSuccess)}
}

func (s *ServiceSQL) Delete(id string, ctx context.Context) (entity.Content, *errorext.HTTPError) {
	b, err := s.readOneInternal(id)
	if err != nil {
		return b, errorext.BuildDBError(err)
	}
	rows, err := s.repository.Delete(id)
	if err != nil {
		return b, errorext.BuildDBError(err)
	}
	if rows > 0 {
		return b, nil
	}
	return b, &errorext.HTTPError{Code: http.StatusBadRequest, Err: errorext.NewError(constant.OperationNotSuccess)}
}
