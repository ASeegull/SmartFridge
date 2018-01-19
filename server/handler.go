package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
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
	UserID string
	sync.Mutex
	sync.WaitGroup
	Conn     *websocket.Conn
	Setup    *pb.Setup
	Shutdown chan struct{}
}

// agentsList contains all connected Agents for quick access to them
var agentsList = make(map[string]*Container)

// createWS opens connection with agent and keeps it until the sigint is recieved
func createWS(w http.ResponseWriter, r *http.Request) {
	agentID := r.Header.Get("agentID")

	if database.CheckAgentRegistration(agentID) {
		if err := database.RegisterNewAgent(agentID); err != nil {
			sendResponse(w, http.StatusInternalServerError, err)
			return
		}
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	log.Infof("websocket connection with %s established", r.Host)

	agent := &Container{Conn: conn}
	agentsList[agentID] = agent

	agent.Add(1)
	go agent.ReadAgentState()
	agent.Wait()
}

// SendAgentSetup takes new settings and sends them to the specified agent via existing websocket connection
func SendAgentSetup(id string, settings *pb.Setup) error {
	agent, ok := agentsList[id]
	if !ok {
		return errors.New("bad id")
	}
	var msg []byte
	var err error

	agent.Lock()
	if agent.Setup == nil {
		agent.Setup = settings
	} else {
		if settings.AgentID != "" {
			agent.Setup.AgentID = settings.AgentID
		}
		if settings.ProductID != "" {
			agent.Setup.ProductID = settings.ProductID
		}
		if settings.Token != "" {
			agent.Setup.Token = settings.Token
		}
		if settings.UserID != "" {
			agent.Setup.UserID = settings.UserID
		}
		if settings.StateExpires != "" {
			agent.Setup.StateExpires = settings.StateExpires
		}
		if settings.Heartbeat != 0 {
			agent.Setup.Heartbeat = settings.Heartbeat
		}
	}

	msg, err = agent.Setup.MarshalStruct()
	if err != nil {
		agent.Unlock()
		return err
	}
	agent.Unlock()

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

			if t == websocket.CloseMessage {
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

			if err = database.SaveState(agentState); err != nil {
				log.Error(err)
				return
			}
			log.Infof("agent state: %v", agentState)
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

	agent := &database.NewAgent{}

	if err := json.NewDecoder(r.Body).Decode(&agent); err != nil {
		sendResponse(w, http.StatusUnauthorized, err)
		return
	}
	agent.ProductName = strings.ToLower(agent.ProductName)
	if err := database.CheckProductName(agent.ProductName); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	if err := database.RegisterAgentWithUser(userID, agent.ID); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	state := &pb.Setup{}
	state.SetParameters(agent.ID, userID, agent.ProductName, defaultHeartBeat)
	state.StateExpires = agent.StateExpires
	if err := SendAgentSetup(agent.ID, state); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}
	sendResponse(w, http.StatusOK, err)
}

func removeAgent(w http.ResponseWriter, r *http.Request) {

}

func updateAgent(w http.ResponseWriter, r *http.Request) {
	agentUpdate := &database.ProductUpdate{}
	if err := json.NewDecoder(r.Body).Decode(agentUpdate); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}
	agentID, err := database.GetAgentIDFromSerial(agentUpdate.AgentID)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	if agentUpdate.Product == "" {

		if err = database.DeleteAgent(agentID); err != nil {
			sendResponse(w, http.StatusInternalServerError, err)
			return
		}
		sendResponse(w, http.StatusOK, nil)
		return
	}
	resp := &pb.Setup{ProductID: agentUpdate.Product, StateExpires: agentUpdate.StateExpires}
	if err := SendAgentSetup(agentUpdate.AgentID, resp); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}
	sendResponse(w, http.StatusOK, nil)
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

func getRecipesByProductName(w http.ResponseWriter, r *http.Request) {
	productName := mux.Vars(r)["name"]

	if err := database.CheckProductName(productName); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

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
