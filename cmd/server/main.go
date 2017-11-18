package main

import (
	"net/http"
	"sync"

	"github.com/ASeegull/SmartFridge/server"
	log "github.com/sirupsen/logrus"
)

func main() {
	var wg *sync.WaitGroup

	server.ReadConfig()
	server.SetWaitGroup(wg)
	host, port := server.GetAddr()
	router := server.NewRouter()
	log.Printf("Server started on %s:%s", host, port)
	log.Fatal(http.ListenAndServe(host+":"+port, router))
	wg.Wait()
}
