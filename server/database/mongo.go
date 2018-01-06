package database

import (
	"fmt"
	"sync"
	"time"

	"github.com/ASeegull/SmartFridge/server/config"
	"github.com/davecheney/errors"
	"github.com/sirupsen/logrus"
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
	//dialInfo, err := mgo.ParseURL(mongoConfig.URI)
	//if err != nil {
	//	return nil, err
	//}
	//tlsConfig := &tls.Config{InsecureSkipVerify: true}
	//
	//dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
	//	return tls.Dial("tcp", addr.String(), tlsConfig)
	//}
	return mgo.Dial(mongoConfig.URI)
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
			if err := c.Find(bson.M{"agentid": val}).Sort("-_id").One(&agent); err != nil {
				mutex.Lock()
				notFound = append(notFound, *val)
				mutex.Unlock()
				return
			}
			mutex.Lock()
			foods = append(foods, FoodInfo{AgentID: agent.AgentID, Product: agent.Product, Weight: agent.Weight, Expires: agent.StateExpires, Condition: checkConditions(agent.StateExpires)})
			mutex.Unlock()
		}(&containersID[i])

	}
	wg.Wait()
	URLs, err := GetImagesByNames(foods)
	if err != nil {
		logrus.Error(err)
		return foods, nil
	}

	for index := range foods {
		var productURL string
		if url, ok := URLs[foods[index].Product]; ok {
			productURL = url
		} else {
			productURL = "http://www.reallanguagerightaway.com/content/images/thumbs/default-image_450.png"
		}
		foods[index].URL = productURL
	}
	return foods, nil
}

func checkConditions(state string) string {
	productTime, err := time.Parse("2006-01-02", state)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	days := int(time.Since(productTime).Hours() / (-24))
	switch {
	case days < 0:
		return "expired"
	case days < 2:
		return "warn"
	default:
		return "ok"
	}
}
