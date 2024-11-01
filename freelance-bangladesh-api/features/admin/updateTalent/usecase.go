package updatetalent

import (
	"context"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/sayeed1999/freelance-bangladesh/api/middlewares"
	"github.com/sayeed1999/freelance-bangladesh/database"
	"github.com/sayeed1999/freelance-bangladesh/models"
)

type updateTalentUseCase struct{}

func NewUpdateTalentUseCase() *updateTalentUseCase {
	return &updateTalentUseCase{}
}

type UpdateTalentCommand struct {
	TalentID   string `json:"talent_id" validate:"required,max=36"`
	IsVerified *bool  `json:"is_verified"` // Pointer will be null if not provided by user
}

func (uc *updateTalentUseCase) Handler(ctx context.Context, claims middlewares.Claims, command UpdateTalentCommand) error {
	db := database.DB.Db

	var validate = validator.New()
	err := validate.Struct(command)
	if err != nil {
		fmt.Println(err)

		return fmt.Errorf("failed to validate command: %v", err)
	}

	var talent models.Talent

	if err = db.First(&talent, "ID = ?", command.TalentID).Error; err != nil {
		return fmt.Errorf("failed to find talent: %v", err)
	}

	// PATCH updates
	if command.IsVerified != nil {
		talent.IsVerified = *command.IsVerified
	}

	if err = db.Save(talent).Error; err != nil {
		return fmt.Errorf("failed to update talent entity: %v", err)
	}

	return nil
}
