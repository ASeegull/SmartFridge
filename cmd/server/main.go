package main

import (
	"net/http"

	"github.com/ASeegull/SmartFridge/server"
	"github.com/ASeegull/SmartFridge/server/database"
	log "github.com/sirupsen/logrus"
)

func main() {
	server.ReadConfig()
	if err := database.InitiateMongoDB(); err != nil {
		log.Println(err)
		return
	}

	host, port := server.GetAddr()
	router := server.NewRouter()

	defer server.GetWaitGroup().Wait()

	log.Printf("Server started on %s:%s", host, port)
	log.Fatal(http.ListenAndServe(host+":"+port, router))
}
