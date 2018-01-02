package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"

	"github.com/ASeegull/SmartFridge/server/database"
	"github.com/davecheney/errors"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"

	pb "github.com/ASeegull/SmartFridge/protoStruct"
	"github.com/gorilla/mux"
)

const (
	defaultHeartBeat = 3
)

func sendResponse(w http.ResponseWriter, status int, err error) {
	if err != nil {
		log.Error(err)
		http.Error(w, fmt.Sprintf("%v", err), status)
		return
	}

	w.WriteHeader(status)
}

func checkSession(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		isNew, err := isNewSession(w, r)
		if err != nil {
			sendResponse(w, http.StatusInternalServerError, err)
			return
		}

		if isNew {
			sendResponse(w, http.StatusUnauthorized, err)
			return
		}

		h.ServeHTTP(w, r)
	})
}

// Container holds ws connection for instance of agent and chan to send done signal to goroutines
type Container struct {
	sync.Mutex
	sync.WaitGroup
	Conn     *websocket.Conn
	Setup    *pb.Setup
	Shutdown chan struct{}
}

// AgentsList contains all connected Agents for quick access to them

var agentsList = &map[string]*Container{}

// createWS opens connection with agent and keeps it until the sigint is recieved
func createWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	log.Infof("websocket connection with %s established", r.Host)

	done := make(chan struct{})
	sign := make(chan os.Signal)
	signal.Notify(sign, os.Interrupt)
	go func() {
		<-sign
		log.Info("SIGINT recieved")
		done <- struct{}{}
	}()

	id, err := agentAuthentication(conn)
	agent := &Container{Conn: conn, Shutdown: done}
	if _, ok := (*agentsList)[id]; !ok {
		(*agentsList)[id] = agent
	}

	agent.Add(1)
	go agent.ReadAgentState()
	agent.Wait()
}

func agentAuthentication(conn *websocket.Conn) (string, error) {
	t, data, err := conn.ReadMessage()
	if err != nil {
		return "", err
	}
	if t != websocket.BinaryMessage {
		return "", errors.New("It is not BinaryMessage")
	}

	agentState := &pb.Agentstate{}

	if err = agentState.UnmarshalToStruct(data); err != nil {
		return "", err
	}

	if database.CheckAgentRegistration(agentState.Token) {
		if err = database.RegisterNewAgent(agentState.AgentID); err != nil {
			return agentState.AgentID, err
		}
	}

	return agentState.AgentID, nil
}

// SendAgentSetup takes new settings and sends them to the specified agent via existing websocket connection
func SendAgentSetup(id string, settings *pb.Setup) error {
	agent := (*agentsList)[id]

	{
		agent.Lock()
		agent.Setup = settings
		agent.Unlock()
	}

	msg, err := settings.MarshalStruct()
	if err != nil {
		return err
	}
	return agent.Conn.WriteMessage(websocket.BinaryMessage, msg)
}

// ReadAgentState reads messages from agent and saves state to mongodb
func (c *Container) ReadAgentState() {
	defer c.Conn.Close()
	defer c.Done()
	for {
		select {
		case <-c.Shutdown:
			return
		default:
			t, data, err := c.Conn.ReadMessage()
			if err != nil {
				log.Errorf("cannot read from websocket %v", err)
				return
			}

			if t == websocket.CloseGoingAway {
				log.Errorf("closed ws connection with %s", c.Conn.RemoteAddr())
				return
			}

			agentState := &pb.Agentstate{}

			if err = agentState.UnmarshalToStruct(data); err != nil {
				log.Error(err)
				return
			}

			if !agentState.CheckToken() {
				log.Error(errors.New("unauthorized agent detected"))
				continue
			}

			log.Infof("agent state: %v", agentState)
			if err = database.SaveState(agentState); err != nil {
				log.Error(err)
				return
			}
		}
	}
}

func getFoodInfo(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	ID := &database.UserID{ID: userID}

	foods, err := ID.GetFoodsInFridge()
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	if err = json.NewEncoder(w).Encode(foods); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
	}
}

func getRecipes(w http.ResponseWriter, r *http.Request) {
	recipes, err := database.AllRecipes()

	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	data, err := json.Marshal(recipes)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	if _, err = w.Write(data); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
	}
}

func searchRecipes(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	ID := &database.UserID{ID: userID}

	foods, err := ID.GetFoodsInFridge()
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	recipes, err := database.Recipes(foods)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	if err = json.NewEncoder(w).Encode(recipes); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
	}
}

