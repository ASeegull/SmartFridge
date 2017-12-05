package agent

import (
	"crypto/tls"
	"fmt"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"

	pb "github.com/ASeegull/SmartFridge/protoStruct"
	log "github.com/sirupsen/logrus"

	"bytes"
	"io/ioutil"
	"net/http"
)

var setupInfo = pb.Setup{}
var containerID string

//Start runs agent
func Start(cfg *Config) error {

	tokenURL := fmt.Sprintf("%s:%s%s", cfg.Host, cfg.Port, cfg.RestURI)
	websocketURL := fmt.Sprintf("%s:%s%s", cfg.Websocket, cfg.Port, cfg.RestURI)
	if err := setToken(tokenURL); err != nil {
		return err
	}

	dialer := websocket.Dialer{}
	dialer.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	conn, resp, err := dialer.Dial(websocketURL, nil)
	if err != nil {
		return err
	}
	defer conn.Close()
	defer resp.Body.Close()

	fmt.Println(resp.Status)

	go timeReader(conn)

	return writeInfo(conn)
}

func setToken(urlForToken string) error {
	containerID = getID()
	req := pb.Request{AgentID: containerID}
	fmt.Println("container ID: ", req.AgentID)

	data, err := req.MarshalStruct()
	if err != nil {
		return err
	}

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	response, err := client.Post(urlForToken, "application/json; charset=utf-8", bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	return setupInfo.UnmarshalToStruct(body)
}

func writeInfo(conn *websocket.Conn) error {
	agentInfo := foodAgentGenerator()

	for {
		time.Sleep(time.Second * time.Duration(setupInfo.Heatrbeat))
		agentInfo.StateExpires = int32(time.Duration(time.Now().Nanosecond()))
		date, err := agentInfo.MarshalStruct()

		if err != nil {
			return err
		}

		if err = conn.WriteMessage(websocket.BinaryMessage, date); err != nil {
			return err
		}
	}
}

func timeReader(conn *websocket.Conn) {

	for {
		types, message, err := conn.ReadMessage()

		if err != nil {
			log.Println("Error answer: ", err)
			return
		}

		if types == websocket.CloseMessage {
			log.Println(string(message))
			return
		}

		fmt.Println("answer: ", string(message))
	}
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func foodAgentGenerator() *pb.Agentstate {
	agentInfo := &pb.Agentstate{
		AgentID:      containerID,
		Token:        setupInfo.Token,
		UserID:       setupInfo.UserID,
		ProductID:    setupInfo.ProductID,
		Weight:       int32(r.Intn(900) + 1),
		StateExpires: int32(time.Now().Second()),
	}
	return agentInfo
}

func getID() string {
	return uuid.NewV4().String()
}
