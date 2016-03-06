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

// Starts the session manager.
func StartSessionManager() {
	sm, err := session.NewManager(cookie_name, max_lifetime, session.NewMemProvider())

	if err != nil {
		panic(err)
	}

	sm.GC()
	sessionManager = sm
}

// Returns the session manager.
func GetSessionManager() *session.Manager {
	return sessionManager
}

// Returns true if user is logged in for request, else false.
func IsLoggedIn(r *http.Request) bool {
	if r == nil {
		return false
	}

	session, _ := sessionManager.GetCurrentSession(r)
	return session.IsValid()
}
