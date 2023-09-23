package comprojroleuser

import "database/sql"

const tableName = "companies_projects_roles_users"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) create(q string, args ...any) (string, error) {
	var id string
	err := r.db.QueryRow(q, args...).Scan(&id)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (r *Repository) readManyForUser(userID string) (*sql.Rows, error) {
	rows, err := r.db.Query("SELECT * FROM "+tableName+" WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *Repository) readOneRoleForCompanyProjectUser(userID, companyID, projectID string) *sql.Row {
	q := "SELECT * FROM " + tableName + " WHERE user_id = $1"
	if companyID == "" && projectID == "" {
		// build query for nil companyID and projectID
		q += " AND company_id IS NULL AND project_id IS NULL LIMIT 1"
		return r.db.QueryRow(q, userID)
	}
	if projectID == "" {
		// build query for nil projectID
		q += " AND company_id = $2 AND project_id IS NULL LIMIT 1"
		return r.db.QueryRow(q, userID, companyID)
	}
	// build query for non-nil companyID and projectID
	q += " AND company_id = $2 AND project_id = $3 LIMIT 1"
	return r.db.QueryRow(q, userID, companyID, projectID)
}
