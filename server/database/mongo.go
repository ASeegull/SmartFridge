package database

import (
	"crypto/tls"
	"io/ioutil"
	"net"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/yaml.v2"

	log "github.com/sirupsen/logrus"
)

//FoodInfo is example struct from agent
type FoodInfo struct {
	Name           string `bson:"name"`
	ExpirationDate string `bson:"exporationdate"`
	Weight         int    `bson:"weight"`
	Amount         int    `bson:"amount"`
}

//FoodAgent is example struct from agent
type FoodAgent struct {
	ClientID    string   `bson:"clientid"`
	ContainerID string   `bson:"containerid"`
	Token       string   `bson:"token"`
	FoodInfo    FoodInfo `bson:"foodinfo"`
}

type config struct {
	MongoURI string `yaml:"mongoURI"`
	Database string `yaml:"database"`
	Table    string `yaml:"table"`
}

var mongoConfig config

func (config *config) getConf() *config {

	yamlFile, err := ioutil.ReadFile("../SmartFridge/mongoConfig.yaml")
	if err != nil {
		log.WithFields(log.Fields{"error: ": err}).Info("Cannot read ../SmartFridge/mongoConfig.yaml")
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		log.WithFields(log.Fields{"error: ": err}).Info("Unmarshal yamlFile error")
	}

	return config
}

func init() {
	mongoConfig.getConf()
}

func getSession() (*mgo.Session, error) {
	var mongoURI = mongoConfig.MongoURI

	dialInfo, _ := mgo.ParseURL(mongoURI)
	tlsConfig := &tls.Config{}

	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}
	return mgo.DialWithInfo(dialInfo)
}

//SaveState saves state from agent
func SaveState(agentInfo *FoodAgent) {

	session, err := getSession()
	if err != nil {
		log.WithFields(log.Fields{"error: ": err}).Info("Cannot create session to mongoDB")
		return
	}
	defer session.Close()

	c := session.DB(mongoConfig.Database).C(mongoConfig.Table)
	err = c.Insert(&agentInfo)
	if err != nil {
		log.WithFields(log.Fields{"error: ": err}).Info("Fail inserting into mongoDB")
	}
}

//GetFoodsInFridge shows all food in a fridge
func GetFoodsInFridge(containersID []string) []FoodInfo {
	session, err := getSession()
	if err != nil {
		log.WithFields(log.Fields{"error: ": err}).Info("Cannot create session to mongoDB")
		return []FoodInfo{}
	}

	defer session.Close()
	c := session.DB(mongoConfig.Database).C(mongoConfig.Table)

	var foods []FoodInfo
	for _, value := range containersID {
		var agent FoodAgent

		if err := c.Find(bson.M{"token": value}).One(&agent); err == nil {
			foods = append(foods, agent.FoodInfo)
		}
	}
	return foods
}
