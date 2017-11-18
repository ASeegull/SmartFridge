package server

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ASeegull/SmartFridge/server/database"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

func getHash(stringToHash string) string {
	hash := md5.Sum([]byte(stringToHash))
	return hex.EncodeToString(hash[:])
}

var upgrader websocket.Upgrader

func setUpgrader() {
	upgrader = websocket.Upgrader{
		ReadBufferSize:  serverConfig.ReadBufferSize,
		WriteBufferSize: serverConfig.WriteBufferSize,
	}
}

func agentAuthentication(w http.ResponseWriter, r *http.Request) {
	tokenExample := getHash("userID+agentID+productName")
	err := json.NewEncoder(w).Encode(tokenExample)

	if err != nil {
		log.Print(err)
	}
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
		time.Sleep(time.Second * serverConfig.WebsocketSleep)
		//here will be reading info from websocket
	}
}

type userID struct {
	id string
}

func getFoodInfo(w http.ResponseWriter, r *http.Request) {
	var u userID
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}

	defer r.Body.Close()

	IDs, err := database.GetAllAgentsIDForClient(u.id)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 404)
		return
	}

	var foods []database.FoodInfo
	foods, err = database.GetFoodsInFridge(IDs)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 404)
		return
	}

	err = json.NewEncoder(w).Encode(foods)
	if err != nil {
		log.Print(err)
	}
}

func getRecipes(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode("Here will be your recipes")
	if err != nil {
		log.Print(err)
	}
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
