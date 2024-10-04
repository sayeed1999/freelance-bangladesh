package admindashboarduc

import (
	"context"
	"fmt"

	"github.com/sayeed1999/freelance-bangladesh/database"
	"github.com/sayeed1999/freelance-bangladesh/domain/entities"
)

type getClientsUseCase struct{}

func NewGetClientsUseCase() *getClientsUseCase {
	return &getClientsUseCase{}
}

func (uc *getClientsUseCase) Handler(ctx context.Context) ([]entities.Client, error) {
	db := database.DB.Db

	talents := []entities.Client{}

	if err := db.Find(&talents).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch talents: %v", err)
	}

	return talents, nil
}
