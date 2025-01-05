package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DockResponse struct {
	Name           string    `json:"name_docks"`             // Nome da doca
	Status         bool      `json:"status"`                 // Status da doca
	PlacaMotorista string    `json:"placa_motorista"`        // Placa do motorista
	HoraEntrada    time.Time `json:"hora_entrada,omitempty"` // Hora de início
	HoraSaida      time.Time `json:"hora_saida,omitempty"`   // Hora de saída
}

type CompanyResponse struct {
	Nome  string         `json:"nome"`  // Nome da empresa
	CNPJ  string         `json:"cnpj"`  // CNPJ da empresa
	Docas []DockResponse `json:"docas"` // Lista de docas
}

type UserResponse struct {
	ID      primitive.ObjectID `json:"id"`      // Identificador único
	Email   string             `json:"email"`   // E-mail do usuário
	Company CompanyResponse    `json:"company"` // Empresa associada
	Docas   []DockResponse     `json:"docas"`   // Docas associadas
}

type UserLoginResponse struct {
	ID      primitive.ObjectID `json:"id"`      // Identificador único
	Email   string             `json:"email"`   // E-mail do usuário
	Company CompanyResponse    `json:"company"` // Empresa associada
	Docas   []DockResponse     `json:"docas"`   // Docas associadas
}
