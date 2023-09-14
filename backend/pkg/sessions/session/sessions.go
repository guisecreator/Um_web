package session

import (
	"crypto/rsa"
	"github.com/guisecreator/um_backend/graphql/model"
	"github.com/guisecreator/um_backend/pkg/sessions"
)

type Session struct {
	model.Roles
	model.User
	Id         int    //ID of the session
	Name       string // Name of the session
	PrivateKey rsa.PrivateKey
}

type typeSessions map[string]*Session

func NewObjectOfSessions() sessions.ManageSessions {
	return &typeSessions{}
}

func (s *typeSessions) AddSession(
	tokenStr string, session *Session) {
	(*s)[tokenStr] = session

	sessionName := session.Name
	if sessionName == "" {
		sessionName = tokenStr
	}

	session.Id = len(*s)
}

func (s *typeSessions) GetSession(
	tokenStr string) (*Session, bool) {

	value, ok := (*s)[tokenStr]
	if !ok {
		return nil, false
	}
	return value, true
}

func (s *typeSessions) RemoveSession(tokenStr string) {
	delete(*s, tokenStr)
}
