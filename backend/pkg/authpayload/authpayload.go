package authpayload

import (
	"context"
	"github.com/graph-gophers/dataloader"
	"github.com/guisecreator/um_backend/graphql/model"
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
	loadersKey               = ctxKey("dataloaders")
)

const CookieName string = "_auth"

func ForContext(ctx context.Context) *AuthPayload {
	raw := ctx.Value(ctxAuthPayloadKey).(*AuthPayload)
	return raw
}

// For returns the dataloader for a given context
func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}

func (auth *AuthPayload) WithContext(
	ctx context.Context) context.Context {
	return context.WithValue(ctx, ctxAuthPayloadKey, auth)
}
