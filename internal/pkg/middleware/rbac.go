package middleware

import (
	"net/http"

	"github.com/tanveerprottoy/stdlib-go-template/internal/template/module/auth"
)

// RBAC Role Based Access Control
type RBAC struct {
	Service *auth.Service
}

func NewRBAC(s *auth.Service) *Auth {
	m := new(Auth)
	m.Service = s
	return m
}

// AuthUserMiddleWare auth user
func (r *RBAC) AuthUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
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
