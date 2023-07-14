package template

import (
	"context"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
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
	bucketName := configpkg.GetEnvValue("BUCKET_NAME")
	customResolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
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
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(s3Region),
		config.WithEndpointResolver(customResolver),
	)
	if err != nil {
		log.Fatalf("Cannot load the AWS configs: %s", err)
	}
	a.ClientS3 = s3pkg.GetInstance()
	/* a.ClientS3.Init(s3.Options{
		Region: config.GetEnvValue("S3_REGION"),
		// Credentials: aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(config.GetEnvValue("S3_ACCESS_KEY"), config.GetEnvValue("S3_SECRET_KEY"), "")),
	}, func(o *s3.Options) {
		o.EndpointOptions = s3.EndpointResolver{
			return s3.ResolveEndpoint()
			Endpoint: config.GetEnvValue("S3_ENDPOINT"),
		}
	}) */
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
