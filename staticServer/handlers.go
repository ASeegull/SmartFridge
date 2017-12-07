package staticServer

import (
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func redirect(httpsAddr string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		target := "https://" + httpsAddr + req.URL.String()

		if len(req.URL.RawQuery) > 0 {
			target += "?" + req.URL.RawQuery
		}
		log.Printf("redirect to: %s", target)
		http.Redirect(w, req, target, http.StatusPermanentRedirect)
	}
}

func fetch(query string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		res, err := http.Get(query)
		if err != nil {
			log.Errorf("Could not get response: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Something bad happened!"))
			return
		}
		defer res.Body.Close()

		jsn, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Errorf("Could not get response: %s", err)
		}

		done, err := w.Write(jsn)
		if err != nil {
			log.Errorf("Could not write response to client: %s", err)
		}
		log.Infof("Success, written down %d bytes", done)
	}
}
