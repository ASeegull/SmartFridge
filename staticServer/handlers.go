package staticServer

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func redirect(httpsAddr string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		target := "https://" + httpsAddr + req.URL.String()
		log.Printf("redirect to: %s", target)
		http.Redirect(w, req, target, http.StatusPermanentRedirect)
	}
}
