package serverLogIn



import (
	"net/http"
	"github.com/gorilla/sessions"
	"github.com/ASeegull2/fridge/SmartFridge-server_db_1/server/database"
	"github.com/ASeegull2/serverLogIn/SmartFridgeAuthorization/serverLogIn/pages"
	"fmt"
	"github.com/gorilla/securecookie"
)


var hfejke string
var hashKey = securecookie.GenerateRandomKey(64)
var blockKey = securecookie.GenerateRandomKey(32)
var store = sessions.NewFilesystemStore("hfejke",hashKey,blockKey)


func SessionSet(w http.ResponseWriter, r *http.Request)  {
	session, err := store.Get(r, "session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["name"] = r.FormValue("name")
	session.Values["authenticated"] = true
	session.Save(r, w)

}
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	pass := r.FormValue("password")
	redirectTarget := "/"

	err:= database.ClientLogin(name , pass)

	if (err == nil) || (pass != "") {
		SessionSet (w , r)
		redirectTarget = "/internal"

	}
	http.Redirect(w, r, redirectTarget, http.StatusFound)
}
func SignUpHandler(w http.ResponseWriter, r *http.Request) {

	Name := r.FormValue("name")
	pass := r.FormValue("password")
	redirectTarget := "/signup"

	err:= database.RegisterNewUser(Name , pass )

	if (err == nil) || (pass != "")  {

		SessionSet (w , r)
		redirectTarget = "/internal"
	}
	http.Redirect(w, r, redirectTarget, http.StatusFound)
}


func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["name"] = ""
	session.Values["authenticated"] = false
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusFound)
}

func IndexPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, pages.IndexPage)
}
func IndexPageTwoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, pages.IndexTwoPage)
}

func InternalPageHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if  session.Values["authenticated"] == true{
		fmt.Fprintf(w, pages.InternalPage,	session.Values["name"])
	} else {
		http.Redirect(w, r, "/", http.StatusFound )
	}
}
