package dbmodels

import (
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel
	ID         string  `bun:"id,pk,autoincrement"`
	Email      string  `bun:"email,notnull"`
	Password   string  `bun:"password,notnull"`
	Roles      string  `bun:"role,notnull"`
	create_at  string  `bun:"createAt"`
	update_at  string  `bun:"updateAt"`
	deleted_at *string `bun:"deletedAt"`
}
