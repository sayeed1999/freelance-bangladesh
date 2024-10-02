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

type verifyClientUseCase struct{}

func NewVerifyClientUseCase() *verifyClientUseCase {
	return &verifyClientUseCase{}
}

type VerifyClientCommand struct {
	ClientID uuid.UUID `validate:"required,uuid"`
}

func (uc *verifyClientUseCase) Handler(ctx context.Context, claims middlewares.Claims, command VerifyClientCommand) error {
	db := database.DB.Db

	var validate = validator.New()
	err := validate.Struct(command)
	if err != nil {
		return fmt.Errorf("failed to validate command: %v", err)
	}

	var client entities.Client

	if err = db.First(&client, "ID = ?", command.ClientID).Error; err != nil {
		return fmt.Errorf("failed to find client: %v", err)
	}

	// Update the prop
	client.IsVerified = true

	if err = db.Save(client).Error; err != nil {
		return fmt.Errorf("failed to update client entity: %v", err)
	}

	return nil
}
