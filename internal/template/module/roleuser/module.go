package comprojroleuser

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
)

// Module holds the components of the current module
type Module struct {
	Service    *Service
	Repository *Repository
}

// NewModule initializes a new Module
func NewModule(db *sql.DB, v *validator.Validate) *Module {
	// init order is reversed of the field decleration in the struct
	// as the dependency is served this way
	r := NewRepository(db)
	s := NewService(r)
	return &Module{Service: s, Repository: r}
}
