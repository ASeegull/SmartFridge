package server

import (
	"sync"

	"github.com/ASeegull/SmartFridge/server/config"
	"github.com/ASeegull/SmartFridge/server/database"
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

//SetWaitGroup sets WaitGroup
func SetWaitGroup(waitGroup *sync.WaitGroup) {
	wg = waitGroup
}

//ReadConfig reads config from file
func ReadConfig() {
	var mongoDBConfig *config.MongoConfig
	if err := config.ReadConfig(); err != nil {
		serverConfig = &config.ServerConfig{
			Port:            defaultPort,
			Host:            defaultHost,
			ReadBufferSize:  defaultReadBufferSize,
			WriteBufferSize: defaultWriteBufferSize,
			WebsocketSleep:  defaultWebsocketSleep}

		log.Println("cannot read config. Used default values")
	} else {
		mongoDBConfig = config.GetMongoConfig()
		serverConfig = config.GetServerConfig()
	}

	setUpgrader()
	database.InitiateMongoDB(mongoDBConfig)
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
