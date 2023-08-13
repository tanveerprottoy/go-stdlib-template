package auth

import (
	"errors"
	"net/http"

	"github.com/tanveerprottoy/stdlib-go-template/internal/template/module/user"
	"github.com/tanveerprottoy/stdlib-go-template/internal/template/module/user/entity"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/httppkg"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/jwtpkg"
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
	splits, err := httppkg.ParseAuthToken(r)
	if err != nil {
		return e, err
	}
	tokenBody := splits[1]
	claims, err := jwtpkg.VerifyToken(tokenBody)
	if err != nil {
		return e, err
	}
	// find user
	e, err = s.userService.ReadOneInternal(claims.Payload.Id)
	if err != nil {
		return e, err
	}
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
