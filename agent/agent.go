package agent

import (
	"context"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"

	pb "github.com/ASeegull/SmartFridge/protoStruct"
	log "github.com/sirupsen/logrus"
)

// Controls is the one struct to rule them all :)
type controls struct {
	wg          sync.WaitGroup
	conn        *websocket.Conn
	ID          string
	publicToken string
	setup       *pb.Setup
}

// Start runs agent
func Start(ctx context.Context, cfg *Config) error {
	agent := &controls{ID: cfg.AgentID, publicToken: cfg.PublicToken}

	// Get endpoints from config and make request to server to register agent entity
	wsURL := cfg.GetEndPoints()

	// Establish ws connection
	dialer := websocket.Dialer{ReadBufferSize: 1024 * 4, WriteBufferSize: 1024 * 4}
	conn, resp, err := dialer.Dial(wsURL, nil)
	if err != nil || resp.StatusCode != http.StatusSwitchingProtocols {
		return errors.Wrapf(err, "could not establish ws connection on %s. Status: %s", wsURL, resp.Status)
	}

	agent.conn = conn
	defer agent.conn.Close()
	defer resp.Body.Close()

	// Start listening and writing on connection
	agent.wg.Add(2)
	messages := make(chan *pb.Setup)
	go streamAgentState(ctx, agent, messages)
	go messageReader(ctx, agent, messages)
	agent.wg.Wait()
	return nil
}

func readPreviousState(id string) (*pb.Agentstate, error) {
	state := &pb.Agentstate{}
	path := "/agentState/" + id + ".yaml"
	if _, err := os.Stat(path); err != nil {
		return nil, err
	}

	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrapf(err, "could not read yaml file %", path)
	}

	if err = yaml.Unmarshal(yamlFile, state); err != nil {
		return nil, errors.Wrapf(err, "could not decode config file %", path)
	}
	return state, nil
}

func streamAgentState(ctx context.Context, agent *controls, messages chan *pb.Setup) {
	log.Info("Starting streaming agent state")
	// Checks if agent is already configured and retrieves current info if available
	agentInfo, err := readPreviousState(agent.ID)
	if os.IsNotExist(err) {
		agentInfo = foodAgentGenerator(agent.ID, agent.publicToken)
	}

	defer agent.wg.Done()

	ticker := time.NewTicker(time.Second * time.Duration(agent.setup.Heartbeat))
	for {
		select {
		case <-ctx.Done():
			ticker.Stop()
			agent.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseGoingAway, "Agent shut down"))
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
			log.Info("Agent state sent")
		case <-messages:
			msg := <-messages
			ticker.Stop()
			state := updateAgent(agent.ID, msg)
			err = saveAgent(state)
			ticker = time.NewTicker(time.Second * time.Duration(agent.setup.Heartbeat))
		}
	}
}

func messageReader(ctx context.Context, agent *controls, messages chan *pb.Setup) {
	log.Info("Starting reading from connection")
	defer agent.wg.Done()
	for {
		select {
		case <-ctx.Done():
			close(messages)
			return
		default:
			for {
				types, message, err := agent.conn.ReadMessage()
				if types == websocket.TextMessage {
					log.Info(string(message))
					continue
				}

				if err != nil {
					log.Errorf("failed to read message from server %s", err)
					break
				}

				setup := &pb.Setup{}
				if err = setup.UnmarshalToStruct(message); err != nil {
					break
				}
				log.Info("Setup message received")
				messages <- setup
			}
		}
	}
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// Generates default agent state
func foodAgentGenerator(id, key string) *pb.Agentstate {
	log.Info("Generating agent state")
	agentInfo := &pb.Agentstate{
		AgentID:      id,
		Token:        key,
		UserID:       "",
		ProductID:    "",
		Weight:       int32(r.Intn(900) + 1),
		StateExpires: time.Now().Format(time.ANSIC),
	}
	return agentInfo
}

// Forms agent state to send if new setup is received
func updateAgent(id string, setup *pb.Setup) *pb.Agentstate {
	log.Info("Updating agent state")
	agentInfo := &pb.Agentstate{
		AgentID:      id,
		Token:        setup.Token,
		UserID:       setup.UserID,
		ProductID:    setup.ProductID,
		Weight:       int32(r.Intn(900) + 1),
		StateExpires: setup.StateExpires,
	}
	return agentInfo
}

// Writes to yaml file current state of container
func saveAgent(state *pb.Agentstate) error {
	path := "/agentState/" + state.AgentID + ".yaml"
	config, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return errors.Wrap(err, "Cannot save agent state to file")
	}
	defer config.Close()

	info, err := yaml.Marshal(state)
	if err != nil {
		return errors.Wrap(err, "cannot marshal to yaml")
	}
	sum, err := config.Write(info)
	if err != nil || len(info) != sum {
		return errors.Wrap(err, "failed to write to file")
	}
	return nil
}
