package database

//FoodInfo is example struct from agent
type FoodInfo struct {
	Name   string `bson:"name"`
	Veight int    `bson:"veight"`
	Date   string `bson:"date"`
	Value  int    `bson:"value"`
}

//FoodInfo is example struct from agent
type FoodAgent struct {
	Token       string   `bson:"token"`
	ContainerID string   `bson:"containerid"`
	GoodWork    bool     `bson:"goodwork"`
	FoodInfo    FoodInfo `bson:"foodinfo"`
}
