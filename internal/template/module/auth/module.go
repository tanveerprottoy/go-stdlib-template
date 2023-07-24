package auth

import "github.com/tanveerprottoy/stdlib-go-template/internal/template/module/user"

type Module struct {
	Service *Service
}

func NewModule(userService *user.Service) *Module {
	m := new(Module)
	m.Service = NewService(userService)
	return m
}
