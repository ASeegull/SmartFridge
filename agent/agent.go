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
	wg           sync.WaitGroup
	conn         *websocket.Conn
	tokenRequest *pb.Request
	setup        *pb.Setup
}

// Start runs agent
func Start(ctx context.Context, cfg *Config) error {
	agent := &controls{}

	// Get endpoints from config and make request to server to register agent entity
	setupURL, wsURL := cfg.GetEndPoints()

	// Checks if agent is already configured and retrieves current info if available
	state, err := readPreviousState(cfg.AgentID)
	if err == os.IsNotExist(err) {

	}

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
	messages := make(chan []byte, 1024)
	go streamAgentState(ctx, agent, messages)
	go timeReader(ctx, agent, messages)
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
		return nil, errors.Annotatef(err, "could not read yaml file %", path)
	}

	if err = yaml.Unmarshal(yamlFile, state); err != nil {
		return nil, errors.Annotatef(err, "could not decode config file %", configPath)
	}
	return config, nil
}

func streamAgentState(ctx context.Context, agent *controls, messages chan []byte) {
	log.Info("Starting streaming agent state")
	agentInfo := foodAgentGenerator(agent.tokenRequest, agent.setup)
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
			ticker.Stop()
			ticker = time.NewTicker(time.Second * time.Duration(agent.setup.Heartbeat))
		}
	}
}

func timeReader(ctx context.Context, agent *controls, messages chan []byte) {
	log.Info("Starting reading from connection")
	defer agent.wg.Done()
	for {
		select {
		case <-ctx.Done():
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
				log.Info("Agent successfully registered and configured")
				messages <- message
			}
		}
	}
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func foodAgentGenerator(tokenRequest *pb.Request, agentSetup *pb.Setup) *pb.Agentstate {
	log.Info("Generating agent state")
	agentInfo := &pb.Agentstate{
		AgentID:      tokenRequest.AgentID,
		Token:        agentSetup.Token,
		UserID:       agentSetup.UserID,
		ProductID:    agentSetup.ProductID,
		Weight:       int32(r.Intn(900) + 1),
		StateExpires: time.Now().Format(time.ANSIC),
	}
	return agentInfo
}
