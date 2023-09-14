package sessions

import "github.com/guisecreator/um_backend/pkg/sessions/session"

type ManageSessions interface {
	AddSession(tokenStr string, session *session.Session)
	GetSession(tokenStr string) (*session.Session, bool)
	RemoveSession(tokenStr string)
}
