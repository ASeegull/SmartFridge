package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/ASeegull/SmartFridge/server/database"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"

	pb "github.com/ASeegull/SmartFridge/protoStruct"
)

//mock data
const (
	defaultHearBeat = 3
	adminID         = "9079744c-ab87-4083-8400-19c14628c26f"
)

func sendErrorMsg(w http.ResponseWriter, err error, status int) {
	log.Debug(err)
	http.Error(w, err.Error(), status)
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
	resp.SetParameters(pbRequest.AgentID, userID, *productName, defaultHearBeat)

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

	IDs, err := database.GetAllProductsName()
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
				log.Debug(err)
			}
		}
	}()

	for {
		var data []byte
		var t int

		t, data, err = conn.ReadMessage()
		if err != nil {
			log.Printf("cannot read from websocket %v", err)
			return
		}

		if t == websocket.CloseNormalClosure {
			log.Printf("closed ws connection with %s", conn.RemoteAddr())
			return
		}

		agentState := &pb.Agentstate{}

		if err = agentState.UnmarshalToStruct(data); err != nil {
			log.Printf("unmarshal data from websocket error: %v", err)
			return
		}

		if !agentState.CheckToken() {
			err = errors.New("unauthorized agent detected")
			log.Println(err)
			return
		}

		if err = database.SaveState(agentState); err != nil {
			err = errors.New("save to db problem: ")
			log.Println(err)
			return
		}

		if err := conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%s %s", time.Now().Format("2-Mon-Jan-2006-15:04:05"), "Ok!"))); err != nil {
			log.Debug(err)
			return
		}
	}
}

func getFoodInfo(w http.ResponseWriter, r *http.Request) {

	if r.Body == nil {
		sendErrorMsg(w, errors.New("please send a request body"), http.StatusBadRequest)
		return
	}

	userID := &database.UserID{adminID}

	defer r.Body.Close()

	foods, err := userID.GetFoodsInFridge()
	if err != nil {
		sendErrorMsg(w, err, http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(foods); err != nil {
		sendErrorMsg(w, err, http.StatusInternalServerError)
	}
}

func getRecipes(w http.ResponseWriter, r *http.Request) {
	recipes := database.RecipesStr{}

	if err := recipes.GetAllRecipes(); err != nil {
		sendErrorMsg(w, err, http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(recipes)
	if err != nil {
		sendErrorMsg(w, err, http.StatusInternalServerError)
		return
	}

	if _, err = w.Write(data); err != nil {
		sendErrorMsg(w, err, http.StatusInternalServerError)
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
	defer r.Body.Close()

	jsn, err := ioutil.ReadAll(r.Body)

	if err != nil {
		sendErrorMsg(w, err, http.StatusInternalServerError)
		return
	}

	login := database.Login{}
	if err = login.Unmarshalling(jsn); err != nil {
		sendErrorMsg(w, err, http.StatusInternalServerError)
		return
	}

	if err = login.LogIn(); err != nil {
		sendErrorMsg(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func clientLogout(w http.ResponseWriter, r *http.Request) {

}

func clientRegister(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	jsn, err := ioutil.ReadAll(r.Body)

	if err != nil {
		sendErrorMsg(w, err, http.StatusInternalServerError)
		return
	}

	login := database.Login{}
	if err = login.Unmarshalling(jsn); err != nil {
		sendErrorMsg(w, err, http.StatusInternalServerError)
		return
	}

	if err = login.Register(); err != nil {
		sendErrorMsg(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
