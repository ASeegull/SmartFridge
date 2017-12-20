package database

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"

	pb "github.com/ASeegull/SmartFridge/protoStruct"
	"github.com/ASeegull/SmartFridge/server/config"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

const (
	avgNumrOfIngInRecepie  = 7
	prognosedNumOfRecepies = 100
	prognosedNumOfProducts = 100
	avgNumOfAgentsOfUser   = 10
)

const (
	PublicToken = "verysecrettoken"
	privateKey  = "muchmoresecrettoken"
)

var dbinfo string
var db *gorm.DB

//CheckToken checks tokens conformity
func CheckToken(agent *pb.Agentstate) bool {
	return agent.Token == getHash(privateKey+agent.AgentID+agent.UserID)
}

func getHash(str string) string {
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}

// SetToken assigns token to the given agent
func SetToken(agent *pb.Agentstate) {
	agent.Token = getHash(privateKey + agent.AgentID + agent.UserID)
}

//InitPostgersDB initiates connection to postgres gatabase
func InitPostgersDB(cfg config.PostgresConfigStr) error {
	var err error
	if db != nil {
		return nil
	}
	dbinfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Dbhost, cfg.Dbport, cfg.DbUser, cfg.DbPassword, cfg.DbName)
	db, err = gorm.Open("postgres", dbinfo)
	if err != nil {
		return err
	}
	db.DB().SetMaxOpenConns(cfg.MaxOpenedConnectionsToDb)
	db.DB().SetMaxIdleConns(cfg.MaxIdleConnectionsToDb)
	db.DB().SetConnMaxLifetime(time.Minute * time.Duration(cfg.MbConnMaxLifetimeMinutes))
	return nil
}

//RegisterNewUser adds a new user, returns error if adding was not successful
func RegisterNewUser(login string, passHash string) (string, error) {
	user := User{}
	err := db.Where("login = ?", login).Find(&user).Error
	if err == nil {
		return "", errors.New("login is already taken")
	}
	user.ID = uuid.NewV4().String()
	user.Login = login
	user.Password = passHash
	rows := db.Create(&user).RowsAffected
	if rows != int64(1) {
		return "", errors.New("user not added")
	}
	return user.ID, nil
}

//ClientLogin checks login and pass for client
func ClientLogin(login string, pass string) error {
	user := User{}
	err := db.Where("login = ?", login).Find(&user).Error
	if err != nil {
		return errors.Wrap(err, "login not found")
	}
	if strings.TrimRight(pass, "\n") != user.Password {
		return errors.New("wrong password")
	}
	return nil
}

//GetUserID checks login and pass for client
func GetUserID(login string) (string, error) {
	user := User{}
	err := db.Where("login = ?", login).Find(&user).Error
	if err != nil {
		return "", errors.Wrapf(err, "cannot find user with login %s", login)
	}
	return user.ID, nil
}

//CheckAgent checks agent registration, if agent is associated with a user returns true as first returning value
func CheckAgent(idUser string, idAgent string) (bool, error) {
	var err error
	agent := Agent{}
	err = db.Where("agents.id = ? AND agents.user_id = ?", idAgent, idUser).Find(&agent).Error
	if err != nil {
		return false, nil
	}
	return agent.ID != "", nil
}

// GetAgentOwner returns user ID registered for particular agent
func GetAgentOwner(agentID string) (string, error) {
	agent := Agent{}
	err := db.Where("agents.id = ?", agentID).Find(&agent).Error
	// If agent has no user id, thats ok, user will sign up soon
	if err == gorm.ErrRecordNotFound {
		return "", nil
	}

	if err != nil {
		return "", err
	}
	return agent.UserID, nil
}

//RegisterNewAgent adds a new agent to user, returns nil if adding was successful
func RegisterNewAgent(idAgent string) error {
	agent := Agent{ID: idAgent}
	rows := db.Create(&agent).RowsAffected
	if rows != 1 {
		return errors.New("can not register an agent")
	}
	return nil
}

//GetAllAgentsIDForClient returns all agent for clientID as a slice of string
func GetAllAgentsIDForClient(userID string) ([]string, error) {
	var err error
	var agentID string
	agentIds := make([]string, 0, avgNumOfAgentsOfUser)
	rows, err := db.Table("agents").Select("agents.id").Where("agents.user_id=?", userID).Rows()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&agentID)
		if err != nil {
			return nil, err
		}
		agentIds = append(agentIds, agentID)
	}
	return agentIds, nil
}

// GetDefaultExpirationDate queries database and returns
// avarage shelf time of a product as int in hours
func GetDefaultExpirationDate(productName string) (int, error) {
	product := Product{}
	err := db.Where("name LIKE ?", strings.ToLower(productName)).First(&product).Error
	if err != nil {
		return 0, errors.Wrapf(err, "for the product %s", productName)
	}
	if product.ShelfLife == 0 {
		return 0, errors.Errorf("there is no shelflife for the product %s", productName)
	}
	return product.ShelfLife, nil
}

