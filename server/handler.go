package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/ASeegull/SmartFridge/server/database"
	"github.com/davecheney/errors"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"

	pb "github.com/ASeegull/SmartFridge/protoStruct"
)

//mock data
const (
	defaultHeartBeat = 3
	adminID          = "9079744c-ab87-4083-8400-19c14628c26f"
	clientURI        = "http://localhost:15080"
)

func sendErrorMsg(w http.ResponseWriter, err error, status int) {
	if err != nil {
		log.Error(err)
	}

	//for local work
	//w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//w.Header().Set("Access-Control-Allow-Credentials", "true")
	//w.Header().Set("Access-Control-Allow-Origin", clientURI)
	w.WriteHeader(status)
}

func checkSession(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		status, err := isNewSession(w, r)
		if err != nil {
			sendErrorMsg(w, err, http.StatusInternalServerError)
			return
		}

		if status {
			sendErrorMsg(w, err, http.StatusUnauthorized)
			return
		}

		h.ServeHTTP(w, r)
	})
}

func agentAuthentication(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		sendErrorMsg(w, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	pbRequest := pb.Request{}
	if err = pbRequest.UnmarshalToStruct(body); err != nil {
		sendErrorMsg(w, err, http.StatusInternalServerError)
		return
	}

	//get product is mock. user has to set it in server page
	productName, err := getProduct()
	if err != nil {
		sendErrorMsg(w, err, http.StatusInternalServerError)
		return
	}

	//adminID is mock. Here must be call to postgress db (table user - agentsID)
	userID := adminID
	if err := database.RegisterNewAgent(userID, pbRequest.AgentID); err != nil {
		sendErrorMsg(w, err, http.StatusInternalServerError)
		return
	}

	resp := pb.Setup{}
	resp.SetParameters(pbRequest.AgentID, userID, *productName, defaultHeartBeat)

	data, err := resp.MarshalStruct()
	if err != nil {
		sendErrorMsg(w, err, http.StatusInternalServerError)
		return
	}

	_, err = w.Write(data)
	if err != nil {
		sendErrorMsg(w, err, http.StatusInternalServerError)
		return
	}
}

//getProduct is mock. Method returns random product from postgres DB (table Products)
func getProduct() (*string, error) {

	IDs, err := database.GetAllProductsNames()
	if err != nil {
		return nil, err
	}
	rand.Seed(time.Now().UTC().UnixNano())

	return &IDs[rand.Intn(len(IDs))], nil
}

func createWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		sendErrorMsg(w, err, http.StatusInternalServerError)
		return
	}

	log.Printf("New websocket connect with %s", r.Host)
	go wsListener(conn)
}

func wsListener(conn *websocket.Conn) {
	var err error

	defer conn.Close()

	defer func() {
		if err != nil {
			if err = conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseInternalServerErr, err.Error())); err != nil {
				log.Error(err)
			}
		}
	}()

	for {
		var data []byte
		var t int

		t, data, err = conn.ReadMessage()
		if err != nil {
			log.Errorf("cannot read from websocket %v", err)
			return
		}

		if t == websocket.CloseNormalClosure {
			log.Errorf("closed ws connection with %s", conn.RemoteAddr())
			return
		}

		agentState := &pb.Agentstate{}

		if err = agentState.UnmarshalToStruct(data); err != nil {
			log.Errorf("unmarshal data from websocket error: %v", err)
			return
		}

		if !agentState.CheckToken() {
			log.Error(errors.New("unauthorized agent detected"))
			return
		}

		if err = database.SaveState(agentState); err != nil {
			log.Error(errors.New("save to db problem: "))
			return
		}

		if err := conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%s %s", time.Now().Format("2-Mon-Jan-2006-15:04:05"), "Ok!"))); err != nil {
			log.Error(err)
			return
		}
	}
}

func getFoodInfo(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		sendErrorMsg(w, err, http.StatusInternalServerError)
		return
	}

	ID := &database.UserID{ID: userID}

	foods, err := ID.GetFoodsInFridge()
	if err != nil {
		sendErrorMsg(w, err, http.StatusInternalServerError)
		return
	}

	//for local work
	//w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//w.Header().Set("Access-Control-Allow-Credentials", "true")
	//w.Header().Set("Access-Control-Allow-Origin", clientURI)
	if err = json.NewEncoder(w).Encode(foods); err != nil {
		sendErrorMsg(w, err, http.StatusInternalServerError)
	}
}

func getRecipes(w http.ResponseWriter, r *http.Request) {
	recipes, err := database.AllRecipes()

	if err != nil {
		sendErrorMsg(w, err, http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(recipes)
	if err != nil {
		sendErrorMsg(w, err, http.StatusInternalServerError)
		return
	}

	//for local work
	//w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//w.Header().Set("Access-Control-Allow-Credentials", "true")
	//w.Header().Set("Access-Control-Allow-Origin", clientURI)
	if _, err = w.Write(data); err != nil {
		sendErrorMsg(w, err, http.StatusInternalServerError)
	}
}

func searchRecipes(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		sendErrorMsg(w, err, http.StatusInternalServerError)
		return
	}

	ID := &database.UserID{ID: userID}

	foods, err := ID.GetFoodsInFridge()
	if err != nil {
		sendErrorMsg(w, err, http.StatusInternalServerError)
		return
	}

	recipes, err := database.Recipes(foods)
	if err != nil {
		sendErrorMsg(w, err, http.StatusInternalServerError)
		return
	}

	//for local work
	//w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//w.Header().Set("Access-Control-Allow-Credentials", "true")
	//w.Header().Set("Access-Control-Allow-Origin", clientURI)
	if err = json.NewEncoder(w).Encode(recipes); err != nil {
		sendErrorMsg(w, err, http.StatusInternalServerError)
	}
}

func addAgent(w http.ResponseWriter, r *http.Request) {

}

func removeAgent(w http.ResponseWriter, r *http.Request) {

}

func updateAgent(w http.ResponseWriter, r *http.Request) {

}

func clientLogin(w http.ResponseWriter, r *http.Request) {
	user := &database.Login{}
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		sendErrorMsg(w, err, http.StatusInternalServerError)
		return
	}

	if err := database.ClientLogin(user.UserName, user.Pass); err != nil {
		sendErrorMsg(w, err, http.StatusUnauthorized)
		return
	}

	userID, err := database.GetUserID(user.UserName)
	if err != nil {
		sendErrorMsg(w, err, http.StatusInternalServerError)
		return
	}

	if err := sessionSet(w, r, userID); err != nil {
		sendErrorMsg(w, err, http.StatusInternalServerError)
		return
	}
	sendErrorMsg(w, nil, http.StatusOK)
}

func clientLogout(w http.ResponseWriter, r *http.Request) {
	if err := closeSession(w, r); err != nil {
		sendErrorMsg(w, err, http.StatusInternalServerError)
		return
	}
	sendErrorMsg(w, nil, http.StatusOK)
}

func clientRegister(w http.ResponseWriter, r *http.Request) {
	newUser := &database.Login{}
	if err := json.NewDecoder(r.Body).Decode(newUser); err != nil {
		sendErrorMsg(w, err, http.StatusInternalServerError)
		return
	}

	userID, err := database.RegisterNewUser(newUser.UserName, newUser.Pass)
	if err != nil {
		sendErrorMsg(w, err, http.StatusInternalServerError)
		return
	}

	if err := sessionSet(w, r, userID); err != nil {
		sendErrorMsg(w, err, http.StatusInternalServerError)
		return
	}
	sendErrorMsg(w, nil, http.StatusOK)
}
