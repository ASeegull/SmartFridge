package server

import (
	"github.com/ASeegull/SmartFridge/server/config"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

const (
	defaultPort = "9000"
	defaultHost = "localhost"
)

var serverConfig = config.GetServerConfig()

// GetAddr sets host and port for server
func GetAddr() (host string, port string) {
	host = serverConfig.Host
	port = serverConfig.Port

	if port == "" {
		port = defaultPort
	}

	if host == "" {
		host = defaultHost
	}

	return host, port
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  serverConfig.ReadBufferSize,
	WriteBufferSize: serverConfig.WriteBufferSize,
}

//NewRouter creates new gorilla router
func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/agent", agentAuthentication).Methods("POST")
	router.HandleFunc("/agent", createWS).Methods("GET")

	s := router.PathPrefix("/client").Subrouter()

	s.HandleFunc("/allRecipes", getRecipes).Methods("GET")
	s.HandleFunc("/searchRecipes", searchRecipes).Methods("POST")
	s.HandleFunc("/fridgeContent", getFoodInfo).Methods("POST")

	s.HandleFunc("/addAgent", addAgent).Methods("POST")
	s.HandleFunc("/removeAgent", removeAgent).Methods("POST")
	s.HandleFunc("/updateAgent", updateAgent).Methods("POST")

	s.HandleFunc("/register", clientRegister).Methods("POST")
	s.HandleFunc("/login", clientLogin).Methods("POST")
	s.HandleFunc("/logout", clientLogout).Methods("POST")

	return router
}
