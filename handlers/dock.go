package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ItKarma/idocks/models"
	"github.com/ItKarma/idocks/repository"
	"github.com/ItKarma/idocks/services"
	"go.mongodb.org/mongo-driver/mongo"
)

func convertDockToResponse(dock models.Docks) models.DockResponse {
	return models.DockResponse{
		Name:           dock.Name,
		Status:         dock.Status,
		PlacaMotorista: dock.PlacaMotorista,
		HoraEntrada:    dock.HoraEntrada.Format(time.RFC3339),
		HoraSaida:      dock.HoraSaida.Format(time.RFC3339),
	}
}

// Função de Handler para o registro de docas da empresa
func RegisterDock(db *mongo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("userID").(string)

		// Estrutura que vai armazenar os dados recebidos
		var doca models.Docks

		// Decodificar os dados JSON
		err := json.NewDecoder(r.Body).Decode(&doca)
		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		//fmt.Println(doca)

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

// Função de Handler para o registro de docas da empresa
func EditDock(db *mongo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("userID").(string)

		// Estrutura que vai armazenar os dados recebidos
		var doca models.Docks

		// Decodificar os dados JSON
		err := json.NewDecoder(r.Body).Decode(&doca)
		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		//fmt.Println(doca)

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

// Função de Handler para a listagem  de docas da empresa
func ListDocks(db *mongo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("userID").(string)

		repo := repository.NewDocksRepository(db)

		// Chama a função de registro
		user, err := services.ListDocks(repo, userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var dockResponses []models.DockResponse
		for _, dock := range user.Docas {
			dockResponses = append(dockResponses, convertDockToResponse(dock))
		}

		userListDocks := &models.UserResponse{
			ID:    user.ID,
			Email: user.Email,
			Company: models.CompanyResponse{
				CNPJ: user.Company.CNPJ,
				Nome: user.Company.Nome,
			},
			Docas: dockResponses,
		}

		// Retorna status 201 - Criado
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(userListDocks)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
