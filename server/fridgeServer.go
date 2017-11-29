package server

import (
	"net/http"

	"github.com/ASeegull/SmartFridge/server/config"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

var upgrader websocket.Upgrader

//Run starts server
func Run(cfg config.ServerConfig) error {
	upgrader = websocket.Upgrader{
		ReadBufferSize:  cfg.ReadBufferSize,
		WriteBufferSize: cfg.WriteBufferSize,
	}

	log.Printf("Server started on %s:%s", cfg.Host, cfg.Port)
	return http.ListenAndServe(cfg.Host+":"+cfg.Port, newRouter())
}

func newRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/agent", agentAuthentication).Methods("POST")
	router.HandleFunc("/agent", createWS).Methods("GET")

	sub := router.PathPrefix("/client").Subrouter()

	sub.HandleFunc("/allRecipes", getRecipes).Methods("GET")
	sub.HandleFunc("/searchRecipes", searchRecipes).Methods("POST")
	sub.HandleFunc("/fridgeContent", getFoodInfo).Methods("POST")

	sub.HandleFunc("/addAgent", addAgent).Methods("POST")
	sub.HandleFunc("/removeAgent", removeAgent).Methods("POST")
	sub.HandleFunc("/updateAgent", updateAgent).Methods("POST")

	sub.HandleFunc("/register", clientRegister).Methods("POST")
	sub.HandleFunc("/login", clientLogin).Methods("POST")
	sub.HandleFunc("/logout", clientLogout).Methods("POST")

	return router
}
