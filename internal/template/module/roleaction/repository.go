package roleaction

import (
	"database/sql"
)

const tableName = "roles_actions"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(q string, args ...any) (string, error) {
	var id string
	err := r.db.QueryRow(q, args...).Scan(&id)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (r *Repository) ReadMany(limit, offset int) (*sql.Rows, error) {
	rows, err := r.db.Query("SELECT * FROM "+tableName+" LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *Repository) ReadManyForRole(roleID string, limit, offset int) (*sql.Rows, error) {
	rows, err := r.db.Query("SELECT * FROM "+tableName+" WHERE role_id = $1 LIMIT $2 OFFSET $3", roleID, limit, offset)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *Repository) ReadManyActionsForRole(roleID string) (*sql.Rows, error) {
	rows, err := r.db.Query("SELECT JSON_AGG(action_ids) as action_ids FROM "+tableName+" WHERE role_id = $1", roleID)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *Repository) ReadManyActionsForRole1(roleID string) (*sql.Rows, error) {
	rows, err := r.db.Query("SELECT json_array_elements(JSON_AGG(action_ids)) as action_ids FROM "+tableName+" WHERE role_id = $1", roleID)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
