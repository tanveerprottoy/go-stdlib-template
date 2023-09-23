package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/constant"
	"github.com/tanveerprottoy/stdlib-go-template/internal/template/module/user"
	"github.com/tanveerprottoy/stdlib-go-template/internal/template/module/user/entity"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/errorext"
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

// Authorize handles authorization for a request
func (s *Service) Authorize(r *http.Request) (*http.Request, error) {
	splits, err := httpext.ParseAuthToken(r)
	if err != nil {
		return r, err
	}
	tokenBody := splits[1]
	jwtToken, err := jwtext.Parse(tokenBody)
	if err != nil {
		return r, err
	}
	// extract the claims
	c := jwtext.ParseClaims(jwtToken)
	if c == nil {
		// handle error
		return r, fmt.Errorf(constant.RBACError)
	}
	// check if id exists
	var id string
	if val, ok := c["id"]; ok {
		// extract the email
		id = val.(string)
	} else {
		return r, fmt.Errorf("missing required data")
	}
	// find user
	_, httpErr := s.userService.ReadOne(id, r.Context())
	if httpErr.Err != nil {
		if httpErr.Code == http.StatusNotFound {
			httpErr.Err = errorext.NewError(constant.RBACError)
		}
		return r, httpErr.Err
	}
	// check with rbac
	// r, d, err := s.authWithRBAC(e.ID, r)
	if err != nil {
		return r, err
	}
	// d[constant.AuthUser] = e
	//ctx := context.WithValue(r.Context(), constant.KeyAuthUser, d)
	return r, nil
	// return r.WithContext(ctx), nil
}

func (s *Service) AuthorizeBasic(r *http.Request) (entity.User, error) {
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

func (s *Service) AuthorizeForRoleBasic(r *http.Request) (entity.User, error) {
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
