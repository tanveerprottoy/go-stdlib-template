package router

import (
	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/constant"
	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/middleware"
	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/router"
	"github.com/tanveerprottoy/stdlib-go-template/internal/template/module/content"

	"github.com/go-chi/chi"
)

func RegisterContentRoutes(router *router.Router, version string, module *content.Module, authMiddleWare *middleware.Auth, rbacMiddleWare *middleware.RBAC) {
	router.Mux.Route(
		constant.ApiPattern+version+constant.ContentsPattern,
		func(r chi.Router) {
			// public routes
			r.Get(constant.RootPattern+"public", module.Handler.Public)
			r.Group(func(r chi.Router) {
				// protected routes
				// r.Use(rbacMiddleWare.AuthRole)
				// r.Use(authMiddleWare.AuthUser)
				r.Get(constant.RootPattern, module.Handler.ReadMany)
				r.Get(constant.RootPattern+"{id}", module.Handler.ReadOne)
				r.Post(constant.RootPattern, module.Handler.Create)
				r.Patch(constant.RootPattern+"{id}", module.Handler.Update)
				r.Delete(constant.RootPattern+"{id}", module.Handler.Delete)
			})
		},
	)
}
