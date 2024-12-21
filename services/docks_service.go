package services

import (
	"context"
	"time"

	"github.com/ItKarma/idocks/models"
	"github.com/ItKarma/idocks/repository"
)

func RegisterDocks(id string, docks models.Dock, repo *repository.DocksRepository) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return repo.CreateDocks(ctx, id, docks)
}

func ListDocks(repo *repository.DocksRepository, id string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Buscar usuário por email
	user, err := repo.ListDocks(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
