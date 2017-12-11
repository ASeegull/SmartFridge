package server

import (
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

var hashKey = securecookie.GenerateRandomKey(64)
var blockKey = securecookie.GenerateRandomKey(32)

//var path = "DbPath"
//var store = sessions.NewFilesystemStore(path,hashKey,blockKey)
var store = sessions.NewCookieStore(hashKey, blockKey)

const sessionName = "sessionName"

func sessionSet(w http.ResponseWriter, r *http.Request, userID string) error {
	session, err := store.Get(r, sessionName)
	if err != nil {
		return err
	}

	session.Options = &sessions.Options{
		MaxAge:   3600,
		Path:     "/",
		HttpOnly: true,
	}
	session.Values["userID"] = userID
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

func checkOutUser(w http.ResponseWriter, r *http.Request) (bool, error) {
	session, err := store.Get(r, sessionName)
	if err != nil {
		return false, err
	}
	if session.IsNew == true {
		session.Options = &sessions.Options{
			MaxAge: -1,
		}
		return false, session.Save(r, w)
	}
	return true, nil
}
