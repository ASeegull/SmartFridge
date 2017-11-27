package database

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
	RecName         string   `json:"recName"`
	Description     string   `json:"description"`
	CoockingTimeMin int      `json:"coockingTimeMin"`
	Complexity      string   `json:"complexity"`
	Ingred          []string `gorm:"-" json:"ingred"`
}

//User represents a user
type User struct {
	ID       string
	Login    string
	Password string
	Role     string
}
