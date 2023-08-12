package middleware

import (
	"fmt"
	"net/http"

	"github.com/tanveerprottoy/stdlib-go-template/internal/template/module/auth"
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
		/* e, err := r.Service.Authorize(r)
		if err != nil {
			response.RespondError(http.StatusForbidden, err, w)
			return
		}
		ctx := context.WithValue(r.Context(), constant.KeyAuthUser, e)
		req := r.WithContext(ctx) */
		next.ServeHTTP(writer, request)
	})
}
