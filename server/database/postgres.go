package database

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
	"errors"
	"github.com/jinzhu/gorm"
//importing Go Postgres driver for database
	_ "github.com/lib/pq"
	"log"
)
//MUnit represents units of measure used for products
type MUnit struct {
	ID   int
	Unit string
}
//Agent represents an agent entity
type Agent struct {
	ID     int
	UserID int
}
//Ingridient represents an ingredient in a recepie
type Ingridient struct {
	ProductID int
	RecipeID  int
	Amount    int
}
//Product represents an product
type Product struct {
	ID        int
	Name      string
	ShelfLife int
	Units     int
}
//Recepie represents a recepie
type Recepie struct {
	ID              int			`json:"id"`
	Name            string		`json:"name"`
	Description     string		`json:"description"`
	CoockingTimeMin int			`json:"coockingTimeMin"`
	Complexity      string		`json:"complexity"`
	Ingred          []string	`json:"ingred"`
}
//User represents a user
type User struct {
	ID       int
	Login    string
	Password string
}

const (
	//postgres connection credentials
	dbUser     = "postgres"
	dbPassword = ""
	dbName     = "postgres"
)

//RegisterNewUser adds a new user, returns error if adding was not successful
func RegisterNewUser(login string, passHash string) error {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		dbUser, dbPassword, dbName)
	db, err := gorm.Open("postgres", dbinfo)
	if err != nil {
		return errors.New("db connection error")
	}
	defer db.Close()
	if err != nil {
		log.Println(err)
	}
	user := User{}
	db.Last(&user)
	user.ID++
	user.Login = login
	user.Password = passHash
	db.Create(&user)
	rows := db.Create(&user).RowsAffected
	if !(rows == 1) {return errors.New("not added")} else {
		return nil
	}
}

//ClientLogin checks login and pass for client
func ClientLogin(login string, pass string) error {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		dbUser, dbPassword, dbName)
	db, err := gorm.Open("postgres", dbinfo)
	if err != nil {
		return errors.New("db connection error")
	}
	defer db.Close()
	if err != nil {
		log.Println(err)
	}
	user := User{}
	db.Where("name = ?", login).Find(&user)
	if user.ID == 0 {
		return errors.New("login not found")
	} else if !(strings.TrimRight(pass, "\n") == user.Password) {
		return errors.New("wrong password")
	} else {
		return nil
	}
}

//CheckAgent checks agent registration, if agent is associated with a user returns true, nil or false
func CheckAgent(idUser int, idAgent int) bool {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		dbUser, dbPassword, dbName)
	db, err := gorm.Open("postgres", dbinfo)
	if err != nil {
		panic(errors.New("db connection error"))
	}
	defer db.Close()
	if err != nil {
		log.Println(err)
	}
	agent := Agent{}
	db.Where("agents.id = ? AND agents.user_id = ?", idAgent, idUser).Find(&agent)
	return agent.ID != 0
}

//RegisterNewAgent adds a new agent to user, returns nil if adding was successful
func RegisterNewAgent(idAgent int, idUser int) error {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		dbUser, dbPassword, dbName)
	db, err := gorm.Open("postgres", dbinfo)
	if err != nil {
		return errors.New("db connection error")
	}
	defer db.Close()
	if err != nil {
		log.Println(err)
	}
	agent := Agent{ID: idAgent, UserID: idUser}
	rows := db.Create(&agent).RowsAffected
	if !(rows == 1) {return errors.New("can not register an agent")} else{
		return nil
	}
}


//GetAllAgentsIDForClient returns all agent for clientID as a sice of string
func GetAllAgentsIDForClient(userID string) ([]string, error) {
		dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
			dbUser, dbPassword, dbName)
		db, err := gorm.Open("postgres", dbinfo)
	if err != nil {
		return nil, errors.New("db connection error")
	}
	defer db.Close()
	if err != nil {
		log.Println(err)
	}
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		return nil, errors.New("uncorrect id value")
	}
		var agentID int
		agentIds := make([]string, 0)
		rows, err := db.Table("agents").Select("agents.id").Where("agents.user_id=?", userIDInt).Rows()
			if err != nil {
		return nil, errors.New("agents id no found")
	}
		for rows.Next() {
		err = rows.Scan(&agentID)
			if err != nil {
				return nil, errors.New("agents id no found")
			}
		agentIds = append(agentIds, strconv.Itoa(agentID)+" ")
	}
		return agentIds, nil
}

//GetDefaultExplorationDate function returns expiration date a product as time.Time object
func GetDefaultExplorationDate(productName string) (time.Time, error) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		dbUser, dbPassword, dbName)
	db, err := gorm.Open("postgres", dbinfo)
	if err != nil {
		return time.Time{}, errors.New("db connection error")
	}
	defer db.Close()
	if err != nil {
		log.Println(err)
	}
	product := Product{}
	db.Where("name LIKE ?", productName).Find(&product)
	return time.Now().Add(time.Hour * 24 * time.Duration(product.ShelfLife)), nil
}
//AllRecipes functions returns all Recipes with ingridients as a JSON
func AllRecipes() ([]byte, error) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		dbUser, dbPassword, dbName)
	db, err := gorm.Open("postgres", dbinfo)
	if err != nil {
		return nil, errors.New("db connection error")
	}
	defer db.Close()
	if err != nil {
		log.Println(err)
	}
	recipes := []Recepie{}

	db.Find(&recipes)
	fmt.Print(recipes)
	fmt.Println()

	var name, unit string
	var amount int
	for k, v := range recipes {
		rows, err := db.Table("recepies").Select("ingridients.amount, m_units.unit, products.name").
			Joins("LEFT JOIN ingridients on ingridients.recipe_id = recepies.id").
			Joins("JOIN products on ingridients.product_id = products.id").
			Joins("JOIN m_units on m_units.id = products.units").
			Where("recepies.id=?", v.ID).Rows()

		if err != nil {
			return nil, errors.New("recepies search error")
		}

		for rows.Next() {
			err = rows.Scan(&amount, &unit, &name)
			if err != nil {
				return nil, errors.New("ingredients search error")
			}
			recipes[k].Ingred = append(recipes[k].Ingred, strconv.Itoa(amount)+unit+name+";")
		}
	}
	allRecipies, err := json.Marshal(recipes)
	if err != nil {
		return nil, errors.New("marshaling error")
	}
	return allRecipies, nil
}

