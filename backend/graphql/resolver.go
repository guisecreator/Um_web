package graphql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/guisecreator/um_backend/pkg/authpayload"
	"github.com/guisecreator/um_backend/pkg/sessions"
	"github.com/uptrace/bun"

	"github.com/99designs/gqlgen/graphql"
	"github.com/graph-gophers/dataloader"
	"github.com/guisecreator/um_backend/graphql/model"
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

//func NewLoaders(conn *sql.DB) *Resolver {
//	// define the data loader
//	userReader := &UserReader{conn: conn}
//	loaders := &Loaders{
//		UserLoader: dataloader.NewBatchedLoader(userReader.GetUsers),
//	}
//	return loaders
//}

func (d *authenticatedDirective) Validate(schema graphql.ExecutableSchema) error {
	// Проверка валидности директивы
	return nil
}

func (d *authenticatedDirective) MutateArgument(rc *graphql.ResolverContext) error {
	// Логика мутации аргумента
	return nil
}

func (r *Resolver) AdminOnlyQuery(ctx context.Context) (model.SomeType, error) {
	result, err := CheckAuthorization(ctx, nil)
	if err != nil {
		return model.SomeType{}, err
	}

	return result.(model.SomeType), err
}

//func (r *queryResolver) AdminOnlyQuery(ctx context.Context) (*model.SomeType, error) {
//	panic(fmt.Errorf("not implemented: AdminOnlyQuery - adminOnlyQuery"))
//}
