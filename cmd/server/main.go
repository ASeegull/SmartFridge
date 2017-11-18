package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ASeegull/SmartFridge/server"
)

func main() {
	host, port := server.GetAddr()
	router := server.NewRouter()
	fmt.Printf("Server started with address %s:%s\n", host, port)
	log.Fatal(http.ListenAndServe(host+":"+port, router))
}
