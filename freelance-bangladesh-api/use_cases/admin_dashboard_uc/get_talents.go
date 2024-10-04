package admindashboarduc

import (
	"context"
	"fmt"

	"github.com/sayeed1999/freelance-bangladesh/database"
	"github.com/sayeed1999/freelance-bangladesh/domain/entities"
)

type getTalentsUseCase struct{}

func NewGetTalentsUseCase() *getTalentsUseCase {
	return &getTalentsUseCase{}
}

func (uc *getTalentsUseCase) Handler(ctx context.Context) ([]entities.Talent, error) {
	db := database.DB.Db

	talents := []entities.Talent{}

	if err := db.Find(&talents).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch talents: %v", err)
	}

	return talents, nil
}
