package staticServer

import (
	"net/http"

	"github.com/ASeegull/SmartFridge/staticServer/handlers"
	config "github.com/ASeegull/SmartFridge/staticServer/staticServerConfig"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/http2"
)

var filePath, prefix = config.GetStaticPath()
var server = config.GetServerAddr()

// newRouter сreates and returns gorilla router
func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.Path("/").Handler(http.FileServer(http.Dir(filePath)))
	r.PathPrefix(prefix).Handler(http.StripPrefix(prefix, http.FileServer(http.Dir(filePath))))
	// r.Path("/login").HandlerFunc(ServeLoginPage).Methods("GET")
	// r.Path("/login/send").HandlerFunc(LoginHandler).Methods("POST")
	// r.Path("/signup").HandlerFunc(ServeSignUpPage).Methods("GET")
	// r.Path("/signup/send").HandlerFunc(SignUpHandler).Methods("GET")
	r.Path("/content").HandlerFunc(handlers.Fetch("/client/fridgeContent")).Methods("GET")
	r.Path("/recipes").HandlerFunc(handlers.Fetch("/client/allRecipes")).Methods("GET")
	return r
}

// Run starts server on defined by yaml.config port
func Run() error {
	r := newRouter()
	addr := config.GetAddr()
	server := new(http.Server)
	http2config := new(http2.Server)
	http2.ConfigureServer(server, http2config)
	log.WithField("host:port", addr).Info("Client server is listening on")
	return http.ListenAndServeTLS(addr, "../../staticServer/staticServerConfig/cert.pem", "../../staticServer/staticServerConfig/key.pem", r)
}

// var cookieHandler = securecookie.New(
// 	securecookie.GenerateRandomKey(64),
// 	securecookie.GenerateRandomKey(32))

// func getUserName(r *http.Request) (userName string) {
// 	if cookie, err := r.Cookie("session"); err == nil {
// 		cookieValue := make(map[string]string)
// 		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
// 			userName = cookieValue["name"]
// 		}
// 	}
// 	return userName
// }

// func setSession(userName string, w http.ResponseWriter) {
// 	value := map[string]string{
// 		"name": userName,
// 	}
// 	if encoded, err := cookieHandler.Encode("session", value); err == nil {
// 		cookie := &http.Cookie{
// 			Name:  "session",
// 			Value: encoded,
// 			Path:  "/",
// 		}
// 		http.SetCookie(w, cookie)
// 	}
// }

// func clearSession(w http.ResponseWriter) {
// 	cookie := &http.Cookie{
// 		Name:   "session",
// 		Value:  "",
// 		Path:   "/",
// 		MaxAge: -1,
// 	}
// 	http.SetCookie(w, cookie)
// }

// func LoginHandler(w http.ResponseWriter, r *http.Request) {
// 	name := r.FormValue("name")
// 	pass := r.FormValue("password")
// 	redirectTarget := "/login"

// 	res, err := http.PostForm(server+"/client/login", url.Values{"name": {name}, "password": {pass}})

// 	if err != nil {
// 		log.Errorf("Could not write response to client: %s", err)
// 		http.Redirect(w, r, redirectTarget, 302)
// 	}
// 	defer res.Body.Close()

// 	body, err := ioutil.ReadAll(res.Body)
// 	if string(body[:]) == "pass" {
// 		setSession(name, w)
// 		redirectTarget = "/"
// 	}
// 	http.Redirect(w, r, redirectTarget, 302)
// }

// func SignUpHandler(w http.ResponseWriter, r *http.Request) {
// 	name := r.FormValue("name")
// 	pass := r.FormValue("password")
// 	passconfirm := r.FormValue("passwordconfirm")
// 	redirectTarget := "/register"

// 	res, err := http.PostForm(server+"/client/register",
// 		url.Values{"name": {name}, "password": {pass}, "passwordconfirm": {passconfirm}})

// 	if err != nil {
// 		log.Errorf("Could not send passw to server: %s", err)
// 		http.Redirect(w, r, redirectTarget, 302)
// 	}
// 	defer res.Body.Close()

// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		log.Errorf("Could read response from server: %s", err)
// 		http.Redirect(w, r, redirectTarget, 302)
// 	}

// 	if string(body[:]) == "pass" {
// 		setSession(name, w)
// 		redirectTarget = "/"
// 	}
// 	http.Redirect(w, r, redirectTarget, 302)
// }

// func LogoutHandler(w http.ResponseWriter, r *http.Request) {
// 	clearSession(w)
// 	http.Redirect(w, r, "/", 302)
// }

// func ServeLoginPage(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, "../../staticServer/static/views/login.html")
// }
// func ServeSignUpPage(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, "../../staticServer/static/views/signup.html")
// }

// func InternalPageHandler(w http.ResponseWriter, r *http.Request) {
// 	userName := getUserName(r)
// 	if userName != "" {
// 		http.FileServer(http.Dir(filePath))
// 	} else {
// 		http.Redirect(w, r, "/login", 302)
// 	}
// }
