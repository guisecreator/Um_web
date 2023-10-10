package authpayload

import (
	"context"
	"github.com/graph-gophers/dataloader"
	"github.com/guisecreator/um_web/graphql/model"
	"log"
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

func ForContext(ctx context.Context) *AuthPayload {
	raw := ctx.Value(ctxAuthPayloadKey).(*AuthPayload)
	if raw == nil {
		log.Println("raw is nil")
	}
	return raw
}

// For returns the dataloader for a given context
func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}

func GetSessionTokenFromContext(ctx context.Context) (string, error) {
	raw := ctx.Value(ctxAuthPayloadKey).(*AuthPayload)
	return raw.AuthInfo.Token, nil
}

func (auth *AuthPayload) WithContext(
	ctx context.Context) context.Context {
	return context.WithValue(ctx, ctxAuthPayloadKey, auth)
}
