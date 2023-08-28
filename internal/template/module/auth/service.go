package auth

import (
	"context"
	"errors"
	"net/http"

	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/constant"
	"github.com/tanveerprottoy/stdlib-go-template/internal/template/module/user"
	"github.com/tanveerprottoy/stdlib-go-template/internal/template/module/user/entity"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/httpext"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/jwtext"
)

type Service struct {
	userService *user.Service
}

func NewService(userService *user.Service) *Service {
	s := new(Service)
	s.userService = userService
	return s
}

func (s *Service) Authorize(r *http.Request) (entity.User, error) {
	var e entity.User
	splits, err := httpext.ParseAuthToken(r)
	if err != nil {
		return e, err
	}
	tokenBody := splits[1]
	claims, err := jwtext.VerifyToken1(tokenBody)
	if err != nil {
		return e, err
	}
	// find user
	e, err = s.userService.ReadOneInternal(claims.Payload.Id)
	if err != nil {
		return e, err
	}
	ctx := context.WithValue(r.Context(), constant.KeyAuthUser, e)
	_ = r.WithContext(ctx)
	return e, nil
}

func (s *Service) AuthorizeForRole(r *http.Request) (entity.User, error) {
	var e entity.User
	role := r.Header.Get("role")
	if role == "" {
		// role is missing
		return e, errors.New("role is missing")
	}
	// find user
	e, err := s.userService.ReadOneInternal(r.Header.Get("id"))
	if err != nil {
		return e, err
	}
	return e, nil
}
