package staticServer

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"time"

	config "github.com/ASeegull/SmartFridge/staticServer/staticServerConfig"
	log "github.com/sirupsen/logrus"
)

var server = config.GetDynamicServer()
var httpsAddr = config.GetHTTPSAddr()

func redirect(w http.ResponseWriter, req *http.Request) {
	target := "https://" + httpsAddr + req.URL.String()

	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}
	log.Printf("redirect to: %s", target)
	http.Redirect(w, req, target, http.StatusPermanentRedirect)
}

//Fetch request to server
func Fetch(query string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tr := &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
			TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		}

		client := &http.Client{Transport: tr}
		res, err := client.Get(server + query)

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
			return
		}

		done, err := w.Write(jsn)
		if err != nil {
			log.Errorf("Could not write response to client: %s", err)
			return
		}
		log.Infof("Success, written down %d bytes", done)
	}
}
