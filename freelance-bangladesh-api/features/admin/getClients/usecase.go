package getclients

import (
	"context"
	"fmt"

	"github.com/sayeed1999/freelance-bangladesh/database"
	"github.com/sayeed1999/freelance-bangladesh/models"
)

type getClientsUseCase struct{}

func NewGetClientsUseCase() *getClientsUseCase {
	return &getClientsUseCase{}
}

func (uc *getClientsUseCase) Handler(ctx context.Context) ([]models.Client, error) {
	db := database.DB.Db

	talents := []models.Client{}

	if err := db.Find(&talents).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch talents: %v", err)
	}

	return talents, nil
}
