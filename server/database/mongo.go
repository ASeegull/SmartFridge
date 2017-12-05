package database

import (
	"crypto/tls"
	"net"

	"github.com/ASeegull/SmartFridge/server/config"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	pb "github.com/ASeegull/SmartFridge/protoStruct"
	log "github.com/sirupsen/logrus"
)

var session *mgo.Session
var mongoConfig config.MongoConfig

//InitiateMongoDB sets config for mongoDB
func InitiateMongoDB(cfg config.MongoConfig) error {
	mongoConfig = cfg
	return createSession()
}

func createSession() error {
	dialInfo, err := mgo.ParseURL(mongoConfig.URI)
	if err != nil {
		log.Debug(err)
		return err
	}
	tlsConfig := &tls.Config{}

	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		return tls.Dial("tcp", addr.String(), tlsConfig)
	}

	session, err = mgo.DialWithInfo(dialInfo)
	session.SetPoolLimit(100)
	if err != nil {
		log.Debug(err)
		return err
	}
	return nil
}

//SaveState saves state from agent
func SaveState(agentInfo *pb.Agentstate) error {
	c := session.DB(mongoConfig.Database).C(mongoConfig.Table)
	err := c.Insert(&agentInfo)
	if err != nil {
		log.Debug(err)
	}
	return err
}

//GetFoodsInFridge shows all food in a fridge
func GetFoodsInFridge(containersID []string) ([]FoodInfo, error) {

	c := session.DB(mongoConfig.Database).C(mongoConfig.Table)

	foods := make([]FoodInfo, 0, len(containersID))
	for _, value := range containersID {
		var agent FoodAgent

		if err := c.Find(bson.M{"agentid": value}).One(&agent); err != nil {
			log.Println(err)
			return nil, err
		}

		foods = append(foods, FoodInfo{agent.Product, agent.Weight})
	}
	return foods, nil
}
