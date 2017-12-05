package main

import (
	"github.com/ASeegull/SmartFridge/server"
	"github.com/ASeegull/SmartFridge/server/config"
	"github.com/ASeegull/SmartFridge/server/database"

	log "github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatalf("Cannot read config from file with error : %v", err)
		return
	}

	if err = database.InitiateMongoDB(cfg.Mongo); err != nil {
		log.Fatal(err)
		return
	}

	if err = database.InitPostgersDB(cfg.Postgres); err != nil {
		log.Fatal(err)
		return
	}

	log.Fatal(server.Run(cfg.Server))
}
