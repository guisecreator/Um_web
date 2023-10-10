package middleware

import (
	"errors"
	"fmt"
	"github.com/guisecreator/um_web/graphql/model"
	"github.com/guisecreator/um_web/pkg/authpayload"
	"log"
	"net/http"
)

type HandleAuthHTTPFunc func(w http.ResponseWriter, r *http.Request) (*http.Request, error)

func AuthMiddleware() func(http.Handler) http.Handler {
	return func(nextHttp http.Handler) http.Handler {
		return http.HandlerFunc(
			func(responseWriter http.ResponseWriter, request *http.Request) {
				ctx := request.Context()

				payload := authpayload.AuthPayload{
					ResponseWriter: responseWriter,
					Request:        *request,
					AuthInfo:       model.AuthInfo{},
				}

				logRequest(responseWriter, request)

				cookie, err := request.Cookie("auth_cookie")
				if err != nil || cookie == nil {
					log.Printf("Куки 'auth_cookie' не установлены или ошибка: %v", err.Error())
					request = request.WithContext(payload.WithContext(ctx))
					nextHttp.ServeHTTP(responseWriter, request)
					return
				} else {
					log.Printf("Значение куки: %s", cookie.Value)
				}

				user := model.User{}
				user.Role = model.Roles(cookie.Value)
				if user.Role != model.RolesAdmin {
					nextHttp.ServeHTTP(responseWriter, request)
					log.Printf("User %s not have admin role", user.Email)
					return
				}

				auth := model.AuthInfo{}
				auth.Token = cookie.Value
				payload.AuthInfo = auth

				request = request.WithContext(ctx)
				nextHttp.ServeHTTP(responseWriter, request)
			},
		)
	}
}

func logRequest(w http.ResponseWriter, r *http.Request) {
	for key, value := range r.Header {
		fmt.Printf("Header: %s = %s\n", key, value)
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
