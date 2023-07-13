package fileupload

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Module struct {
	Handler    *Handler
	Service    *Service
}

func NewModule(s3Client *s3.Client) *Module {
	m := new(Module)
	// init order is reversed of the field decleration
	// as the dependency is served this way
	m.Service = NewService(s3Client)
	m.Handler = NewHandler(m.Service)
	return m
}
