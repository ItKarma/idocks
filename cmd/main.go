package main

import (
	"net/http"

	"github.com/ItKarma/idocks/database"
	"github.com/ItKarma/idocks/handlers"
	"github.com/ItKarma/idocks/middleware"
	"github.com/gorilla/mux"
)

func main() {
	client := database.ConnectDB("mongodb://localhost:27017")
	db := client.Database("auth-api").Collection("users")
	registerDockHandler := middleware.TokenValidationMiddleware(http.HandlerFunc(handlers.RegisterDock(db)))

	r := mux.NewRouter()
	r.HandleFunc("/register", handlers.RegisterHandler(db)).Methods("POST")
	r.HandleFunc("/login", handlers.LoginHandler(db)).Methods("POST")
	r.Handle("/doca/register", registerDockHandler).Methods("POST")

	r.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	http.ListenAndServe(":8000", r)
}
