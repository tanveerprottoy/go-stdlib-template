package user

import (
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/tanveerprottoy/stdlib-go-template/internal/app/template/module/user/entity"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/data/sqlxpkg"
)

type Module struct {
	Handler    *Handler
	Service    *Service
	Repository sqlxpkg.Repository[entity.User]
}

func NewModule(db *sqlx.DB, validate *validator.Validate) *Module {
	m := new(Module)
	// init order is reversed of the field decleration
	// as the dependency is served this way
	m.Repository = NewRepository(db)
	m.Service = NewService(m.Repository)
	m.Handler = NewHandler(m.Service, validate)
	return m
}
