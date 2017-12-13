package main

import (
	"flag"
	"os"
	"os/signal"

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

	sign := make(chan os.Signal)
	endconn := make(chan struct{})
	signal.Notify(sign, os.Interrupt)
	go func() {
		<-sign
		close(endconn)
	}()

	log.Fatal(agent.Start(cfg, endconn))
}
