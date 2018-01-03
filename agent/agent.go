package agent

import (
	"context"
	"io"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"

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

var agent *controls

const defaultHeartBeat = 10

//Start runs agent
func Start(cfg *Config, ctx context.Context) error {
	agent = &controls{}
	agent.setup = &pb.Setup{}
	var err error
	agent.tokenRequest = agentInit()

	// Get endpoints from config and make request to server to register agent entity
	setupURL, wsURL := cfg.GetEndPoints()
	//err = agentRegistration(setupURL, agent.tokenRequest)
	if err != nil {
		return errors.Wrapf(err, "could not set token for %s", setupURL)
	}

	dialer := websocket.Dialer{ReadBufferSize: 1024 * 4, WriteBufferSize: 1024 * 4}

	header := http.Header{}
	header.Add("agentID", agent.tokenRequest.AgentID)

	conn, resp, err := dialer.Dial(wsURL, header)
	if err != nil || resp.StatusCode != http.StatusSwitchingProtocols {
		return errors.Wrapf(err, "could not establish ws connection on %s. Status: %s", wsURL, resp.Status)
	}
	defer resp.Body.Close()

	// Establish ws connection
	agent.conn = conn
	defer agent.conn.Close()

	// Start listen and write on connection
	agent.wg.Add(2)
	go timeReader(ctx, agent)
	go streamAgentState(ctx, agent)
	agent.wg.Wait()
	return nil
}

func streamAgentState(ctx context.Context, agent *controls) {
	log.Info("Starting streaming agent state")
	defer agent.wg.Done()
	var timeHeart int32
	if agent.setup.Heartbeat == 0 {
		timeHeart = defaultHeartBeat
	} else {
		timeHeart = agent.setup.Heartbeat
	}

	ticker := time.NewTicker(time.Second * time.Duration(timeHeart))
	for {
		select {
		case <-ctx.Done():
			ticker.Stop()
			agent.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseGoingAway, "Agent shut down"))
			return
		case <-ticker.C:
			agentInfo := foodAgentGenerator(agent.tokenRequest.AgentID)
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
		}
	}
}

func agentInit() *pb.Request {
	id := uuid.NewV4().String()
	log.Infof("Container %s is starting", id)
	return &pb.Request{AgentID: id}
}

func timeReader(ctx context.Context, agent *controls) {
	log.Info("Starting reading from connection")
	defer agent.wg.Done()
	for {
		select {
		case <-ctx.Done():
			agent.conn.Close()
			return
		default:
			for {
				types, message, err := agent.conn.ReadMessage()
				if types == websocket.TextMessage || err == io.ErrUnexpectedEOF {
					log.Info(message)
					continue
				}
				if types == websocket.CloseMessage {
					return
				}
				if err = agent.setup.UnmarshalToStruct(message); err != nil {
					log.Errorf("failed unmarshal messages: %s", err)
					return
				}
			}
		}
	}
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func foodAgentGenerator(agentID string) *pb.Agentstate {
	log.Info("Generating agent state")
	return &pb.Agentstate{
		AgentID:      agentID,
		Token:        agent.setup.Token,
		UserID:       agent.setup.UserID,
		ProductID:    agent.setup.ProductID,
		Weight:       int32(r.Intn(900) + 1),
		StateExpires: agent.setup.StateExpires,
	}
}
