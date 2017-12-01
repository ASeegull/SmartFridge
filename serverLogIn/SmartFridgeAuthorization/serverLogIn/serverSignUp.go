package serverLogIn

/*

import (
	"fmt"

	"net/http"
	"github.com/ASeegull2/fridge/SmartFridge-server_db_1/server/database"
	"github.com/ASeegull2/serverLogIn/SmartFridgeAuthorization/serverLogIn/pages"
)


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
	redirectTarget := "/signup"

	err:= database.RegisterNewUser(name , pass )

	if (err == nil) || (pass != "")  {

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
*/
