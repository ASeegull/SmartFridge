package main

import (
	"context"
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

	ctx, cancel := context.WithCancel(context.Background())
	sign := make(chan os.Signal)
	signal.Notify(sign, os.Interrupt)
	go func(ctx context.Context, sign chan os.Signal) {
		<-sign
		cancel()
	}(ctx, sign)

	if err = agent.Start(cfg, ctx); err != nil {
		log.Fatal(err)
	}
}
