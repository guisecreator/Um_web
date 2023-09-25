package middleware

import (
	"errors"
	"github.com/guisecreator/um_web/graphql/model"
	"github.com/guisecreator/um_web/pkg/authpayload"
	"log"
	"net/http"
)

// AuthMiddleware is a middleware that checks and extracts user ID from cookie
func AuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				ctx := r.Context()

				payload := authpayload.AuthPayload{
					ResponseWriter: w,
					Request:        *r,
				}

				cookie, err := r.Cookie(authpayload.CookieName)
				if err != nil {
					log.Printf("Cookie retrieval error: %v ",
						err.Error())
					r = r.WithContext(payload.
						WithContext(ctx))
					next.ServeHTTP(w, r)
					return
				}

				if err != nil || cookie == nil {
					next.ServeHTTP(w, r)
					return
				}

				auth := model.AuthInfo{}
				auth.Token = cookie.Value
				payload.AuthInfo = auth

				r = r.WithContext(ctx)
				next.ServeHTTP(w, r)
			},
		)
	}
}

// function to check and extract user ID from cookie
func validateAndGetUserID(c *http.Cookie) (string, error) {
	userId := c.Value
	if userId == "" {
		return "0", errors.New("invalid cookie")
	}
	return userId, nil
}
