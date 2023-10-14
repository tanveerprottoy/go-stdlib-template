package content

import (
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/data/sqlxext"
	"github.com/tanveerprottoy/stdlib-go-template/internal/template/module/content/entity"
)

type Module struct {
	Handler    *Handler
	Service    *Service
	Repository sqlxext.Repository[entity.Content]
}

func NewModule(db *sqlx.DB, v *validator.Validate) *Module {
	// init order is reversed of the field decleration
	// as the dependency is served this way
	r := NewRepository(db)
	s := NewService(r)
	h := NewHandler(s, v)
	return &Module{Handler: h, Service: s, Repository: r}
}
