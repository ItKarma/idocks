package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// armazenar as informações da Doca
type Dock struct {
	Name           string    `json:"name_docks"`      // Número identificador da doca
	Status         bool      `json:"status"`          // Status da doca (ocupada ou não)
	PlacaMotorista string    `json:"placa_motorista"` // Placa do motorista utilizando a doca
	HoraEntrada    time.Time `json:"hora_entrada"`    // Hora em que o motorista começou a descarregar
	HoraSaida      time.Time `json:"hora_saida"`      // Hora de saída (quando o descarregamento terminou)
}

// armazenar as informações da Empresa
type Company struct {
	Nome  string `bson:"nome"`  // Nome da empresa
	CNPJ  string `bson:"cnpj"`  // CNPJ da empresa
	Docas []Dock `bson:"docas"` // Lista de docas associadas à empresa
}

// armazenar o Usuário
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"` // Identificador único do usuário
	Email    string             `bson:"email"`         // E-mail do usuário
	Password string             `bson:"password"`      // Senha do usuário
	Company  Company            `bson:"company"`       // Detalhes da empresa associada ao usuário
}
