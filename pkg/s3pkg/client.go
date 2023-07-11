package s3pkg

import (
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/config"
)

var (
	instance *Client
	once     sync.Once
)

type Client struct {
	// Uploader declaration
	Uploader *s3manager.Uploader

	// Session declaration
	Session *session.Session

	// Session declaration
	S3Svc *s3.S3
}

func GetInstance() *Client {
	once.Do(func() {
		instance = new(Client)
		instance.init()
	})
	return instance
}

func (d *Client) init() {
	r := config.GetEnvValue("AWS_REGION")
	d.Session = session.Must(
		session.NewSession(
			&aws.Config{
				Region: aws.String(r),
				Credentials: credentials.NewStaticCredentials(
					AccessKeyID,
					SecretAccessKey,
					"", // a token will be created when the session it's used.
				),
			},
		),
	)
	// Create an uploader with the session and default options
	d.Uploader = s3manager.NewUploader(d.Session)
	d.S3Svc = s3.New(d.Session)
}
