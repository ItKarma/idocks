package repository

import (
	"context"
	"fmt"

	"github.com/ItKarma/idocks/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	db *mongo.Collection
}

func NewUserRepository(db *mongo.Collection) *UserRepository {
	return &UserRepository{
		db: db,
	}

}

func (r *DocksRepository) FindUserById(ctx context.Context, id string) (*models.User, error) {

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %v", err)
	}

	var user models.User
	err = r.db.FindOne(ctx, bson.M{"_id": objectID}).Decode(&user)
	//fmt.Println("User found:", user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *DocksRepository) FindUserByEmail(ctx context.Context, email string) (*models.User, error) {

	var user models.User
	err := r.db.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	//fmt.Println("User found:", user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
