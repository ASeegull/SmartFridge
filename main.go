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
	}

	if err = database.InitiateMongoDB(cfg.Mongo); err != nil {
		log.Fatal(err)
	}

	if err = database.InitPostgersDB(cfg.Postgres); err != nil {
		log.Fatal(err)
	}

	log.Fatal(server.Run(cfg.Server))
}
