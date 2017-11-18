package server

import (
	"sync"

	"github.com/ASeegull/SmartFridge/server/config"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

const (
	defaultPort            = "9000"
	defaultHost            = "localhost"
	defaultReadBufferSize  = 1024
	defaultWriteBufferSize = 1024
	defaultWebsocketSleep  = 2
)

var serverConfig *config.ServerConfig
var wg *sync.WaitGroup

//GetWaitGroup returns waitGroup
func GetWaitGroup() *sync.WaitGroup {
	return wg
}

//ReadConfig reads config from file
func ReadConfig() {
	if err := config.ReadConfig(); err != nil {
		serverConfig = &config.ServerConfig{
			Port:            defaultPort,
			Host:            defaultHost,
			ReadBufferSize:  defaultReadBufferSize,
			WriteBufferSize: defaultWriteBufferSize,
			WebsocketSleep:  defaultWebsocketSleep}

		log.Println("Cannot read config. Used default values")
	} else {
		serverConfig = config.GetServerConfig()
	}

	setUpgrader()
}

// GetAddr sets host and port for server
func GetAddr() (string, string) {
	return serverConfig.Host, serverConfig.Port
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
