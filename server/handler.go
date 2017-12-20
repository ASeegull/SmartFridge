package server

import (
	"encoding/json"
	"net/http"
	"os"
	"os/signal"
	"sync"

	"github.com/ASeegull/SmartFridge/server/database"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	pb "github.com/ASeegull/SmartFridge/protoStruct"
)

// mock data
const (
	defaultHeartBeat = 3
	adminID          = "9079744c-ab87-4083-8400-19c14628c26f"
)

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

func sendErrorMsg(w http.ResponseWriter, err error, status int) {
	log.Error(err)
	http.Error(w, err.Error(), status)
}

func checkSession(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if status, err := checkOutUser(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			if status == true {
				h.ServeHTTP(w, r)
			} else {
				http.Error(w, err.Error(), http.StatusUnauthorized)
			}
		}
	})
}

// createWS opens connection with agent and keeps it until the sigint is recieved
func createWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		sendErrorMsg(w, err, http.StatusInternalServerError)
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
	if err != nil || t != websocket.BinaryMessage {
		return "", errors.Wrapf(err, "cannot read from websocket %v")
	}

	agentState := &pb.Agentstate{}

	if err = agentState.UnmarshalToStruct(data); err != nil {
		return "", errors.Wrapf(err, "unmarshal data from websocket error: %v")
	}

	if agentState.Token == database.PublicToken {
		if err = database.RegisterNewAgent(agentState.AgentID); err != nil {
			return "", err
		}
	} else {

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
		return errors.Wrapf(err, "failed to marshal setup")
	}
	return agent.Conn.WriteMessage(websocket.BinaryMessage, msg)
}

// ReadAgentState reads messages from agent and saves state to mongodb
func (c *Container) ReadAgentState() {
	defer c.Done()
	defer c.Conn.Close()
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
				log.Error(errors.Wrap(err, "failed to unmarshal agentstate"))
				return
			}

			if !agentState.CheckToken() {
				log.Error(errors.New("unauthorized agent detected"))
				return
			}

			log.Infof("agent state: %v", agentState)
			if err = database.SaveState(agentState); err != nil {
				log.Error(errors.Wrap(err, "saving to db failed: "))
				return
			}
		}
	}
}

func getFoodInfo(w http.ResponseWriter, r *http.Request) {

	if r.Body == nil {
		sendErrorMsg(w, errors.New("please send a request body"), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	userID := &database.UserID{ID: adminID}

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
	user := &database.Login{}
	err := json.NewDecoder(r.Body).Decode(user)

	if err := database.ClientLogin(user.UserName, user.Pass); err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	userID, err := database.GetUserID(user.UserName)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := sessionSet(w, r, userID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func clientLogout(w http.ResponseWriter, r *http.Request) {
	if err := closeSession(w, r); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func clientRegister(w http.ResponseWriter, r *http.Request) {
	newUser := &database.Login{}
	err := json.NewDecoder(r.Body).Decode(newUser)

	userID, err := database.RegisterNewUser(newUser.UserName, newUser.Pass)
	if err != nil {
		log.Error(errors.Wrap(err, "user already exist"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := sessionSet(w, r, userID); err != nil {
		log.Error(errors.Wrap(err, "couldn't create session"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
