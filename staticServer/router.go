package staticServer

import (
	"context"
	"net/http"
	"os"
	"sync"

	config "github.com/ASeegull/SmartFridge/staticServer/config"
	"github.com/davecheney/errors"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Saves http methods to constants with short names
const (
	GET = http.MethodGet
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
