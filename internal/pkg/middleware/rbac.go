package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/constant"
	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/rbac"
	"github.com/tanveerprottoy/stdlib-go-template/internal/template/module/auth"

	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/response"
)

// RBAC Role Based Access Control middleware
type RBAC struct {
	Service *auth.Service
}

func NewRBAC(s *auth.Service) *RBAC {
	m := new(RBAC)
	m.Service = s
	return m
}

// AuthUserMiddleWare auth user
func (r *RBAC) AuthRole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("AuthRole.method", request.Method)
		fmt.Println("AuthRole.RequestURI", request.RequestURI)
		fmt.Println("AuthRole.RemoteAddr", request.RemoteAddr)
		fmt.Println("AuthRole.URL", request.URL)
		fmt.Println("AuthRole.URL.RawPath", request.URL.RawPath)
		fmt.Println("AuthRole.URL.Fragment", request.URL.Fragment)
		fmt.Println("AuthRole.URL.Path", request.URL.Path)
		fmt.Println("AuthRole.URL.EscapedPath", request.URL.EscapedPath())
		d := rbac.GetRBAC(request.URL.Path, request.Method)
		if d == nil {
			// could not resolve access control stop the request
			response.RespondError(http.StatusForbidden, constant.Error, constant.Unauthorized, writer)
			return
		}
		d = d.(rbac.RBACModel)
		fmt.Println("GetRBAC: ", d)
		e, err := r.Service.AuthorizeForRole(request)
		if err != nil {
			response.RespondError(http.StatusForbidden, constant.Error, err.Error(), writer)
			return
		}
		ctx := context.WithValue(request.Context(), constant.KeyAuthUser, e)
		_ = request.WithContext(ctx)
		next.ServeHTTP(writer, request)
	})
}
