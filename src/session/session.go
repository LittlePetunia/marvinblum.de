// Package session provides session protection for ONE user.
// This will be me only. If I login again, the session will be reset.
// On logout the session token will be set to an empty string.
package session

import (
	"crypto/sha256"
	"encoding/base64"
	"strconv"
	"sync"
	"time"
)

type Session struct {
	m     sync.Mutex
	token string
}

// Creates a new session token.
func (s *Session) Login(login, password string) {
	s.m.Lock()
	s.token = generateSessionToken(login, password)
	s.m.Unlock()
}

// Resets the session token.
func (s *Session) Logout() {
	s.m.Lock()
	s.token = ""
	s.m.Unlock()
}

// Checks if tokens match.
func (s *Session) LoggedIn(token string) bool {
	return token == s.token
}

// Returns the session token.
func (s *Session) GetToken() string {
	return s.token
}

// Generates session token.
func generateSessionToken(login, password string) string {
	now := strconv.Itoa(time.Now().Nanosecond())
	hash := sha256.New()
	hash.Write([]byte(login + password + now))
	tokenString := base64.URLEncoding.EncodeToString(hash.Sum(nil))

	return tokenString
}
