package entity

type Role struct {
	ID        string `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	Key       string `db:"key" json:"key"`
	IsDeleted bool   `db:"is_deleted" json:"isDeleted"`
	CreatedAt int64  `db:"created_at" json:"createdAt"`
	UpdatedAt int64  `db:"updated_at" json:"updatedAt"`
}
