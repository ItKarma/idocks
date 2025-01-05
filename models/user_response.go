package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DockResponse struct {
	Name           string `json:"name_docks"`
	Status         bool   `json:"status"`
	PlacaMotorista string `json:"placa_motorista"`
	HoraEntrada    string `json:"hora_entrada"` // Talvez seja melhor em formato de string
	HoraSaida      string `json:"hora_saida"`   // Talvez seja melhor em formato de string
}

type CompanyResponse struct {
	Nome string `json:"nome"` // Nome da empresa
	CNPJ string `json:"cnpj"` // CNPJ da empresa
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
