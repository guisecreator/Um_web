package sessions

import (
	"crypto/rsa"
	"github.com/guisecreator/um_web/graphql/model"
	"strconv"
)

type ManageSessions interface {
	AddSession(tokenStr string, session Session)
	GetSession(tokenStr string) (Session, bool)
	RemoveSession(tokenStr string)
}

type Session struct {
	Roles      model.Roles
	Login      model.User
	Id         string //ID of the session
	Name       string // Name of the session
	PrivateKey rsa.PrivateKey
}

type typeSessions map[string]*Session

func NewObjectOfSessions() ManageSessions {
	return &typeSessions{}
}

func (s *typeSessions) AddSession(
	tokenStr string, session Session) {
	(*s)[tokenStr] = &session

	sessionName := session.Name
	if sessionName == "" {
		sessionName = tokenStr
	}

	session.Id = strconv.Itoa(len(*s))
}

func (s *typeSessions) GetSession(
	tokenStr string) (Session, bool) {

	value, ok := (*s)[tokenStr]
	if !ok {
		return Session{}, false
	}
	return *value, true
}

func (s *typeSessions) RemoveSession(tokenStr string) {
	delete(*s, tokenStr)
}
