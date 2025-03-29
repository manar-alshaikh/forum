package functions

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

const (
	sessionCookieName = "session_token"
	sessionDuration   = 24 * time.Hour // Session lasts for 24 hours
)

var(
	sessions      = map[string]Session{}
)

type Session struct {
	Username string
	Expiry   time.Time
	ID       string
}

func (s Session) IsExpired() bool {
	return s.Expiry.Before(time.Now())
}

func CreateSession(username string, w http.ResponseWriter) {
	sessionToken := uuid.New().String()
	expiresAt := time.Now().Add(sessionDuration)

	sessions[sessionToken] = Session{
		Username: username,
		Expiry:   expiresAt,
	}

	log.Println("Session created for:", username, "Token:", sessionToken)

	http.SetCookie(w, &http.Cookie{
		Name:     sessionCookieName,
		Value:    sessionToken,
		Expires:  expiresAt,
		HttpOnly: true,
		Path:     "/",
	})
}

func GetSession(r *http.Request) (*Session, error) {
	c, err := r.Cookie(sessionCookieName)
	if err != nil {
		return nil, err
	}

	session, ok := sessions[c.Value]
	if !ok {
		return nil, fmt.Errorf("session not found")
	}

	if session.IsExpired() {
		delete(sessions, c.Value)
		return nil, fmt.Errorf("session expired")
	}

	return &session, nil
}

func DeleteSession(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie(sessionCookieName)
	if err != nil {
		return
	}

	delete(sessions, c.Value)

	// Set cookie to expire
	http.SetCookie(w, &http.Cookie{
		Name:     sessionCookieName,
		Value:    "",
		Expires:  time.Now(),
		HttpOnly: true,
		Path:     "/",
	})
}