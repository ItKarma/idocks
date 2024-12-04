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
