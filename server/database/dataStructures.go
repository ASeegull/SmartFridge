package database

//FoodInfo is example struct for agent
type FoodInfo struct {
	Name   string `bson:"name"`
	Veight int    `bson:"veight"`
	Date   string `bson:"date"`
	Value  int    `bson:"value"`
}

//FoodAgent is example struct for agent
type FoodAgent struct {
	Token       string   `bson:"token"`
	ContainerID string   `bson:"containerid"`
	GoodWork    bool     `bson:"goodwork"`
	FoodInfo    FoodInfo `bson:"foodinfo"`
}
