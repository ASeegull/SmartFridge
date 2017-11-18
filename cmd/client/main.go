package main

import (
	"github.com/ASeegull/SmartFridge/staticServer"
	log "github.com/sirupsen/logrus"
)

func main() {
	addr := staticServer.Run()
	log.WithField("host:port", addr).Info("Client server is listening on")
}
