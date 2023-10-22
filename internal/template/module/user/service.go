package user

import (
	"context"
	"errors"
	"net/http"

	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/constant"
	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/data/sqlxext"
	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/errorext"
	"github.com/tanveerprottoy/stdlib-go-template/internal/template/module/user/dto"
	"github.com/tanveerprottoy/stdlib-go-template/internal/template/module/user/entity"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/timeext"
)

type Service struct {
	repository sqlxext.Repository[entity.User]
}

func NewService(r sqlxext.Repository[entity.User]) *Service {
	s := new(Service)
	s.repository = r
	return s
}

func (s *Service) ReadOneInternal(id string, ctx context.Context) (entity.User, error) {
	return s.repository.ReadOne(id, ctx)
}

func (s *Service) Create(d *dto.CreateUpdateUserDTO, ctx context.Context) (entity.User, errorext.HTTPError) {
	// convert dto to entity
	e := entity.User{}
	e.Name = d.Name
	// b.Role = d.Role
	n := timeext.NowUnixMilli()
	e.CreatedAt = n
	e.UpdatedAt = n
	err := s.repository.Create(e, ctx)
	if err != nil {
		return e, errorext.BuildDBError(err)
	}
	return e, errorext.HTTPError{}
}

func (s *Service) ReadMany(limit, page int, ctx context.Context) (map[string]any, errorext.HTTPError) {
	m := make(map[string]any)
	m["items"] = make([]entity.User, 0)
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

func (s *Service) ReadOne(id string, ctx context.Context) (entity.User, errorext.HTTPError) {
	b, err := s.ReadOneInternal(id, ctx)
	if err != nil {
		return b, errorext.BuildDBError(err)
	}
	return b, errorext.HTTPError{}
}

func (s *Service) Update(id string, d *dto.CreateUpdateUserDTO, ctx context.Context) (entity.User, errorext.HTTPError) {
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

func (s *Service) Delete(id string, ctx context.Context) (entity.User, errorext.HTTPError) {
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
