package dbmodels

import (
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel
	ID        int     `bun:"id,pk,autoincrement"`
	Name      string  `bun:"name,notnull"`
	Email     string  `bun:"email,notnull"`
	Password  string  `bun:"password,notnull"`
	Roles     string  `bun:"role,notnull"`
	CreateAt  string  `bun:"createAt"`
	UpdateAt  string  `bun:"updateAt"`
	DeletedAt *string `bun:"deletedAt"`
}
