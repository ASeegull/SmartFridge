package database

//FoodInfo is example struct for agent
type FoodInfo struct {
	Product   string `bson:"product" json:"product"`
	Weight    int32  `bson:"weight" json:"weight"`
	Expires   string `bson:"stateexpires" json:"stateExpires"`
	Condition string `bson:"condition" json:"condition"`
	URL       string `json:"imageURL"`
}

//FoodAgent is example struct for agent
type FoodAgent struct {
	Token        string `bson:"token"`
	UserID       string `bson:"userid"`
	AgentID      string `bson:"agentid"`
	Product      string `bson:"productid"`
	Weight       int32  `bson:"weight"`
	StateExpires string `bson:"stateexpires"`
}

//MUnit represents units of measure used for products
type MUnit struct {
	ID   string
	Unit string
}

//Agent represents an agent entity
type Agent struct {
	ID          string
	AgentSerial string
}

//UserAgent is a row in joining table between user and agent
type UserAgent struct {
	UserID  string
	AgentID string
}

//Ingridient represents an ingredient in a recepie
type Ingridient struct {
	ProductID string
	RecipeID  string
	Amount    int
}

//Product represents an product
type Product struct {
	ID        string `json:"ID"`
	Name      string `json:"name"`
	ShelfLife int    `json:"shelfLife"`
	Units     string `json:"units"`
	Image     string `json:"imageURL"`
}

//Recepie represents a recepie
type Recepie struct {
	ID              string   `json:"-"`
	RecName         string   `json:"title"`
	Description     string   `json:"description"`
	CoockingTimeMin int      `json:"coockingTimeMin"`
	Complexity      string   `json:"complexity"`
	Ingred          []string `gorm:"-" json:"ingredients"`
}

//User represents a user
type User struct {
	ID       string
	Login    string
	Password string
	Role     string
}

//UserRole represents the authorities of an user
type UserRole struct {
	ID   string
	Role string
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

//LogIn logged in
func (log *Login) LogIn() error {
	return ClientLogin(log.UserName, log.Pass)
}

//Register registers new user
func (log *Login) Register() error {
	RegisterNewUser(log.UserName, log.Pass)
	return nil
}
