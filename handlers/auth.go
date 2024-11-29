package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ItKarma/idocks/services"
	"go.mongodb.org/mongo-driver/mongo"
)

// func para lidar com os registro de usuarios
func RegisterHandler(db *mongo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// estrutura vai armazenar os dados recebidos
		var data struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		// Decodificar os dados json
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		}

		// Chama a função de registro
		err = services.RegisterUser(db, data.Email, data.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		//Retorn status 201 - criado
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("User registered successfully"))
	}
}

// Func para lidar com autenticação

func LoginHandler(db *mongo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		// Decodifica o corpo JSON da requisição
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		// Chama a função de login no serviço
		token, err := services.LoginUser(db, data.Email, data.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// Define o cabeçalho Content-Type para JSON
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Retorna o token em formato JSON
		json.NewEncoder(w).Encode(map[string]string{"token": token})
	}
}
