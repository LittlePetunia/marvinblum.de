package util

import (
	"github.com/DeKugelschieber/go-session"
	"net/http"
)

const (
	cookie_name  = "mb_session"
	max_lifetime = 3600 * 24
)

var (
	sessionManager *session.Manager
)

func StartSessionManager() {
	sm, err := session.NewManager(cookie_name, max_lifetime, session.NewMemProvider())

	if err != nil {
		panic(err)
	}

	sm.GC()
	sessionManager = sm
}

func GetSessionManager() *session.Manager {
	return sessionManager
}

func IsLoggedIn(r *http.Request) bool {
	session, _ := sessionManager.GetCurrentSession(r)
	return session.IsValid()
}
