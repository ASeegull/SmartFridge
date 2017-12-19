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
	agentID := flag.String("agentID", "12345", "Set agent ID if it's a new instance")

	flag.Parse()

	cfg, err := agent.ReadConfig(*cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	cfg.AgentID = *agentID

	ctx, cancel := context.WithCancel(context.Background())
	sign := make(chan os.Signal)
	signal.Notify(sign, os.Interrupt)
	go func() {
		<-sign
		log.Info("SIGINT recieved")
		cancel()
	}()

	if err = agent.Start(ctx, cfg); err != nil {
		log.Fatal(err)
	}
}
