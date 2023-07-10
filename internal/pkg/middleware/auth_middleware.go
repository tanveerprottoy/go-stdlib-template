package middleware

import (
	"context"
	"net/http"

	"github.com/tanveerprottoy/go-stdlib-template/internal/app/template/module/auth"
	"github.com/tanveerprottoy/go-stdlib-template/internal/pkg/constant"
	"github.com/tanveerprottoy/go-stdlib-template/pkg/response"
)

type AuthMiddleware struct {
	Service *auth.Service
}

func NewAuthMiddleware(s *auth.Service) *AuthMiddleware {
	m := new(AuthMiddleware)
	m.Service = s
	return m
}

// AuthUserMiddleWare auth user
func (m *AuthMiddleware) AuthUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		e, err := m.Service.Authorize(r)
		if err != nil {
			response.RespondError(http.StatusForbidden, err, w)
			return
		}
		ctx := context.WithValue(r.Context(), constant.KeyAuthUser, e)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}
