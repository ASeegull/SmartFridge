package main

import (
	"log"

	"github.com/ASeegull/SmartFridge/agent"
)

func main() {
	cfg, err := agent.ReadConfig()
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Fatal(agent.Start(cfg))
}
