package database

import (
	"crypto/tls"
	"net"
	"sync"

	"github.com/ASeegull/SmartFridge/server/config"
	"github.com/davecheney/errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	pb "github.com/ASeegull/SmartFridge/protoStruct"
)

var session *mgo.Session
var mongoConfig config.MongoConfig

//InitiateMongoDB sets config for mongoDB
func InitiateMongoDB(cfg config.MongoConfig) error {
	mongoConfig = cfg

	if session != nil {
		return errors.New("session already exist")
	}

	var err error
	session, err = createSession()
	if err == nil {
		session.SetPoolLimit(mongoConfig.ConnectionsPool)
	}

	return err
}

func createSession() (*mgo.Session, error) {
	dialInfo, err := mgo.ParseURL(mongoConfig.URI)
	if err != nil {
		return nil, err
	}
	tlsConfig := &tls.Config{InsecureSkipVerify: true}

	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		return tls.Dial("tcp", addr.String(), tlsConfig)
	}

	return mgo.DialWithInfo(dialInfo)
}

//SaveState saves state from agent
func SaveState(agentInfo *pb.Agentstate) error {
	return session.DB(mongoConfig.Database).C(mongoConfig.Table).Insert(&agentInfo)
}

//GetFoodsInFridge shows all food in a fridge
func GetFoodsInFridge(containersID []string) ([]FoodInfo, error) {

	c := session.DB(mongoConfig.Database).C(mongoConfig.Table)
	cLen := len(containersID)
	foods := make([]FoodInfo, 0, cLen)
	var wg sync.WaitGroup
	var mutex = &sync.Mutex{}
	wg.Add(cLen)
	notFound := make([]string, 0, cLen)

	for i := range containersID {

		go func(val *string) {
			var agent FoodAgent
			defer wg.Done()
			if err := c.Find(bson.M{"agentid": val}).One(&agent); err != nil {
				mutex.Lock()
				notFound = append(notFound, *val)
				mutex.Unlock()
				return
			}
			mutex.Lock()
			foods = append(foods, FoodInfo{agent.Product, agent.Weight, agent.StateExpires, ""})
			mutex.Unlock()
		}(&containersID[i])

	}
	wg.Wait()
	return foods, nil
}
