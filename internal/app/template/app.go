package template

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/tanveerprottoy/stdlib-go-template/internal/app/template/module/auth"
	"github.com/tanveerprottoy/stdlib-go-template/internal/app/template/module/user"
	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/constant"
	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/middleware"
	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/router"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/data/sqlxpkg"
)

// App struct
type App struct {
	DBClient    *sqlxpkg.Client
	router      *router.Router
	Middlewares []any
	AuthModule  *auth.Module
	UserModule  *user.Module
	Validate    *validator.Validate
}

func NewApp() *App {
	a := new(App)
	a.initComponents()
	return a
}

func (a *App) initDB() {
	a.DBClient = sqlxpkg.GetInstance()
}

func (a *App) initMiddlewares() {
	authMiddleWare := middleware.NewAuthMiddleware(a.AuthModule.Service)
	a.Middlewares = append(a.Middlewares, authMiddleWare)
}

func (a *App) initModules() {
	a.UserModule = user.NewModule(a.DBClient.DB, a.Validate)
	a.AuthModule = auth.NewModule(a.UserModule.Service)
}

func (a *App) initModuleRouters() {
	m := a.Middlewares[0].(*middleware.AuthMiddleware)
	router.RegisterUserRoutes(a.router, constant.V1, a.UserModule, m)
}

// Init app
func (a *App) initComponents() {
	a.initDB()
	a.router = router.NewRouter()
	a.initModules()
	a.initMiddlewares()
	a.initModuleRouters()
}

// Run app
func (a *App) Run() {
	err := http.ListenAndServe(":8080", a.router.Mux)
	if err != nil {
		panic(err)
	}
}

// Run app
func (a *App) RunTLS() {
	err := http.ListenAndServeTLS(":443", "cert.crt", "key.key", a.router.Mux)
	if err != nil {
		panic(err)
	}
}
