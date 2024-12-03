package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ItKarma/idocks/models"
	"github.com/ItKarma/idocks/repository"
	"github.com/ItKarma/idocks/services"
	"go.mongodb.org/mongo-driver/mongo"
)

// Função de Handler para o registro de docas da empresa
func RegisterDock(db *mongo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("userID").(string)

		// Estrutura que vai armazenar os dados recebidos
		var doca models.Dock

		// Decodificar os dados JSON
		err := json.NewDecoder(r.Body).Decode(&doca)
		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		// Verificar os dados recebidos (para depuração)
		fmt.Println(doca)

		repo := repository.NewDocksRepository(db)

		// Chama a função de registro
		err = services.RegisterDocks(userID, doca, repo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Retorna status 201 - Criado
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Dock registered successfully"))
	}
}
