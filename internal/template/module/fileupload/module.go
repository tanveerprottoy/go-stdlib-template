package fileupload

import (
	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/s3pkg"
)

type Module struct {
	Handler *Handler
	Service *Service
}

func NewModule(clientsS3 *s3pkg.Clients) *Module {
	m := new(Module)
	// init order is reversed of the field decleration
	// as the dependency is served this way
	m.Service = NewService(clientsS3)
	m.Handler = NewHandler(m.Service)
	return m
}
