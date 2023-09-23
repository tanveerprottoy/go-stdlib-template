package action

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/data/postgres"
	"github.com/tanveerprottoy/stdlib-go-template/internal/template/module/action/entity"
)

// Module holds the components of the current module
type Module struct {
	Handler    *Handler
	Service    *Service
	Repository postgres.Repository[entity.Action]
}

// NewModule initializes a new Module
func NewModule(db *sql.DB, v *validator.Validate) *Module {
	// init order is reversed of the field decleration in the struct
	// as the dependency is served this way
	r := NewRepository(db)
	s := NewService(r)
	h := NewHandler(s, v)
	return &Module{Handler: h, Service: s, Repository: r}
}
