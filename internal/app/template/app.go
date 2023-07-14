package template

import (
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/go-playground/validator/v10"
	"github.com/tanveerprottoy/stdlib-go-template/internal/app/template/module/auth"
	"github.com/tanveerprottoy/stdlib-go-template/internal/app/template/module/fileupload"
	"github.com/tanveerprottoy/stdlib-go-template/internal/app/template/module/user"
	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/constant"
	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/middleware"
	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/router"
	configpkg "github.com/tanveerprottoy/stdlib-go-template/pkg/config"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/data/sqlxpkg"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/file"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/s3pkg"
)

// App struct
type App struct {
	DBClient         *sqlxpkg.Client
	router           *router.Router
	Middlewares      []any
	AuthModule       *auth.Module
	UserModule       *user.Module
	FileUploadModule *fileupload.Module
	Validate         *validator.Validate
	ClientS3         *s3pkg.Client
}

func NewApp() *App {
	a := new(App)
	a.initComponents()
	return a
}

func (a *App) initDB() {
	a.DBClient = sqlxpkg.GetInstance()
}

func (a *App) initDir() {
	file.CreateDirIfNotExists("./uploads")
}

func (a *App) initS3() {
	s3Region := configpkg.GetEnvValue("S3_REGION")
	s3Endpoint := configpkg.GetEnvValue("S3_ENDPOINT")
	endpointResolverFunc := s3.EndpointResolverFunc(func(region string, options s3.EndpointResolverOptions) (aws.Endpoint, error) {
		if s3Endpoint != "" {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           s3Endpoint,
				SigningRegion: s3Region,
			}, nil
		}
		// returning EndpointNotFoundError will allow the service to fallback to it's default resolution
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})
	a.ClientS3 = s3pkg.GetInstance()
	a.ClientS3.Init(s3.Options{
		Region: configpkg.GetEnvValue("S3_REGION"),
		// Credentials: aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(config.GetEnvValue("S3_ACCESS_KEY"), config.GetEnvValue("S3_SECRET_KEY"), "")),
		// EndpointResolver: endpointResolverFunc,
		// UsePathStyle: true,
	}, func(o *s3.Options) {
		o.EndpointResolver = endpointResolverFunc
		o.UsePathStyle = true
	})
}

func (a *App) initMiddlewares() {
	authMiddleWare := middleware.NewAuthMiddleware(a.AuthModule.Service)
	a.Middlewares = append(a.Middlewares, authMiddleWare)
}

func (a *App) initModules() {
	a.UserModule = user.NewModule(a.DBClient.DB, a.Validate)
	a.AuthModule = auth.NewModule(a.UserModule.Service)
	a.FileUploadModule = fileupload.NewModule(a.ClientS3.S3Client)
}

func (a *App) initModuleRouters() {
	m := a.Middlewares[0].(*middleware.AuthMiddleware)
	router.RegisterUserRoutes(a.router, constant.V1, a.UserModule, m)
	router.RegisterFileUploadRoutes(a.router, constant.V1, a.FileUploadModule)
}

// Init app
func (a *App) initComponents() {
	a.initDB()
	a.initDir()
	a.router = router.NewRouter()
	a.initS3()
	a.initModules()
	a.initMiddlewares()
	a.initModuleRouters()
}

// Run app
func (a *App) Run() {
	err := http.ListenAndServe(":"+configpkg.GetEnvValue("APP_PORT"), a.router.Mux)
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
