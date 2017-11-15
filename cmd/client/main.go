package main

import (
	"fmt"
	"log"
	"net/http"

	s "github.com/ASeegull/SmartFridge/staticServer"
)

func main() {
	r := s.NewRouter()
	port := ":" + s.GetPort()
	fmt.Printf("Client server is listening on port %s", port)
	log.Fatal(http.ListenAndServe(port, r))
}
