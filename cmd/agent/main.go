package main

import (
	"flag"

	"github.com/ASeegull/SmartFridge/agent"
	log "github.com/sirupsen/logrus"
)

func main() {
	cfgPath := flag.String("config", "agent/config.yaml", "Location of config File")

	flag.Parse()
	cfg, err := agent.ReadConfig(*cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(agent.Start(cfg))
}
