package graphql

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"github.com/guisecreator/um_web/graphql/model"
	"github.com/guisecreator/um_web/pkg/authpayload"
	"github.com/guisecreator/um_web/pkg/token"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"log"
	"time"
)

func CheckAuthorization(ctx context.Context,
	next graphql.Resolver) (interface{}, error) {
	notAuth := errors.New("not authenticated")

	user, err := getUserFromContext(ctx)
	if err != nil {
		return nil, notAuth
	}

	insufficientAccess := errors.New(
		"insufficient access rights")
	if !IsValid(model.Roles(user.Role)) {
		return nil, insufficientAccess
	}

	return next(ctx)
}

func getUserFromContext(ctx context.Context) (
	*model.User, error) {
	user, ok := ctx.Value("user").(*model.User)
	if !ok {
		return nil, errors.New("failed to get user from context")
	}

	return user, nil
}

func IsValid(roles model.Roles) bool {
	switch roles {
	case model.RolesAdmin, model.RolesUser:
		return true
	}
	return false
}

func (r *Resolver) HasRole(
	roles []model.Roles,
	roleToCheck model.Roles) bool {
	for _, r := range roles {
		if model.Roles(r) == roleToCheck {
			return true
		}
	}
	return false
}

func (r *Resolver) Authentication(
	ctx context.Context, _ interface{},
	next graphql.Resolver) (interface{}, error) {
	authenticationError := errors.New("access is denied")

	auth := authpayload.ForContext(ctx)
	sessionKey := auth.AuthInfo.Token
	if sessionKey == "" {
		log.Println("No token")
		return nil, authenticationError
	}

	session, found := r.Sessions.GetSession(sessionKey)
	if !found {
		log.Println("No session for key -", found)
		return nil, authenticationError
	}

	userRoles := []model.Roles{
		session.Login.Role,
	}

	onlyModerator := r.HasRole(userRoles, model.RolesModerator)
	onlyAdmin := r.HasRole(userRoles, model.RolesAdmin)
	if !onlyAdmin && !onlyModerator {
		log.Printf("User %s not have admin role or moderator role", session.Login.Email)
		return nil, authenticationError
	}

	errPrivateKey := session.PrivateKey.Validate()
	if errPrivateKey != nil {
		log.Println("Private key error -", errPrivateKey)
		return nil, authenticationError
	}

	tokenJSON, err := token.PullTokenFromJSON(auth.Token)
	if err != nil {
		log.Println("Token error from JSON -", err)
		return nil, authenticationError
	}

	claims, err := token.
		DecryptBytes(tokenJSON, session.PrivateKey)
	if err != nil {
		log.Println("Decrypt error -", err)
		return nil, authenticationError
	}

	if claims.Expired < time.Now().Unix() {
		log.Printf("Token expired for %v [%v]\n",
			claims.Id, time.Unix(claims.Expired, 0))
		return nil, &gqlerror.Error{
			Message: "Your session has expired. Please re-login",
			Extensions: map[string]interface{}{
				"code": "expired",
			},
		}
	}

	return next(ctx)
}
