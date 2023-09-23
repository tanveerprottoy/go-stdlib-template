package entity

type Action struct {
	ID        string `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	Key       string `db:"keys" json:"keys"`
	IsDeleted bool   `db:"is_deleted" json:"isDeleted"`
	CreatedAt int64  `db:"created_at" json:"createdAt"`
	UpdatedAt int64  `db:"updated_at" json:"updatedAt"`
}
