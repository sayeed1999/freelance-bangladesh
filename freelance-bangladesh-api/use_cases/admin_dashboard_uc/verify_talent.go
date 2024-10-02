package admindashboarduc

import (
	"context"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/sayeed1999/freelance-bangladesh/api/middlewares"
	"github.com/sayeed1999/freelance-bangladesh/database"
	"github.com/sayeed1999/freelance-bangladesh/domain/entities"
)

type verifyTalentUseCase struct{}

func NewVerifyTalentUseCase() *verifyTalentUseCase {
	return &verifyTalentUseCase{}
}

type VerifyTalentCommand struct {
	TalentID uuid.UUID `validate:"required,uuid"`
}

func (uc *verifyTalentUseCase) Handler(ctx context.Context, claims middlewares.Claims, command VerifyTalentCommand) error {
	db := database.DB.Db

	var validate = validator.New()
	err := validate.Struct(command)
	if err != nil {
		return fmt.Errorf("failed to validate command: %v", err)
	}

	var talent entities.Talent

	if err = db.First(&talent, "ID = ?", command.TalentID).Error; err != nil {
		return fmt.Errorf("failed to find talent: %v", err)
	}

	// Update the prop
	talent.IsVerified = true

	if err = db.Save(talent).Error; err != nil {
		return fmt.Errorf("failed to update talent entity: %v", err)
	}

	return nil
}
