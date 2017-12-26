package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/ASeegull/SmartFridge/server/database"
	"github.com/davecheney/errors"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"

	pb "github.com/ASeegull/SmartFridge/protoStruct"
	"github.com/gorilla/mux"
)

const (
	defaultHeartBeat = 3
)

func sendResponse(w http.ResponseWriter, status int, err error) {
	if err != nil {
		log.Error(err)
		http.Error(w, fmt.Sprintf("%v", err), status)
		return
	}

	w.WriteHeader(status)
}

func checkSession(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		isNew, err := isNewSession(w, r)
		if err != nil {
			sendResponse(w, http.StatusInternalServerError, err)
			return
		}

		if isNew {
			sendResponse(w, http.StatusUnauthorized, err)
			return
		}

		h.ServeHTTP(w, r)
	})
}

func agentAuthentication(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		sendResponse(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	pbRequest := pb.Request{}
	if err = pbRequest.UnmarshalToStruct(body); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	//get product is mock. user has to set it in server page
	productName, err := getProduct()
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	//adminID is mock. Here must be call to postgress db (table user - agentsID)
	userID, err := getUserID(r)
	if err := database.RegisterNewAgent(userID, pbRequest.AgentID); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	resp := pb.Setup{}
	resp.SetParameters(pbRequest.AgentID, userID, *productName, defaultHeartBeat)

	data, err := resp.MarshalStruct()
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	_, err = w.Write(data)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
	}
}

//getProduct is mock. Method returns random product from postgres DB (table Products)
func getProduct() (*string, error) {

	IDs, err := database.GetAllProductsNames()
	if err != nil {
		return nil, err
	}
	rand.Seed(time.Now().UTC().UnixNano())

	return &IDs[rand.Intn(len(IDs))], nil
}

func createWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	log.Printf("New websocket connect with %s", r.Host)
	go wsListener(conn)
}

func wsListener(conn *websocket.Conn) {
	var err error

	defer conn.Close()

	defer func() {
		if err != nil {
			if err = conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseInternalServerErr, err.Error())); err != nil {
				log.Error(err)
			}
		}
	}()

	for {
		var data []byte
		var t int

		t, data, err = conn.ReadMessage()
		if err != nil {
			log.Errorf("cannot read from websocket %v", err)
			return
		}

		if t == websocket.CloseNormalClosure {
			log.Errorf("closed ws connection with %s", conn.RemoteAddr())
			return
		}

		agentState := &pb.Agentstate{}

		if err = agentState.UnmarshalToStruct(data); err != nil {
			log.Errorf("unmarshal data from websocket error: %v", err)
			return
		}

		if !agentState.CheckToken() {
			log.Error(errors.New("unauthorized agent detected"))
			return
		}

		if err = database.SaveState(agentState); err != nil {
			log.Error(errors.New("save to db problem: "))
			return
		}

		if err := conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%s %s", time.Now().Format("2-Mon-Jan-2006-15:04:05"), "Ok!"))); err != nil {
			log.Error(err)
			return
		}
	}
}

func getFoodInfo(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	ID := &database.UserID{ID: userID}

	foods, err := ID.GetFoodsInFridge()
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	if err = json.NewEncoder(w).Encode(foods); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
	}
}

func getRecipes(w http.ResponseWriter, r *http.Request) {
	recipes, err := database.AllRecipes()

	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	data, err := json.Marshal(recipes)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	if _, err = w.Write(data); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
	}
}

func searchRecipes(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	ID := &database.UserID{ID: userID}

	foods, err := ID.GetFoodsInFridge()
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	recipes, err := database.Recipes(foods)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	if err = json.NewEncoder(w).Encode(recipes); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
	}
}

func addAgent(w http.ResponseWriter, r *http.Request) {

}

func removeAgent(w http.ResponseWriter, r *http.Request) {

}

func updateAgent(w http.ResponseWriter, r *http.Request) {

}

func clientLogin(w http.ResponseWriter, r *http.Request) {
	user := &database.Login{}
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	if err := database.ClientLogin(user.UserName, user.Pass); err != nil {
		sendResponse(w, http.StatusUnauthorized, err)
		return
	}

	userID, err := database.GetUserID(user.UserName)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	if err := sessionSet(w, r, userID); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	sendResponse(w, http.StatusOK, nil)
}

func clientLogout(w http.ResponseWriter, r *http.Request) {
	if err := closeSession(w, r); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	sendResponse(w, http.StatusOK, nil)
}

func clientRegister(w http.ResponseWriter, r *http.Request) {
	newUser := &database.Login{}
	if err := json.NewDecoder(r.Body).Decode(newUser); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	userID, err := database.RegisterNewUser(newUser.UserName, newUser.Pass)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	if err := sessionSet(w, r, userID); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	sendResponse(w, http.StatusOK, nil)
}

func productAdd(w http.ResponseWriter, r *http.Request) {
	newProduct := &database.Product{}

	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	if err := database.AddProduct(newProduct.Name, newProduct.ShelfLife, newProduct.Units); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	sendResponse(w, http.StatusOK, nil)
}

func getAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := database.AllProducts()
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	data, err := json.Marshal(products)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	if _, err = w.Write(data); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
	}
}

func getProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["id"]
	product, err := database.FindProductByID(ID)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	data, err := json.Marshal(product)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	if _, err = w.Write(data); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
	}
}

func getProductByName(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	product, err := database.FindProductByName(name)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	data, err := json.Marshal(product)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	if _, err = w.Write(data); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
	}
}

func productUpdate(w http.ResponseWriter, r *http.Request) {
	newProduct := &database.Product{}

	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	if err := database.UpdateProduct(newProduct.ID, newProduct.Name, newProduct.ShelfLife, newProduct.Units); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	sendResponse(w, http.StatusOK, nil)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["id"]

	if err := database.DeleteProductByID(ID); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	sendResponse(w, http.StatusOK, nil)
}

func getRecipesByProductName(w http.ResponseWriter, r *http.Request) {
	productName := mux.Vars(r)["productName"]
	recipes, err := database.GetRecepiesByProductName(productName)

	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	data, err := json.Marshal(recipes)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	if _, err = w.Write(data); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
	}
}

func recipesByProductNames(w http.ResponseWriter, r *http.Request) {
	var productNames []string
	err := json.NewDecoder(r.Body).Decode(&productNames)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	recipes, err := database.RecepiesByProducts(productNames)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	data, err := json.Marshal(recipes)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
		return
	}

	if _, err = w.Write(data); err != nil {
		sendResponse(w, http.StatusInternalServerError, err)
	}
}