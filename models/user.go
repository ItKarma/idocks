package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// "Esqueleto" de como deve ser armazenado no mongodb
type User struct {
	ID       primitive.ObjectID `bson:"_id, omitempty"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
}
