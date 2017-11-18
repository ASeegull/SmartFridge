package database

import (
	"crypto/tls"
	"net"

	"github.com/ASeegull/SmartFridge/server/config"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	log "github.com/sirupsen/logrus"
)

const (
	defaultURI      = "mongodb://umbrella:Djkjlz007@smartfridge-shard-00-00-rxof9.mongodb.net:27017,smartfridge-shard-00-01-rxof9.mongodb.net:27017,smartfridge-shard-00-02-rxof9.mongodb.net:27017/admin?replicaSet=SmartFridge-shard-0"
	defaultDatabase = "agent"
	defaultTable    = "agent_info"
)

var session *mgo.Session
var mongoConfig *config.MongoConfig

//InitiateMongoDB sets config for mongoDB
func InitiateMongoDB(inputConfig *config.MongoConfig) {
	if inputConfig != nil {
		mongoConfig = inputConfig
		log.Println("Used config value for mongoDB")
	} else {
		mongoConfig = &config.MongoConfig{
			URI:      defaultURI,
			Database: defaultDatabase,
			Table:    defaultTable}

		log.Println("Use default value for mongoDB")
	}
	createSession()
}

func createSession() {
	dialInfo, err := mgo.ParseURL(mongoConfig.URI)
	if err != nil {
		log.Error(err)
	}
	tlsConfig := &tls.Config{}

	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}

	session, err = mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Error(err)
	}
}

//SaveState saves state from agent
func SaveState(agentInfo *FoodAgent) error {
	c := session.DB(mongoConfig.Database).C(mongoConfig.Table)
	err := c.Insert(&agentInfo)
	if err != nil {
		log.Error(err)
	}
	return err
}

//GetFoodsInFridge shows all food in a fridge
func GetFoodsInFridge(containersID []string) ([]FoodInfo, error) {

	c := session.DB(mongoConfig.Database).C(mongoConfig.Table)

	var foods []FoodInfo
	for _, value := range containersID {
		var agent FoodAgent

		if err := c.Find(bson.M{"containerID": value}).One(&agent); err != nil {
			log.Error(err)
			return nil, err
		}

		foods = append(foods, agent.FoodInfo)
	}
	return foods, nil
}
