package authpayload

import (
	"context"
	"errors"
	"github.com/graph-gophers/dataloader"
	"github.com/guisecreator/um_web/graphql/model"
	"net/http"
)

type ctxKey string

type AuthPayload struct {
	http.ResponseWriter
	http.Request
	model.AuthInfo
}

type Loaders struct {
	UserLoader *dataloader.Loader
}

const (
	ctxAuthPayloadKey ctxKey = "ctxAuthPayloadKey"
	ctxAuthInfoKey    ctxKey = "ctxAuthInfoKey"
)

func ForContext(ctx context.Context) *AuthPayload {
	raw, _ := ctx.Value(ctxAuthPayloadKey).(*AuthPayload)
	return raw
}

func GetSessionTokenFromContext(ctx context.Context) (string, error) {
	auth := ForContext(ctx)

	if auth.AuthInfo.Token == "" {
		return "", errors.New("no token")
	}

	return auth.AuthInfo.Token, nil
}

func (auth *AuthPayload) WithContext(
	ctx context.Context) context.Context {
	return context.WithValue(ctx, ctxAuthPayloadKey, auth)
}
