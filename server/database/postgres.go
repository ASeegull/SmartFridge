package database

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ASeegull/SmartFridge/server/config"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/satori/go.uuid"
)

const (
	avgNumrOfIngInRecepie  = 7
	prognosedNumOfRecepies = 100
	prognosedNumOfProducts = 100
	avgNumOfAgentsOfUser   = 10
)

var dbinfo string
var db *gorm.DB

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
func RegisterNewUser(login string, passHash string) error {
	var err error
	user := User{}
	err = db.Where("login = ?", login).Find(&user).Error
	switch {
	case err != nil:
		return err
	case user.ID != "":
		return errors.New("login is already taken")
	}
	user.ID = uuid.NewV4().String()
	user.Login = login
	user.Password = passHash
	rows := db.Create(&user).RowsAffected
	if rows != int64(1) {
		return errors.New("user not added")
	}
	return nil
}

//ClientLogin checks login and pass for client
func ClientLogin(login string, pass string) error {
	var err error
	user := User{}
	err = db.Where("login = ?", login).Find(&user).Error
	switch {
	case err != nil:
		return err
	case user.ID == "":
		return errors.New("login not found")
	case strings.TrimRight(pass, "\n") != user.Password:
		return errors.New("wrong password")
	}
	return nil
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

//RegisterNewAgent adds a new agent to user, returns nil if adding was successful
func RegisterNewAgent(idUser string, idAgent string) error {
	agent := Agent{ID: idAgent, UserID: idUser}
	rows := db.Create(&agent).RowsAffected
	if !(rows == 1) {
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

//GetDefaultExplorationDate function returns expiration date a product as time.Time object
func GetDefaultExplorationDate(productName string) (time.Time, error) {
	var err error
	product := Product{}
	err = db.Where("name LIKE ?", strings.ToLower(productName)).First(&product).Error
	switch {
	case err != nil:
		return time.Time{}, err
	case product.Name == "":
		return time.Time{}, errors.New("product not found")
	}
	return time.Now().Add(time.Hour * 24 * time.Duration(product.ShelfLife)), nil
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

//GetAllProductsNames returns a slice, containing IDs of all products
func GetAllProductsName() ([]string, error) {
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
