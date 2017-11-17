package main

import (
	"net/http"

	"github.com/ASeegull/SmartFridge/staticServer"
	log "github.com/sirupsen/logrus"
)

func main() {
	r := staticServer.NewRouter()
	port := staticServer.GetPort()
	log.WithField("port", port).Info("Client server is listening on")
	log.Fatal(http.ListenAndServe(port, r))
}
