package entity

type User struct {
	ID   string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Role string `db:"role" json:"role"`
	/* Email         string            `db:"email" json:"email"`
	Age           uint8             `db:"age" json:"age"`
	Phone         string            `db:"phone" json:"phone"`
	FavoriteColor string            `db:"favorite_color" json:"favoriteColor"`
	Addresses     []*UserAddressDTO `db:"addresses" json:"addresses"` */
	CreatedAt int64 `db:"created_at" json:"createdAt"`
	UpdatedAt int64 `db:"updated_at" json:"updatedAt"`
}

// Address houses a users address information
type UserAddressDTO struct {
	Street string `db:"street" json:"street"`
	City   string `db:"city" json:"city"`
}
