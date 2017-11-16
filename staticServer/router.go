package staticServer

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// NewRouter сreates and returns gorilla router
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.Handle("/", http.FileServer(http.Dir("../../staticServer/static")))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../../staticServer/static"))))
	r.Path("/content").HandlerFunc(fetch("fridgeContent")).Methods("GET")
	r.Path("/recipes").HandlerFunc(fetch("allRecipes")).Methods("GET")
	return r
}

// GetPort sets port for dev or production
func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5080"
	}
	return port
}

func fetch(query string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := http.Get("http://localhost:9000/" + query)
		if err != nil {
			log.Println("Could not get response", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Something bad happened!"))
			return
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
