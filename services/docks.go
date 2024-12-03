package services

import (
	"context"
	"errors"
	"time"

	"github.com/ItKarma/idocks/models"
	"github.com/ItKarma/idocks/repository"
)

func RegisterDocks(id string, docks models.Dock, repo *repository.DocksRepository) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// verificar se existe uma doca antes de inserir
	_, err := repo.FindDockByName(ctx, docks.Name)

	if err != nil {
		return errors.New("doca já registrada")
	}

	return repo.CreateDocks(ctx, id, docks)
}
