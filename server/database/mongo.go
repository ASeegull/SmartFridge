package database

import (
	"crypto/tls"
	"net"

	"github.com/ASeegull/SmartFridge/server/config"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	log "github.com/sirupsen/logrus"
)

var mongoConfig = config.GetMongoConfig()
var session = createSession()

func createSession() *mgo.Session {
	dialInfo, _ := mgo.ParseURL(mongoConfig.URI)
	tlsConfig := &tls.Config{}

	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}

	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Error(err)
	}
	return session
}

//SaveState saves state from agent
func SaveState(agentInfo *FoodAgent) {
	c := session.DB(mongoConfig.Database).C(mongoConfig.Table)
	err := c.Insert(&agentInfo)
	if err != nil {
		log.Error(err)
	}
}

//GetFoodsInFridge shows all food in a fridge
func GetFoodsInFridge(containersID []string) []FoodInfo {
	c := session.DB(mongoConfig.Database).C(mongoConfig.Table)

	var foods []FoodInfo
	for _, value := range containersID {
		var agent FoodAgent

		if err := c.Find(bson.M{"containerID": value}).One(&agent); err == nil {
			foods = append(foods, agent.FoodInfo)
		} else {
			log.Error(err)
		}
	}
	return foods
}
