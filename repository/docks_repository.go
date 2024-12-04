package repository

import (
	"context"
	"fmt"

	"github.com/ItKarma/idocks/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DocksRepository struct {
	db *mongo.Collection
}

func NewDocksRepository(db *mongo.Collection) *DocksRepository {
	return &DocksRepository{db: db}
}

func (r *DocksRepository) CreateDocks(ctx context.Context, userId string, dock models.Dock) error {
	// Buscar o usuário com o id
	user, err := r.FindUserById(ctx, userId)
	if err != nil {
		return fmt.Errorf("error fetching user: %v", err)
	}
	if user == nil {
		return fmt.Errorf("user not found with id: %v", userId)
	}

	// Verifica se já existe uma doca com o mesmo nomewaa
	for _, existingDock := range user.Docas {
		if existingDock.Name == dock.Name {
			return fmt.Errorf("dock with name '%s' already exists", dock.Name)
		}
	}

	// Se não existir, adiciona a nova doca à lista
	user.Docas = append(user.Docas, dock)

	// Atualiza a empresa no banco, incluindo a nova doca
	_, err = r.db.UpdateOne(ctx,
		bson.M{"_id": user.ID},
		bson.M{"$set": bson.M{"docas": user.Docas}})

	if err != nil {
		return fmt.Errorf("erro ao atualizar empresa: %v", err)
	}

	return nil
}