func addAgent(w http.ResponseWriter, r *http.Request) {

	userID, err := getUserID(r)
	if err != nil {
		sendResponse(w, http.StatusUnauthorized, err)
		return
	}

	type NewAgent struct {
		ID           string `json:"agentID"`
		ProductName  string `json:"product"`
		StateExpires string `json:"stateExpires"`
	}
	agent := &NewAgent{}

	if err := json.NewDecoder(r.Body).Decode(&agent); err != nil {
		sendResponse(w, http.StatusUnauthorized, err)
		return
	}
	if err := database.RegisterAgentWithUser(userID, agent.ID); err != nil {
		sendResponse(w, http.StatusUnauthorized, err)
		return
	}

	if err := database.CheckProductName(agent.ProductName); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}
	state := &pb.Setup{}
	state.SetParameters(agent.ID, userID, agent.ProductName, defaultHeartBeat)
	if err := SendAgentSetup(agent.ID, state); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}
	sendResponse(w, http.StatusOK, err)
	//--------------------------------------------------------------------------------

	//userId, err := getUserID(r)
	//if err != nil {
	//	sendResponse(w, http.StatusInternalServerError, err)
	//	return
	//}
	//if err := json.NewDecoder(r.Body).Decode(agent); err != nil {
	//	sendResponse(w, http.StatusInternalServerError, err)
	//	return
	//}
	//if err := database.CheckAgent(userId, agent.ID); err != errors.New("unregistered agent") {
	//	sendResponse(w, http.StatusInternalServerError, err)
	//	return
	//}
	//
	//if err := database.RegisterAgentWithUser(userId, agent.ID); err != nil {
	//	sendResponse(w, http.StatusInternalServerError, err)
	//	return
	//}
	//w.WriteHeader(http.StatusOK)
}

func removeAgent(w http.ResponseWriter, r *http.Request) {

}

func updateAgent(w http.ResponseWriter, r *http.Request) {

}

func clientLogin(w http.ResponseWriter, r *http.Request) {
	user := &database.Login{}
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	if err := database.ClientLogin(user.UserName, user.Pass); err != nil {
		sendResponse(w, http.StatusUnauthorized, err)
		return
	}

	userID, err := database.GetUserID(user.UserName)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	if err := sessionSet(w, r, userID); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	sendResponse(w, http.StatusOK, nil)
}

func clientLogout(w http.ResponseWriter, r *http.Request) {
	if err := closeSession(w, r); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	sendResponse(w, http.StatusOK, nil)
}

func clientRegister(w http.ResponseWriter, r *http.Request) {
	newUser := &database.Login{}
	if err := json.NewDecoder(r.Body).Decode(newUser); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	userID, err := database.RegisterNewUser(newUser.UserName, newUser.Pass)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	if err := sessionSet(w, r, userID); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	sendResponse(w, http.StatusOK, nil)
}

func productAdd(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}
	isAdmin, err := checkAdmin(userID)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}
	if !isAdmin {
		sendResponse(w, http.StatusMethodNotAllowed, errors.New("not allowed"))
		return
	}
	newProduct := &database.Product{}

	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	if err := database.AddProduct(newProduct); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	sendResponse(w, http.StatusOK, nil)
}

func getAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := database.AllProducts()
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	data, err := json.Marshal(products)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	if _, err = w.Write(data); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
	}
}

func getProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["id"]
	product, err := database.FindProductByID(ID)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	data, err := json.Marshal(product)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	if _, err = w.Write(data); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
	}
}

func getProductByName(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	product, err := database.FindProductByName(name)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	data, err := json.Marshal(product)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	if _, err = w.Write(data); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
	}
}

func productUpdate(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}
	isAdmin, err := checkAdmin(userID)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	if !isAdmin {
		sendResponse(w, http.StatusMethodNotAllowed, errors.New("not allowed"))
		return
	}
	newProduct := &database.Product{}

	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	if err := database.UpdateProduct(newProduct.ID, newProduct.Name, newProduct.Image, newProduct.ShelfLife, newProduct.Units); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	sendResponse(w, http.StatusOK, nil)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}
	isAdmin, err := checkAdmin(userID)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	if !isAdmin {
		sendResponse(w, http.StatusMethodNotAllowed, errors.New("not allowed"))
		return
	}
	ID := mux.Vars(r)["id"]

	if err := database.DeleteProductByID(ID); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	sendResponse(w, http.StatusOK, nil)
}

func getRecipesByProductName(w http.ResponseWriter, r *http.Request) {
	productName := mux.Vars(r)["name"]
	recipes, err := database.GetRecepiesByProductName(productName)

	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	data, err := json.Marshal(recipes)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	if _, err = w.Write(data); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
	}
}

func recipesByProductNames(w http.ResponseWriter, r *http.Request) {
	var productNames []string
	err := json.NewDecoder(r.Body).Decode(&productNames)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	recipes, err := database.RecepiesByProducts(productNames)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	data, err := json.Marshal(recipes)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	if _, err = w.Write(data); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
	}
}

func checkAdmin(userID string) (bool, error) {
	return database.CheckAdmin(userID)
}
