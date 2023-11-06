package graphql

import (
	"context"
	"database/sql"
	"github.com/guisecreator/um_web/internal/sessions"
	"github.com/uptrace/bun"

	"github.com/guisecreator/um_web/graphql/model"
)

type Resolver struct {
	Db       *bun.DB
	Sessions sessions.SessionManager
}

type UserReader struct {
	conn *sql.DB
}

type authenticatedDirective struct{}

func NewUserReader(db *sql.DB) *UserReader {
	return &UserReader{
		conn: db,
	}
}

func (r *Resolver) AdminOnlyQuery(ctx context.Context) (model.SomeType, error) {
	result, err := CheckAuthorization(ctx, nil)
	if err != nil {
		return model.SomeType{}, err
	}

	return result.(model.SomeType), err
}
