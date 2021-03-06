package main

import (
	"context"
	"flag"
	"fmt"
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

	ctx, cancel := context.WithCancel(context.Background())
	sign := make(chan os.Signal)
	signal.Notify(sign, os.Interrupt)
	go func() {
		<-sign
		cancel()
	}()

	go func() {
		var word string
		fmt.Scan(&word)
		cancel()
	}()

	if err = agent.Start(cfg, ctx); err != nil {
		log.Fatal(err)
	}
}
