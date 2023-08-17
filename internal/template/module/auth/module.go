package auth

import "github.com/tanveerprottoy/stdlib-go-template/internal/template/module/user"

type Module struct {
	Service *Service
}

func NewModule(s *user.Service) *Module {
	m := new(Module)
	m.Service = NewService(s)
	return m
}
