package database

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
	"errors"
	"github.com/jinzhu/gorm"

	_ "github.com/lib/pq"
)

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
			recipes[k].Ingred = append(recipes[k].Ingred, strconv.Itoa(amount), unit, name)
		}
	}
	allRecipies, err := json.Marshal(recipes)
	if err != nil {
		return nil, errors.New("marshaling error")
	}
	return allRecipies, nil
}
//Recipes(foodInfoSlice []FoodInfo) takes the slice of FoodInfo strucktures, representing all available products in all agents and return all recepies, which can be offered as a JSON
func Recipes(foodInfoSlice []FoodInfo) ([]byte, error) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		dbUser, dbPassword, dbName)
	db, err := gorm.Open("postgres", dbinfo)
	if err != nil {
		return nil, errors.New("db connection error")
	}
	defer db.Close()
	productNameSlice := make([]string, 0)
	productMap := make(map[string]int)
	for _, v := range foodInfoSlice {
		productNameSlice = append(productNameSlice, strings.ToLower(v.Name))
		productMap[strings.ToLower(v.Name)] = v.Weight
	}
	recipes := []Recepie{}
	db.Table("recepies").
		Joins("FULL JOIN ingridients on ingridients.recipe_id = recepies.id").
		Joins("JOIN products on ingridients.product_id = products.id").
		Where("products.name IN (?)", productNameSlice).
		Having("count(products.id) <= ?", len(productNameSlice)).
		Group("recepies.id").
		Find(&recipes)
	var name, unit string
	var amount int
OUTER:
	for k := len(recipes) - 1; k >= 0; k-- {
		rows, _ := db.Table("recepies").Select("ingridients.amount, m_units.unit, products.name").
			Joins("LEFT JOIN ingridients on ingridients.recipe_id = recepies.id").
			Joins("JOIN products on ingridients.product_id = products.id").
			Joins("JOIN m_units on m_units.id = products.units").
			Where("recepies.id=?", recipes[k].ID).Rows()
		if err != nil {
			return nil, errors.New("recepies search error")
		}
		for rows.Next() {
			rows.Scan(&amount, &unit, &name)
			if contains(productNameSlice, name) && amount <= productMap[name] {
				recipes[k].Ingred = append(recipes[k].Ingred, strconv.Itoa(amount), unit, name)
			} else {
				recipes = append(recipes[:k], recipes[k+1:]...)
				continue OUTER
			}
		}
	}
	allRecipies, err := json.Marshal(recipes)
	if err != nil {
		return nil, errors.New("marshaling error")
	}
	return allRecipies, nil
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

//CreteTables() can be used to create tables, needed for the project
func CreteTables() error {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		dbUser, dbPassword, dbName)
	db, err := gorm.Open("postgres", dbinfo)
	if err != nil {
		return  errors.New("db connection error")
	}
	defer db.Close()
	db.Exec("CREATE TABLE IF NOT EXISTS m_units (id INT primary KEY, unit VARCHAR (15));")

	db.Exec("CREATE TABLE IF NOT EXISTS products (id INT primary KEY, name VARCHAR (15), shelf_life INT, units INT REFERENCES m_units(id));")

	db.Exec("CREATE TABLE IF NOT EXISTS recepies (id INT primary KEY, name VARCHAR (15), description TEXT, coocking_time_min INT, complexity VARCHAR (15));")

	db.Exec("CREATE TABLE IF NOT EXISTS ingridients (product_id INT REFERENCES products(id),	recipe_id INT REFERENCES recepies(id),	amount INT);")

	db.Exec("CREATE TABLE IF NOT EXISTS users (id INT primary KEY, login VARCHAR (15), password VARCHAR (15));")

	db.Exec("CREATE TABLE IF NOT EXISTS agents (id INT primary KEY, user_id INT REFERENCES users(id));")
	return nil
}

//FillTables() puts some information to tables to work with
func FillTables() error {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		dbUser, dbPassword, dbName)
	db, err := gorm.Open("postgres", dbinfo)
	if err != nil {
		return  errors.New("db connection error")
	}
	defer db.Close()

	mu := MUnit{ID: 1, Unit: "gramm"}
	db.Create(&mu)
	mu = MUnit{ID: 2, Unit: "ml"}
	db.Create(&mu)
	mu = MUnit{ID: 3, Unit: "pieces"}
	db.Create(&mu)

	p := Product{ID: 1, Name: "tomato", ShelfLife: 15, Units: 1}
	db.Create(&p)
	p = Product{ID: 2, Name: "potato", ShelfLife: 30, Units: 1}
	db.Create(&p)
	p = Product{ID: 3, Name: "milk", ShelfLife: 2, Units: 2}
	db.Create(&p)
	p = Product{ID: 4, Name: "onion", ShelfLife: 30, Units: 1}
	db.Create(&p)
	p = Product{ID: 5, Name: "cucumber", ShelfLife: 7, Units: 1}
	db.Create(&p)
	p = Product{ID: 6, Name: "sausage", ShelfLife: 7, Units: 1}
	db.Create(&p)
	p = Product{ID: 7, Name: "butter", ShelfLife: 15, Units: 1}
	db.Create(&p)
	p = Product{ID: 8, Name: "egg", ShelfLife: 15, Units: 3}
	db.Create(&p)
	p = Product{ID: 9, Name: "meat", ShelfLife: 5, Units: 1}
	db.Create(&p)
	p = Product{ID: 10, Name: "chicken", ShelfLife: 7, Units: 1}
	db.Create(&p)
	p = Product{ID: 11, Name: "bread", ShelfLife: 7, Units: 1}
	db.Create(&p)

	r := Recepie{ID: 1, Name: "Salat", Description: "...", CoockingTimeMin: 15, Complexity: "easy"}
	db.Create(&r)
	r = Recepie{ID: 2, Name: "Sandwich", Description: "...", CoockingTimeMin: 5, Complexity: "easy"}
	db.Create(&r)
	r = Recepie{ID: 3, Name: "Soup", Description: "...", CoockingTimeMin: 35, Complexity: "easy"}
	db.Create(&r)

	i := Ingridient{RecipeID: 1, ProductID: 1, Amount: 300}
	db.Create(&i)
	i = Ingridient{RecipeID: 1, ProductID: 2, Amount: 600}
	db.Create(&i)
	i = Ingridient{RecipeID: 1, ProductID: 4, Amount: 300}
	db.Create(&i)
	i = Ingridient{RecipeID: 1, ProductID: 5, Amount: 100}
	db.Create(&i)
	i = Ingridient{RecipeID: 2, ProductID: 11, Amount: 50}
	db.Create(&i)
	i = Ingridient{RecipeID: 2, ProductID: 7, Amount: 50}
	db.Create(&i)
	i = Ingridient{RecipeID: 2, ProductID: 6, Amount: 50}
	db.Create(&i)
	i = Ingridient{RecipeID: 2, ProductID: 1, Amount: 20}
	db.Create(&i)

	u := User{ID: 1, Login: "login", Password: "password"}
	db.Create(&u)

	a := Agent{ID: 1, UserID: 1}
	db.Create(&a)
	a = Agent{ID: 2, UserID: 1}
	db.Create(&a)
	return nil
}


