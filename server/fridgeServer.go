package server

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ASeegull/SmartFridge/server/database"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"gopkg.in/yaml.v2"

	log "github.com/sirupsen/logrus"
)

const (
	defaultPort = "9000"
	defaultHost = "localhost"
)

type config struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}

var serverConfig config

func (config *config) getConf() *config {

	yamlFile, err := ioutil.ReadFile("../SmartFridge/fridgeServerConfig.yaml")

	if err != nil {
		log.WithFields(log.Fields{"error: ": err}).Info("Cannot read ../SmartFridge/fridgeServerConfig.yaml")
	}
	err = yaml.Unmarshal(yamlFile, config)

	if err != nil {
		log.WithFields(log.Fields{"error: ": err}).Info("Unmarshal yamlFile error")
	}

	return config
}

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

func init() {
	serverConfig.getConf()
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func getHash(stringToHash string) string {
	hash := md5.Sum([]byte(stringToHash))
	return hex.EncodeToString(hash[:])
}

//NewRouter creates new gorilla router
func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/agent", agentAuthentication).Methods("POST")
	router.HandleFunc("/agent", createWS).Methods("GET")

	router.HandleFunc("/client/allRecipes", getRecipes).Methods("GET")
	router.HandleFunc("/client/fridgeContent", getFoodInfo).Methods("POST")
	router.HandleFunc("/client/searchRecipes", searchRecipes).Methods("POST")

	router.HandleFunc("/client/addAgent", addAgent).Methods("POST")
	router.HandleFunc("/client/removeAgent", removeAgent).Methods("POST")
	router.HandleFunc("/client/updateAgent", updateAgent).Methods("POST")

	router.HandleFunc("/client/register", clientRegister).Methods("POST")
	router.HandleFunc("/client/login", clientLogin).Methods("POST")
	router.HandleFunc("/client/logout", clientLogout).Methods("POST")

	return router
}

func agentAuthentication(w http.ResponseWriter, r *http.Request) {
	tokenExample := getHash("userID+agentID+productName")
	json.NewEncoder(w).Encode(tokenExample)
}

func createWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.WithFields(log.Fields{"error: ": err}).Info("Problem with creation websocket")
		return
	}

	fmt.Printf("New websocket connect with %s\n", r.Host)
	go wsListener(conn)
}

func wsListener(conn *websocket.Conn) {

}

func getFoodInfo(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.WithFields(log.Fields{"error: ": err}).Info("Problem with reading request body")
		return
	}

	clientID := string(body[1 : len(body)-2])

	IDs, err := database.GetAllAgentsIDForClient(clientID)
	if err != nil {
		log.WithFields(log.Fields{"error: ": err}).Info("Problem with reading database")
		return
	}

	var foods []database.FoodInfo

	if len(IDs) != 0 {
		foods = database.GetFoodsInFridge(IDs)
	}

	json.NewEncoder(w).Encode(foods)
}

func getRecipes(w http.ResponseWriter, r *http.Request) {

}

func searchRecipes(w http.ResponseWriter, r *http.Request) {

}

func addAgent(w http.ResponseWriter, r *http.Request) {

}

func removeAgent(w http.ResponseWriter, r *http.Request) {

}

func updateAgent(w http.ResponseWriter, r *http.Request) {

}

func clientLogin(w http.ResponseWriter, r *http.Request) {

}

func clientLogout(w http.ResponseWriter, r *http.Request) {

}

func clientRegister(w http.ResponseWriter, r *http.Request) {

}
