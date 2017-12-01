package serverLogIn



import (
	"fmt"

	"github.com/gorilla/securecookie"
	"net/http"
	"github.com/SmartFridge/server/database"
	"github.com/SmartFridge/serverLogIn/pages"
)



var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func getUserName(r *http.Request) (userName string) {
	if cookie, err := r.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}
	return userName
}

func setSession(userName string, w http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
	}
}

func clearSession(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}



func LoginHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	pass := r.FormValue("password")
	redirectTarget := "/"

	err:= database.ClientLogin(name , pass)

	if (err == nil) || (pass != "") {

		setSession(name, w)
		redirectTarget = "/internal"
	}
	http.Redirect(w, r, redirectTarget, 302)
}


func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	pass := r.FormValue("password")
	passconfirm := r.FormValue("passwordconfirm")
	redirectTarget := "/register"

	err:= database.RegisterNewClient(name , pass , passconfirm)

	if (err == nil) || (pass == passconfirm) || (pass != "")  {

		setSession(name, w)
		redirectTarget = "/internal"
	}
	http.Redirect(w, r, redirectTarget, 302)
	}


func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	clearSession(w)
	http.Redirect(w, r, "/", 302)
}

func IndexPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, pages.IndexPage)
}
func IndexPageTwoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, pages.IndexTwoPage)
}


func InternalPageHandler(w http.ResponseWriter, r *http.Request) {
	userName := getUserName(r)
	if userName != "" {
		fmt.Fprintf(w, pages.InternalPage, userName)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}
// use jwt (JSON web token)
