package main

import (

	"net/http"


	"github.com/gorilla/mux"
	"github.com/ASeegull2/SmartFridge/serverLogIn"
)


var router = mux.NewRouter()

func main() {

	router.HandleFunc("/", serverLogIn.IndexPageHandler)
	router.HandleFunc("/register", serverLogIn.IndexPageTwoHandler)
	router.HandleFunc("/internal", serverLogIn.InternalPageHandler)


	router.HandleFunc("/signup", serverLogIn.SignUpHandler).Methods("POST")
	router.HandleFunc("/login", serverLogIn.LoginHandler).Methods("POST")
	router.HandleFunc("/logout", serverLogIn.LogoutHandler).Methods("POST")

	http.Handle("/", router)
	http.ListenAndServe(":8000", nil)

}
