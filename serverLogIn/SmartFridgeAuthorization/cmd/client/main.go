package main

import (
	"net/http"


	"github.com/gorilla/mux"
	"github.com/ASeegull2/serverLogIn/SmartFridgeAuthorization/serverLogIn"
)



func main() {
	var router = mux.NewRouter()


	router.HandleFunc("/", serverLogIn.IndexPageHandler)

	router.HandleFunc("/signup", serverLogIn.IndexPageTwoHandler)

	router.HandleFunc("/internal", serverLogIn.InternalPageHandler)



	router.HandleFunc("/register", serverLogIn.SignUpHandler).Methods("POST")
	router.HandleFunc("/login", serverLogIn.LoginHandler).Methods("POST")
	router.HandleFunc("/logout", serverLogIn.LogoutHandler).Methods("POST")

	http.Handle("/", router)
	http.ListenAndServe(":8000", nil)

}
