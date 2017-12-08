package main

import (
	"flag"

	"github.com/ASeegull/SmartFridge/staticServer"
	"github.com/ASeegull/SmartFridge/staticServer/config"
	"github.com/davecheney/errors"
	log "github.com/sirupsen/logrus"
)

func main() {
	cfgPath := flag.String(
		"config",
		"staticServer/config/files/clientConfig.yaml",
		"Location of config File",
	)

	log.Info(*cfgPath)
	flag.Parse()
	cfg, err := config.GetSettings(*cfgPath)
	if err != nil {
		log.Fatal(errors.Annotate(err, "Server couldn't start"))
	}

	staticServer.Run(cfg)
}
