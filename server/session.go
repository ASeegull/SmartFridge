package server

import (
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

var hashKey = securecookie.GenerateRandomKey(64)
var blockKey = securecookie.GenerateRandomKey(32)

var store = sessions.NewCookieStore(hashKey, blockKey)

const (
	sessionName = "sessionName"
	ID          = "userID"
)

func sessionSet(w http.ResponseWriter, r *http.Request, userID string) error {
	session, err := store.Get(r, sessionName)
	if err != nil {
		session.Options = &sessions.Options{
			MaxAge:   3600,
			Path:     "/",
			HttpOnly: true,
		}
		session.Values[ID] = userID
		return session.Save(r, w)
	}

	session.Options = &sessions.Options{
		MaxAge:   3600,
		Path:     "/",
		HttpOnly: true,
	}
	session.Values[ID] = userID
	return session.Save(r, w)
}

func closeSession(w http.ResponseWriter, r *http.Request) error {

	session, err := store.Get(r, sessionName)
	if err != nil {
		return err
	}

	session.Options = &sessions.Options{
		MaxAge: -1,
	}

	return session.Save(r, w)
}

func isNewSession(w http.ResponseWriter, r *http.Request) (bool, error) {
	session, err := store.Get(r, sessionName)
	if err != nil {
		return false, err
	}

	if session.IsNew {
		session.Options = &sessions.Options{
			MaxAge: -1,
		}
		return true, session.Save(r, w)
	}

	return false, nil
}

func getUserID(r *http.Request) (string, error) {
	session, err := store.Get(r, sessionName)
	if err != nil {
		return "", err
	}
	return session.Values[ID].(string), err
}