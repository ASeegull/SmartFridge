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

// AgentsList contains all connected Agents for quick access to them

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

//func agentAuthentication(w http.ResponseWriter, r *http.Request) {
//	//t, data, err := conn.ReadMessage()
//	//if err != nil {
//	//	return "", err
//	//}
//	//if t != websocket.BinaryMessage {
//	//	return "", errors.New("It is not BinaryMessage")
//	//}
//
//	data, err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		sendResponse(w, http.StatusInternalServerError, err)
//		return
//	}
//	request := &pb.Request{}
//
//	if err = request.UnmarshalToStruct(data); err != nil {
//		sendResponse(w, http.StatusInternalServerError, err)
//		return
//	}
//	fmt.Println("agentID", request.AgentID)
//
//	if database.CheckAgentRegistration(request.AgentID) {
//		if err = database.RegisterNewAgent(request.AgentID); err != nil {
//			sendResponse(w, http.StatusInternalServerError, err)
//			return
//		}
//	}
//
//	sendResponse(w, http.StatusOK, nil)
//}

// SendAgentSetup takes new settings and sends them to the specified agent via existing websocket connection
func SendAgentSetup(id string, settings *pb.Setup) error {
	agent, ok := agentsList[id]
	if !ok {
		return errors.New("bad id")
	}

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
	agent.ProductName = strings.ToLower(agent.ProductName)
	if err := database.CheckProductName(agent.ProductName); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	if err := database.RegisterAgentWithUser(userID, agent.ID); err != nil {
		sendResponse(w, http.StatusUnauthorized, err)
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
	type updateAgent struct {
		Product      string `json:"product"`
		Weight       int    `json:"weight"`
		StateExpires string `json:"statExpires"`
		Condition    string `json:"condition"`
		ImageURL     string `json:"imageURL"`
	}
	agentUpdate := &updateAgent{}
	if err := json.NewDecoder(r.Body).Decode(agentUpdate); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}
	if agentUpdate.Product == "" {
		agentID, err := database.GetAgentIDFromSerial("e1379bc0-dce7-46e8-9bd0-a19ffa1f3bad")
		if err != nil {
			sendResponse(w, http.StatusInternalServerError, err)
			return
		}
		err = database.DeleteAgent(agentID)
		if err != nil {
			sendResponse(w, http.StatusInternalServerError, err)
			return
		}
		sendResponse(w, http.StatusOK, nil)
		return
	}
	//userID, err := getUserID(r)
	//if err != nil {
	//	sendResponse(w, http.StatusInternalServerError, err)
	//	return
	//}
	//if err := database.CheckAgent(userID, agentUpdate.AgentID); err != nil {
	//	sendResponse(w, http.StatusUnauthorized, err)
	//	return
	//}
	//resp := pb.Setup{}
	//resp.UpdateParameters(agentUpdate.Product, agentUpdate.StateExpires, defaultHeartBeat)
	//
	////if err:= SendAgentSetup(agentUpdate.AgentID, &resp); err!=nil {
	////	sendResponse(w, http.StatusInternalServerError, err)
	////return
	////}
	//data, err := resp.MarshalStruct()
	//if err != nil {
	//	sendResponse(w, http.StatusInternalServerError, err)
	//	return
	//}
	//
	//_, err = w.Write(data)
	//if err != nil {
	//	sendResponse(w, http.StatusInternalServerError, err)
	//}

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
