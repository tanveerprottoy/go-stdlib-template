package content

import (
	"context"
	"net/http"

	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/constant"
	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/data/postgres"
	"github.com/tanveerprottoy/stdlib-go-template/internal/template/module/content/dto"
	"github.com/tanveerprottoy/stdlib-go-template/internal/template/module/content/entity"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/errorpkg"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/timepkg"
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
	return postgres.GetEntity[entity.Content](row, e, e.Id, e.Name, e.CreatedAt, e.UpdatedAt)
}

// Create defines the business logic for create post request
func (s *ServiceSQL) Create(d dto.CreateUpdateContentDTO, ctx context.Context) (entity.Content, *errorpkg.HTTPError) {
	// build entity
	n := timepkg.NowUnixMilli()
	e := entity.Content{
		Name:      d.Name,
		CreatedAt: n,
		UpdatedAt: n,
	}
	err := s.repository.Create(&e)
	if err != nil {
		return e, errorpkg.MakeDBError(err)
	}
	return e, nil
}

func (s *ServiceSQL) ReadMany(limit, page int, ctx context.Context) (map[string]any, *errorpkg.HTTPError) {
	m := make(map[string]any)
	m["items"] = make([]entity.Content, 0)
	m["limit"] = limit
	m["page"] = page
	offset := limit * (page - 1)
	d, err := s.repository.ReadMany(limit, offset)
	if err != nil {
		return m, errorpkg.MakeDBError(err)
	}
	m["items"] = d
	return m, nil
}

func (s *ServiceSQL) ReadOne(id string, ctx context.Context) (entity.Content, *errorpkg.HTTPError) {
	b, err := s.readOneInternal(id)
	if err != nil {
		return b, errorpkg.MakeDBError(err)
	}
	return b, nil
}

func (s *ServiceSQL) Update(id string, d *dto.CreateUpdateContentDTO, ctx context.Context) (entity.Content, *errorpkg.HTTPError) {
	b, err := s.readOneInternal(id)
	if err != nil {
		return b, errorpkg.MakeDBError(err)
	}
	b.Name = d.Name
	b.UpdatedAt = timepkg.NowUnixMilli()
	rows, err := s.repository.Update(id, &b)
	if err != nil {
		return b, errorpkg.MakeDBError(err)
	}
	if rows > 0 {
		return b, nil
	}
	return b, &errorpkg.HTTPError{Code: http.StatusBadRequest, Err: errorpkg.NewError(constant.OperationNotSuccess)}
}

func (s *ServiceSQL) Delete(id string, ctx context.Context) (entity.Content, *errorpkg.HTTPError) {
	b, err := s.readOneInternal(id)
	if err != nil {
		return b, errorpkg.MakeDBError(err)
	}
	rows, err := s.repository.Delete(id)
	if err != nil {
		return b, errorpkg.MakeDBError(err)
	}
	if rows > 0 {
		return b, nil
	}
	return b, &errorpkg.HTTPError{Code: http.StatusBadRequest, Err: errorpkg.NewError(constant.OperationNotSuccess)}
}
