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
	wg          *sync.WaitGroup
	conn        *websocket.Conn
	containerID string
	setup       *pb.Setup
	mutex       *sync.Mutex
}

var (
	MarshaledAgentState []byte
	agent               *controls
	containerState      *pb.Agentstate
	r                   *rand.Rand
	ticker              *time.Ticker
)

const defaultHeartBeat = 5

//Start runs agent
func Start(cfg *Config, ctx context.Context) error {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	containerID := getSerialID()
	log.Infof("Container %s is starting", containerID)

	dialer := websocket.Dialer{ReadBufferSize: 1024 * 4, WriteBufferSize: 1024 * 4}

	agent = &controls{wg: &sync.WaitGroup{}, containerID: containerID, setup: &pb.Setup{}, mutex: &sync.Mutex{}}

	header := http.Header{}
	header.Add("agentID", containerID)

	wsURL := cfg.GetEndPoints()
	conn, resp, err := dialer.Dial(wsURL, header)
	if err != nil || resp.StatusCode != http.StatusSwitchingProtocols {
		return errors.Wrapf(err, "could not establish ws connection on %s. Status: %s", wsURL, resp.Status)
	}
	defer resp.Body.Close()

	agent.conn = conn
	defer agent.conn.Close()

	ticker = time.NewTicker(time.Second * time.Duration(defaultHeartBeat))
	defer ticker.Stop()

	agent.wg.Add(2)
	go timeReader(ctx, agent)
	go streamAgentState(ctx, agent)
	agent.wg.Wait()

	return nil
}

func streamAgentState(ctx context.Context, agent *controls) {
	defer agent.wg.Done()
	log.Info("Starting streaming agent state")

	MarshaledAgentState = foodAgentGenerator(agent.containerID)

	for range ticker.C {
		select {
		case <-ctx.Done():
			agent.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseGoingAway, "Agent shut down"))
			log.Println("stop streaming")
			return
		default:
			agent.mutex.Lock()
			if err := agent.conn.WriteMessage(websocket.BinaryMessage, MarshaledAgentState); err != nil {
				log.Errorf("Failed to send agent data: %s", err)
				return
			}
			agent.mutex.Unlock()
			log.Info("Agent state sent")
		}
	}
}

func timeReader(ctx context.Context, agent *controls) {

	log.Info("Starting reading from connection")
	defer agent.wg.Done()

	for range ticker.C {
		select {
		case <-ctx.Done():
			log.Println("stop reading")
			return
		default:
			types, message, err := agent.conn.ReadMessage()
			if err != nil {
				log.Info(err)
			}
			if types == websocket.CloseMessage {
				return
			}
			if types == websocket.TextMessage || err == io.ErrUnexpectedEOF {
				log.Info(message)
				continue
			}
			newSetup := &pb.Setup{}
			if err = newSetup.UnmarshalToStruct(message); err != nil {
				log.Errorf("failed unmarshal messages: %s", err)
				return
			}
			log.Println(newSetup)
			agent.mutex.Lock()
			updateAgentState(newSetup)
			agent.mutex.Unlock()
		}
	}
}

func foodAgentGenerator(agentID string) []byte {
	log.Info("Generating agent state")
	containerState = &pb.Agentstate{
		AgentID:      agentID,
		Token:        agent.setup.Token,
		UserID:       agent.setup.UserID,
		ProductID:    agent.setup.ProductID,
		Weight:       int32(r.Intn(900) + 1),
		StateExpires: agent.setup.StateExpires,
	}

	date, err := containerState.MarshalStruct()
	if err != nil {
		log.Fatalf("failed unmarshal messages: %s", err)
	}
	return date
}

func updateAgentState(newSetup *pb.Setup) {
	containerState.Token = newSetup.Token
	containerState.UserID = newSetup.UserID
	containerState.ProductID = newSetup.ProductID
	containerState.StateExpires = newSetup.StateExpires

	var err error
	MarshaledAgentState, err = containerState.MarshalStruct()
	if err != nil {
		log.Fatalf("failed unmarshal messages: %s", err)
	}
}

func getSerialID() string {
	s, _ := uuid.NewV4()
	return s.String()
}
