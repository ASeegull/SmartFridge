package staticServer

import (
	"net/http"

	"github.com/ASeegull/SmartFridge/staticServer/handlers"
	config "github.com/ASeegull/SmartFridge/staticServer/staticServerConfig"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// newRouter сreates and returns gorilla router
func newRouter() *mux.Router {
	filePath, prefix := config.GetStaticPath()
	r := mux.NewRouter()
	r.Handle("/", http.FileServer(http.Dir(filePath)))
	r.PathPrefix(prefix).Handler(http.StripPrefix(prefix, http.FileServer(http.Dir(filePath))))
	r.Path("/content").HandlerFunc(handlers.Fetch("/client/fridgeContent")).Methods("GET")
	r.Path("/recipes").HandlerFunc(handlers.Fetch("/client/allRecipes")).Methods("GET")
	return r
}

// Run starts server on defined by yaml.config port
func Run() string {
	r := newRouter()
	addr := config.GetAddr()
	log.Fatal(http.ListenAndServe(addr, r))
	return addr
}
