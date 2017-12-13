package agent

import (
	"bytes"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"

	pb "github.com/ASeegull/SmartFridge/protoStruct"
	"github.com/davecheney/errors"
	log "github.com/sirupsen/logrus"
)

// Controls is the one struct to rule them all :)
type controls struct {
	conn         *websocket.Conn
	wg           *sync.WaitGroup
	stop         chan struct{}
	tokenRequest *pb.Request
	setup        *pb.Setup
}

//Start runs agent
func Start(cfg *Config, endconn chan struct{}) error {
	agent := &controls{stop: endconn}
	var err error
	// Get random agent ID and log it
	agent.tokenRequest = agentInit()

	// Get endpoints from config and make request to server to register agent entity
	setupURL, wsURL := cfg.GetEndPoints()
	agent.setup, err = agentRegistration(setupURL, agent.tokenRequest)
	if err != nil {
		return errors.Annotatef(err, "could not set token for %s", setupURL)
	}

	// Establish ws connection
	dialer := websocket.Dialer{ReadBufferSize: 1024 * 4, WriteBufferSize: 1024 * 4}
	conn, resp, err := dialer.Dial(wsURL, nil)
	if err != nil || resp.StatusCode != http.StatusSwitchingProtocols {
		return errors.Annotatef(err, "could not esteblish ws connection on %s. Status: %s", wsURL, resp.Status)
	}
	agent.conn = conn
	defer agent.conn.Close()
	defer resp.Body.Close()

	// Start listen and write on connection
	go streamAgentState(agent)
	go timeReader(agent)
	agent.wg.Wait()
	return nil
}

func agentRegistration(tokenSetupURL string, req *pb.Request) (*pb.Setup, error) {
	data, err := req.MarshalStruct()
	if err != nil {
		return nil, err
	}

	response, err := http.Post(tokenSetupURL, "application/octet-stream", bytes.NewBuffer(data))
	if err != nil {
		return nil, errors.Annotatef(err, "could not send token to %s", tokenSetupURL)
	}

	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}

	setup := &pb.Setup{}
	if err = proto.Unmarshal(body, setup); err != nil {
		return nil, err
	}

	log.Infof("Agent successfully registered and configured")
	return setup, nil
}

func streamAgentState(agent *controls) {
	agent.wg.Add(1)
	agentInfo := foodAgentGenerator(agent.tokenRequest, agent.setup)
	defer agent.wg.Done()
	ticker := time.NewTicker(time.Second * time.Duration(agent.setup.Heartbeat))
	for {
		select {
		case <-agent.stop:
			ticker.Stop()
			return
		case <-ticker.C:
			date, err := agentInfo.MarshalStruct()
			if err != nil {
				log.Errorf("Failed to marshal agent data: %s", err)
				return
			}
			if err = agent.conn.WriteMessage(websocket.BinaryMessage, date); err != nil {
				log.Errorf("Failed to send agent data: %s", err)
				return
			}

		}
	}
}

func agentInit() *pb.Request {
	id := uuid.NewV4().String()
	log.Infof("Container %s is starting", id)
	return &pb.Request{id}
}

func timeReader(agent *controls) {
	agent.wg.Add(1)
	defer agent.wg.Done()
	for {
		select {
		case <-agent.stop:
			return
		default:
			types, message, err := agent.conn.ReadMessage()
			if err != nil && types != websocket.BinaryMessage {
				log.Error(err)
				return
			}
			if types == websocket.CloseMessage {
				return
			}
			if err = agent.setup.UnmarshalToStruct(message); err != nil {
				log.Error(err)
				return
			}
		}
	}
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func foodAgentGenerator(tokenRequest *pb.Request, agentSetup *pb.Setup) *pb.Agentstate {
	agentInfo := &pb.Agentstate{
		AgentID:      tokenRequest.AgentID,
		Token:        agentSetup.Token,
		UserID:       agentSetup.UserID,
		ProductID:    agentSetup.ProductID,
		Weight:       int32(r.Intn(900) + 1),
		StateExpires: int32(time.Now().Second()),
	}
	return agentInfo
}
