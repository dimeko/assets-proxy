package api

import (
	"net/http"

	"github.com/gorilla/sessions"
)

/* Authentication related controllers */
const (
	sessionId  = "sessionId"
	sessionKey = "secret_key"
)

var store = sessions.NewCookieStore([]byte(sessionKey))

func Session(r *http.Request) (*sessions.Session, error) {
	return store.Get(r, sessionId)
}

func SessionUser(r *http.Request) string {
	session, _ := Session(r)
	username := session.Values["username"].(string)

	return username
}
