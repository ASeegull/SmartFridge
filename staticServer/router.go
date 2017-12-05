package staticServer

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"sync"

	config "github.com/ASeegull/SmartFridge/staticServer/config"
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

// newRouter сreates and returns gorilla router
func newRouter(staticPath string) *mux.Router {
	r := mux.NewRouter()
	staticDir := http.FileServer(http.Dir(staticPath))
	r.Handle("/", staticDir)
	r.PathPrefix("/{_:.*}").Handler(staticDir)
	r.Path("/login").HandlerFunc(LoginHandler).Methods(POST)
	r.Path("/signup").HandlerFunc(SignUpHandler).Methods(POST)
	r.Path("/content").HandlerFunc(Fetch("/client/fridgeContent")).Methods(GET)
	r.Path("/recipes").HandlerFunc(Fetch("/client/allRecipes")).Methods(GET)
	return r
}

// Run starts servers for http and https connections
func Run(cfg *config.Config, ctx context.Context) {

	var wg sync.WaitGroup

	wg.Add(2)

	r := newRouter(cfg.StaticPath)
	port := os.Getenv("PORT")
	switch {
	case port != "":
		log.WithField("port", port).Info("Server started")
		log.Fatal(http.ListenAndServe(":"+port, r))
	}

	go func() {
		addr := cfg.HTTPAddr()
		log.WithField("address", addr).Info("HTTP Server started")

		if err := http.ListenAndServe(
			addr, http.HandlerFunc(redirect(cfg.HTTPSAddr())),
		); err != nil {
			log.Fatal(errors.Annotate(err, "HTTP server crushed"))
		}
	}()

	go func() {
		addr := cfg.HTTPSAddr()
		log.WithField("address", addr).Info("HTTPS Server started")

		if err := http.ListenAndServeTLS(
			addr, cfg.Cert, cfg.Key, r,
		); err != nil {
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
