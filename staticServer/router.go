package staticServer

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

//Creates gorilla router
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.Handle("/", http.FileServer(http.Dir("static")))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	r.Path("/content").HandlerFunc(fetch("fridgeContent")).Methods("GET")
	r.Path("/recipes").HandlerFunc(fetch("allRecipes")).Methods("GET")
	return r
}

//Sets port for dev or production
func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5080"
	}
	return port
}

func fetch(query string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		res, serverErr := http.Get("localhost:9000/" + query)
		if serverErr != nil {
			log.Println("Could not get response", serverErr)
		}
		defer res.Body.Close()
		jsn, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			log.Println("Could not get response", readErr)
		}
		_, writeErr := w.Write(jsn)
		if writeErr != nil {
			log.Println("Could not write response to client", writeErr)
		}
	}
}
