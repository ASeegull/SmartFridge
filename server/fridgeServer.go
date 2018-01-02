package server

import (
	"net/http"

	"github.com/ASeegull/SmartFridge/server/config"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/kabukky/httpscerts"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
)

// Saves methods names to variables to avoid using magic strings
const (
	GET  = http.MethodGet
	POST = http.MethodPost
	DEL  = http.MethodDelete
)

var upgrader websocket.Upgrader

//Run starts server
func Run(cfg config.ServerConfig) error {
	upgrader = websocket.Upgrader{
		ReadBufferSize:  cfg.ReadBufferSize,
		WriteBufferSize: cfg.WriteBufferSize,
	}

	err := httpscerts.Check("cert.pem", "key.pem")
	// If they are not available, generate new ones.
	if err != nil {
		if err = httpscerts.Generate("cert.pem", "key.pem", "127.0.0.1:8081"); err != nil {
			log.Fatal("Error: Couldn't create https certs.")
		}
	}

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200", "http://localhost:5080", "https://aseegull.github.io"},
		AllowCredentials: true,
		AllowedMethods:   []string{"HEAD", "GET", "POST", "DELETE"},
	}).Handler(newRouter())

	log.Printf("Server started on %s:%s", cfg.Host, cfg.Port)
	// return http.ListenAndServeTLS(cfg.Host+":"+cfg.Port, "cert.pem", "key.pem", handler)
	return http.ListenAndServe(cfg.Host+":"+cfg.Port, handler)
}

func newRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/agent", agentAuthentication).Methods(POST)
	router.HandleFunc("/agent", createWS).Methods(GET)

	sub := router.PathPrefix("/client").Subrouter()

	sub.HandleFunc("/allRecipes", getRecipes).Methods(GET)
	sub.HandleFunc("/searchRecipes", checkSession(searchRecipes)).Methods(GET)
	sub.HandleFunc("/fridgeContent", checkSession(getFoodInfo)).Methods(GET)
	sub.HandleFunc("/recipes/getByProductName/{name}", checkSession(getRecipesByProductName)).Methods(GET)
	sub.HandleFunc("/recipes/recipesByProductNames", checkSession(recipesByProductNames)).Methods(POST)

	sub.HandleFunc("/addProduct", checkSession(productAdd)).Methods(POST)
	sub.HandleFunc("/getProducts", checkSession(getAllProducts)).Methods(GET)
	sub.HandleFunc("/updateProduct", checkSession(productUpdate)).Methods(POST)
	sub.HandleFunc("/products/getByID/{id}", checkSession(getProductByID)).Methods(GET)
	sub.HandleFunc("/products/getByName/{name}", checkSession(getProductByName)).Methods(GET)
	sub.HandleFunc("/products/remove/{id}", checkSession(deleteProduct)).Methods(DEL)

	sub.HandleFunc("/addAgent", checkSession(addAgent)).Methods(POST)
	sub.HandleFunc("/removeAgent", checkSession(removeAgent)).Methods(DEL)
	sub.HandleFunc("/updateAgent", checkSession(updateAgent)).Methods(POST)

	sub.HandleFunc("/signup", clientRegister).Methods(POST)
	sub.HandleFunc("/login", clientLogin).Methods(POST)
	sub.HandleFunc("/logout", checkSession(clientLogout)).Methods(POST)

	return router
}
