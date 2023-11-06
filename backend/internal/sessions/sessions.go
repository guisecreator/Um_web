package sessions

import (
	"crypto/rsa"
	"github.com/guisecreator/um_web/graphql/model"
)

type SessionManager interface {
	AddSession(token string, session *Session)
	GetSession(token string) (Session, bool)
	RemoveSession(tokenStr string)
}

type Session struct {
	Id         string
	Name       string
	Roles      model.Roles
	Login      model.User
	PrivateKey rsa.PrivateKey
}

type typeSessions map[string]Session

func NewObjectOfSessions() SessionManager {
	return &typeSessions{}
}

func (s typeSessions) AddSession(token string, session *Session) {
	s[token] = *session
}

func (s typeSessions) GetSession(token string) (Session, bool) {
	value, ok := s[token]
	return value, ok
}

func (s *typeSessions) RemoveSession(tokenStr string) {
	delete(*s, tokenStr)
}
