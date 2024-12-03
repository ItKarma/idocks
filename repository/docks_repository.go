package repository

import (
	"context"
	"fmt"

	"github.com/ItKarma/idocks/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DocksRepository struct {
	db *mongo.Collection
}

func NewDocksRepository(db *mongo.Collection) *DocksRepository {
	return &DocksRepository{db: db}
}

func (r *DocksRepository) CreateDocks(ctx context.Context, id string, dock models.Dock) error {

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid ID format: %v", err)
	}

	var company models.Company
	err = r.db.FindOne(ctx, bson.M{"_id": objectID}).Decode(&company)
	if err != mongo.ErrNoDocuments && err != nil {

		return fmt.Errorf("failed to create company: %v", err)
	}

	fmt.Println(company)

	if company.Docas == nil {
		company.Docas = []models.Dock{}
	}

	company.Docas = append(company.Docas, dock)

	// Atualiza a empresa no banco, incluindo a lista de docas
	_, err = r.db.UpdateOne(ctx,
		bson.M{"_id": objectID},                        // Filtro para encontrar a empresa pelo ID
		bson.M{"$set": bson.M{"docas": company.Docas}}) // Atualizar o campo 'docas'

	if err != nil {
		return fmt.Errorf("erro ao atualizar empresa: %v", err)
	}

	return nil
}

func (r *DocksRepository) FindDockByName(ctx context.Context, name string) (*models.Dock, error) {
	var dock models.Dock
	err := r.db.FindOne(ctx, bson.M{"Name": name}).Decode(&dock)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &dock, nil
}
