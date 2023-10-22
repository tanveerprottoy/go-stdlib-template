package content

import (
	"context"
	"errors"
	"net/http"

	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/constant"
	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/data/sqlxext"
	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/errorext"
	"github.com/tanveerprottoy/stdlib-go-template/internal/template/module/content/dto"
	"github.com/tanveerprottoy/stdlib-go-template/internal/template/module/content/entity"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/timeext"
)

type Service struct {
	repository sqlxext.Repository[entity.Content]
}

func NewService(r sqlxext.Repository[entity.Content]) *Service {
	s := new(Service)
	s.repository = r
	return s
}

func (s *Service) ReadOneInternal(id string, ctx context.Context) (entity.Content, error) {
	return s.repository.ReadOne(id, ctx)
}

func (s *Service) Create(d dto.CreateUpdateContentDTO, ctx context.Context) (entity.Content, errorext.HTTPError) {
	// convert dto to entity
	b := entity.Content{}
	b.Name = d.Name
	n := timeext.NowUnixMilli()
	b.CreatedAt = n
	b.UpdatedAt = n
	err := s.repository.Create(b, ctx)
	if err != nil {
		return b, errorext.BuildDBError(err)
	}
	return b, errorext.HTTPError{}
}

func (s *Service) ReadMany(limit, page int, ctx context.Context) (map[string]any, errorext.HTTPError) {
	m := make(map[string]any)
	m["items"] = make([]entity.Content, 0)
	m["limit"] = limit
	m["page"] = page
	offset := limit * (page - 1)
	d, err := s.repository.ReadMany(limit, offset, ctx)
	if err != nil {
		return m, errorext.BuildDBError(err)
	}
	m["items"] = d
	return m, errorext.HTTPError{}
}

func (s *Service) ReadOne(id string, ctx context.Context) (entity.Content, errorext.HTTPError) {
	b, err := s.ReadOneInternal(id, ctx)
	if err != nil {
		return b, errorext.BuildDBError(err)
	}
	return b, errorext.HTTPError{}
}

func (s *Service) Update(id string, d dto.CreateUpdateContentDTO, ctx context.Context) (entity.Content, errorext.HTTPError) {
	b, err := s.ReadOneInternal(id, ctx)
	if err != nil {
		return b, errorext.BuildDBError(err)
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

func (s *Service) Delete(id string, ctx context.Context) (entity.Content, errorext.HTTPError) {
	b, err := s.ReadOneInternal(id, ctx)
	if err != nil {
		return b, errorext.BuildDBError(err)
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
