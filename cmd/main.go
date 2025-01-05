package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ItKarma/idocks/database"
	"github.com/ItKarma/idocks/handlers"
	"github.com/ItKarma/idocks/middleware"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	MONGODB_URL := os.Getenv("MONGODB_URL")
	client := database.ConnectDB(MONGODB_URL)
	db := client.Database("auth-api").Collection("users")

	registerDockHandler := middleware.TokenValidationMiddleware(http.HandlerFunc(handlers.RegisterDock(db)))
	ListDocksHandler := middleware.TokenValidationMiddleware(http.HandlerFunc(handlers.ListDocks(db)))

	r := mux.NewRouter()
	r.HandleFunc("/register", handlers.RegisterHandler(db)).Methods("POST")
	r.HandleFunc("/login", handlers.LoginHandler(db)).Methods("POST")
	r.Handle("/doca/register", registerDockHandler).Methods("POST")
	r.Handle("/doca/list", ListDocksHandler).Methods("GET")

	r.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	http.ListenAndServe(":8000", r)
}
