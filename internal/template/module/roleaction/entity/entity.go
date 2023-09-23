package entity

import "github.com/tanveerprottoy/stdlib-go-template/internal/pkg/data/postgres"

type RoleAction struct {
	RoleID    string                   `db:"role_id" json:"roleId"`
	ActionIDs postgres.JsonStringArray `db:"action_ids" json:"actionIds"`
}
