package staticServer

import (
	"net/http"
	"os"
	"sync"

	"github.com/ASeegull/SmartFridge/staticServer/config"
	"github.com/davecheney/errors"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// newRouter сreates and returns gorilla router
func newRouter(staticPath string) *mux.Router {
	r := mux.NewRouter()
	staticDir := http.FileServer(http.Dir(staticPath))
	r.Handle("/", staticDir)
	r.PathPrefix("/{_:.*}").Handler(staticDir)
	return r
}

// Run starts servers for http and https connections
func Run(cfg *config.Config) {

	var wg sync.WaitGroup

	wg.Add(2)
	r := newRouter(cfg.StaticPath)
	port := os.Getenv("PORT")
	if port != "" {
		// production

		log.WithField("port", port).Info("Server started")
		log.Fatal(http.ListenAndServe(":"+port, r))
	} else {
		// dev
		serve(wg, cfg, r)
	}

	wg.Wait()
}

func serve(wg sync.WaitGroup, cfg *config.Config, r *mux.Router) {

	go func() {
		defer wg.Done()
		addr := cfg.HTTPAddr()
		log.WithField("address", addr).Info("HTTP Server started")

		if err := http.ListenAndServe(
			addr, http.HandlerFunc(redirect(cfg.HTTPSAddr())),
		); err != nil {
			log.Fatal(errors.Annotate(err, "HTTP server crushed"))
		}
	}()

	go func() {
		defer wg.Done()
		addr := cfg.HTTPSAddr()
		log.WithField("address", addr).Info("HTTPS Server started")

		if err := http.ListenAndServeTLS(
			addr, cfg.Cert, cfg.Key, r,
		); err != nil {
			log.Fatal(errors.Annotate(err, "HTTPS server crushed"))
		}
	}()
}
