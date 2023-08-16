package s3pkg

import (
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	instance *Clients
	once     sync.Once
)

// S3Client struct
type Clients struct {
	S3Client      *s3.Client
	PresignClient *s3.PresignClient
}

// GetInstance returns a singleton of Client
func GetInstance() *Clients {
	once.Do(func() {
		instance = new(Clients)
	})
	return instance
}

// Init initializes the client with options
//
//	ex: options := s3.Options{
//	   Region:      "us-west-2",
//	   Credentials: aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
//	}
//
// ex: override
//
//	func(o *s3.Options) {
//		o.Region = "us-east-1"
//		o.UseAccelerate = true
//	})
func (c *Clients) Init(options s3.Options, optFn func(*s3.Options)) {
	if optFn != nil {
		c.S3Client = s3.New(options, optFn)
	} else {
		c.S3Client = s3.New(options)
	}
	// init presignClient
	c.PresignClient = s3.NewPresignClient(c.S3Client)
}

// InitWithConfig initializes the client with the
// passed config and if override needed pass the optFn
// ex: cfg, err := config.LoadDefaultConfig(context.TODO())
//
//	optFun = func(o *s3.Options) {
//		o.Region = "us-west-2"
//		o.UseAccelerate = true
//	})
func (c *Clients) InitWithConfig(cfg aws.Config, optFn func(*s3.Options)) {
	if optFn != nil {
		c.S3Client = s3.NewFromConfig(cfg, optFn)
	} else {
		c.S3Client = s3.NewFromConfig(cfg)
	}
}
