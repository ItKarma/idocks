package main

import (
	"net/http"

	"github.com/ItKarma/idocks/database"
	"github.com/ItKarma/idocks/handlers"
	"github.com/gorilla/mux"
)

func main() {
	client := database.ConnectDB("mongodb://localhost:27017")
	db := client.Database("auth-api").Collection("users")

	r := mux.NewRouter()
	r.HandleFunc("/register", handlers.RegisterHandler(db)).Methods("POST")
	r.HandleFunc("/login", handlers.LoginHandler(db)).Methods("POST")

	http.ListenAndServe(":8000", r)
}
