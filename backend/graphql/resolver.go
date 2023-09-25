package graphql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/guisecreator/um_web/pkg/authpayload"
	"github.com/guisecreator/um_web/pkg/sessions"
	"github.com/uptrace/bun"

	"github.com/graph-gophers/dataloader"
	"github.com/guisecreator/um_web/graphql/model"
)

type Resolver struct {
	Db       *bun.DB
	Sessions sessions.ManageSessions
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

func (U *UserReader) GetUser(ctx context.Context, userID string) (*model.User, error) {
	loaders := authpayload.For(ctx)
	thunk := loaders.
		UserLoader.
		Load(ctx, dataloader.StringKey(userID))
	result, err := thunk()
	if err != nil {
		return nil, err
	}

	user, ok := result.(*model.User)
	if !ok {
		return nil, fmt.Errorf("user not found %s", userID)
	}

	result = user
	return result.(*model.User), nil
}

func (r *Resolver) AdminOnlyQuery(ctx context.Context) (model.SomeType, error) {
	result, err := CheckAuthorization(ctx, nil)
	if err != nil {
		return model.SomeType{}, err
	}

	return result.(model.SomeType), err
}