// SetExpirationDate sets default expiration date if none is provided by user
func SetExpirationDate(shelftime int) string {
	return time.Now().Add(time.Duration(shelftime) * time.Duration(24) * time.Hour).Format(time.ANSIC)
}

// CheckCondition sets Condition of FoodInfo
// depending on expiration date
func (product *FoodInfo) CheckCondition() error {
	expdate, err := time.Parse(time.ANSIC, product.Expires)
	if err != nil {
		return err
	}
	delta := time.Now().Sub(expdate).Hours()
	switch {
	case delta < 48:
		product.Condition = "warn"
	case delta < 0:
		product.Condition = "expired"
	default:
		product.Condition = "ok"
	}
	return nil
}

//AllRecipes functions returns all Recipes with ingridients
func AllRecipes() ([]Recepie, error) {
	var err error
	recipes := make([]Recepie, 0, prognosedNumOfRecepies)
	var recName, description, complexity, name, unit string
	var id, amount, coockingTimeMin int
	rows, err := db.Table("recepies").Select("recepies.id, recepies.rec_name, recepies.description, recepies.coocking_time_min, recepies.complexity, ingridients.amount, m_units.unit, products.name").
		Joins("LEFT JOIN ingridients on ingridients.recipe_id = recepies.id").
		Joins("JOIN products on ingridients.product_id = products.id").
		Joins("JOIN m_units on m_units.id = products.units").
		Rows()
	if err != nil {
		return nil, err
	}
	var newRec string
	k := 0
	for rows.Next() {
		err = rows.Scan(&id, &recName, &description, &coockingTimeMin, &complexity, &amount, &unit, &name)
		if err != nil {
			return nil, err
		}
		if recName != newRec {
			ing := make([]string, 0, avgNumrOfIngInRecepie)
			recipes = append(recipes, Recepie{ID: id, RecName: recName, Complexity: complexity, CoockingTimeMin: coockingTimeMin, Description: description, Ingred: append(ing, strconv.Itoa(amount)+" "+name+" "+unit)})
			newRec = recName
			k++
		} else {
			recipes[k-1].Ingred = append(recipes[k-1].Ingred, strconv.Itoa(amount)+" "+unit+" "+name)
		}
	}
	return recipes, nil
}

// GetAllProductsNames returns a slice, containing IDs of all products
func GetAllProductsNames() ([]string, error) {
	var err error
	var productName string
	productNames := make([]string, 0, prognosedNumOfProducts)
	rows, err := db.Table("products").Select("products.name").Rows()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&productName)
		if err != nil {
			return nil, err
		}
		productNames = append(productNames, productName)
	}
	return productNames, nil
}

//Recipes takes the slice of FoodInfo strucktures, representing all available products in all agents and return all recepies, which can be offered
func Recipes(foodInfoSlice []FoodInfo) ([]Recepie, error) {
	var err error
	productNameSlice := make([]string, 0, avgNumrOfIngInRecepie)
	productMap := make(map[string]int)
	for _, v := range foodInfoSlice {
		productNameSlice = append(productNameSlice, strings.ToLower(v.Product))
		productMap[strings.ToLower(v.Product)] = int(v.Weight)
	}

	recipes := make([]Recepie, 0, prognosedNumOfRecepies)
	err = db.Table("recepies").
		Joins("FULL JOIN ingridients on ingridients.recipe_id = recepies.id").
		Joins("JOIN products on ingridients.product_id = products.id").
		Where("products.name IN (?)", productNameSlice).
		Having("count(products.id) <= ?", len(productNameSlice)).
		Group("recepies.id").
		Find(&recipes).Error
	if err != nil {
		return nil, err
	}

	var name, unit string
	var amount int
	copyRec := make([]Recepie, 0, len(recipes))

OUTER:
	for _, recipe := range recipes {
		rows, err := db.Table("recepies").Select("ingridients.amount, m_units.unit, products.name").
			Joins("LEFT JOIN ingridients on ingridients.recipe_id = recepies.id").
			Joins("JOIN products on ingridients.product_id = products.id").
			Joins("JOIN m_units on m_units.id = products.units").
			Where("recepies.id=?", recipe.ID).Rows()
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			err := rows.Scan(&amount, &unit, &name)
			if err != nil {
				return nil, err
			}
			if contains(productNameSlice, name) && amount <= productMap[name] {
				recipe.Ingred = append(recipe.Ingred, strconv.Itoa(amount), unit, name)
			} else {
				continue OUTER
			}
		}
		copyRec = append(copyRec, recipe)
	}
	return copyRec, nil
}

//contains shows if a slice contains given value
func contains(slice []string, v string) bool {
	for _, a := range slice {
		if a == v {
			return true
		}
	}
	return false
}
