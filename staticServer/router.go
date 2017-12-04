package staticServer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	config "github.com/ASeegull/SmartFridge/staticServer/staticServerConfig"
	"github.com/davecheney/errors"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	log "github.com/sirupsen/logrus"
)

// Saves http methods to constants with short names
const (
	GET  = http.MethodGet
	POST = http.MethodPost
)

var SIGN_KEY = []byte("secret")

// newRouter сreates and returns gorilla router
func newRouter() *mux.Router {
	cfg := config.GetStaticFilesPath()
	r := mux.NewRouter()
	r.Path("/").Handler(http.FileServer(http.Dir(cfg.Path)))
	r.PathPrefix(cfg.Prefix).Handler(http.StripPrefix(cfg.Prefix, http.FileServer(http.Dir(cfg.Path))))
	r.Path("/main").HandlerFunc(ServeInternal).Methods(GET)
	r.Path("/login").HandlerFunc(LoginHandler).Methods(POST)
	r.Path("/signup").HandlerFunc(SignUpHandler).Methods(POST)
	r.Path("/content").HandlerFunc(Fetch("/client/fridgeContent")).Methods(GET)
	r.Path("/recipes").HandlerFunc(Fetch("/client/allRecipes")).Methods(GET)
	return r
}

// Run starts servers for http and https connections
func Run() {
	var wg sync.WaitGroup
	wg.Add(2)

	var err error
	r := newRouter()
	go func() {
		addr := config.GetHTTPAddr()
		log.WithField("address", addr).Info("HTTP Server started")

		if err = http.ListenAndServe(addr, http.HandlerFunc(redirect)); err != nil {
			log.Fatal(errors.Annotate(err, "HTTP server crushed"))
		}
	}()
	go func() {
		cfg := config.GetStaticHTTPScfg()
		addr := config.GetHTTPSAddr()
		log.WithField("address", addr).Info("HTTPS Server started")

		if err = http.ListenAndServeTLS(addr, cfg.Cert, cfg.Key, r); err != nil {
			log.Fatal(errors.Annotate(err, "HTTPS server crushed"))
		}
	}()
	wg.Wait()
}

func createSecureCookie() *securecookie.SecureCookie {
	hashKey, blockKey := []byte("very-secret"), []byte("a-lot-secret")
	return securecookie.New(hashKey, blockKey)
}

type signup struct {
	Name     string
	Password string
	Email    string
}

func SignUpHandler(w http.ResponseWriter, req *http.Request) {
	signup := &signup{}

	defer req.Body.Close()
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(signup)
	if err != nil {
		log.Error(errors.Annotate(err, "Failed to read signup data"))
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed to read signup data"))
		return
	}

	target := "/main"
	log.Printf("redirect to: %s", target)
	http.Redirect(w, req, target, http.StatusFound)
}

func LoginHandler(w http.ResponseWriter, req *http.Request) {
	lg := &signup{}

	defer req.Body.Close()
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(lg)
	if err != nil {
		log.Error(errors.Annotate(err, "Failed to read signup data"))
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed to read signup data"))
		return
	}
}

func ServeInternal(w http.ResponseWriter, req *http.Request) {
	fmt.Print("tried to redirect")
	http.ServeFile(w, req, "../../staticServer/static/views/main.html")
}
