package gettalents

import (
	"context"
	"fmt"

	"github.com/sayeed1999/freelance-bangladesh/database"
	"github.com/sayeed1999/freelance-bangladesh/models"
)

type getTalentsUseCase struct{}

func NewGetTalentsUseCase() *getTalentsUseCase {
	return &getTalentsUseCase{}
}

func (uc *getTalentsUseCase) Handler(ctx context.Context) ([]models.Talent, error) {
	db := database.DB.Db

	talents := []models.Talent{}

	if err := db.Find(&talents).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch talents: %v", err)
	}

	return talents, nil
}
