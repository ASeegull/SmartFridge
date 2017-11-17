package server

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/ASeegull/SmartFridge/server/database"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

var wg sync.WaitGroup

func getHash(stringToHash string) string {
	hash := md5.Sum([]byte(stringToHash))
	return hex.EncodeToString(hash[:])
}

func agentAuthentication(w http.ResponseWriter, r *http.Request) {
	tokenExample := getHash("userID+agentID+productName")
	json.NewEncoder(w).Encode(tokenExample)
}

func createWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.WithField("error: ", err).Info("Problem with creation websocket")
		return
	}

	fmt.Printf("New websocket connect with %s\n", r.Host)
	wg.Add(1)
	go wsListener(conn)
}

func wsListener(conn *websocket.Conn) {
	defer wg.Done()
	defer conn.Close()

	for {
		fmt.Println("here will be reading info from websocket")
	}
}

type userId struct {
	id string
}

func getFoodInfo(w http.ResponseWriter, r *http.Request) {
	var u userId
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	defer r.Body.Close()

	IDs, err := database.GetAllAgentsIDForClient(u.id)
	if err != nil {
		log.Error(err)
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
