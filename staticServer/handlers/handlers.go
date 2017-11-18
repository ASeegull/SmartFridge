package handlers

import (
	"io/ioutil"
	"net/http"

	config "github.com/ASeegull/SmartFridge/staticServer/staticServerConfig"
	log "github.com/sirupsen/logrus"
)

var server = config.GetServerAddr()

func Fetch(query string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := http.Get(server + query)
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
