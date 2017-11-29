package database

import (
	"encoding/json"
)

//FoodInfo is example struct for agent
type FoodInfo struct {
	ProductID string `bson:"productID"`
	Weight    int32  `bson:"weight"`
}

//FoodAgent is example struct for agent
type FoodAgent struct {
	Token        string `bson:"token"`
	UserID       string `bson:"userID"`
	AgentID      string `bson:"agentID"`
	ProductID    string `bson:"productID"`
	Weight       int32  `bson:"weight"`
	StateExpires int32  `bson:"stateExpires"`
}

//MUnit represents units of measure used for products
type MUnit struct {
	ID   int
	Unit string
}

//Agent represents an agent entity
type Agent struct {
	ID     string
	UserID string
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
	ID              int      `json:"id"`
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	CoockingTimeMin int      `json:"coockingTimeMin"`
	Complexity      string   `json:"complexity"`
	Ingred          []string `gorm:"-" json:"ingred"`
}

//RecipesStr contains all recipes
type RecipesStr struct {
	Recipes []Recepie `json:"recipes"`
}

//User represents a user
type User struct {
	ID       string
	Login    string
	Password string
}

//Login contains username and password
type Login struct {
	UserName string `json:"login"`
	Pass     string `json:"pass"`
}

//UserID contains user ID
type UserID struct {
	ID string `json:"id"`
}

//GetAllAgentIDs returns all agents ID for this userID
func (lg *UserID) GetAllAgentIDs() ([]string, error) {
	return GetAllAgentsIDForClient(lg.ID)
}

//GetFoodsInFridge returns all food from fridge for this userID
func (lg *UserID) GetFoodsInFridge() ([]FoodInfo, error) {
	IDs, err := lg.GetAllAgentIDs()
	if err != nil {
		return nil, err
	}

	return GetFoodsInFridge(IDs)
}

//GetAllRecipes returns all recipes
func (rec *RecipesStr) GetAllRecipes() error {
	var err error
	if rec.Recipes, err = AllRecipes(); err != nil {
		return err
	}
	return nil
}

//LogIn logged in
func (log *Login) LogIn() error {
	return ClientLogin(log.UserName, log.Pass)
}

//Register registers new user
func (log *Login) Register() error {
	return RegisterNewUser(log.UserName, log.Pass)
}

//Unmarshalling unmarchals to this struct
func (log Login) Unmarshalling(data []byte) error {
	return json.Unmarshal(data, &log)
}
