package updateclient

import (
	"context"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/sayeed1999/freelance-bangladesh/api/middlewares"
	"github.com/sayeed1999/freelance-bangladesh/database"
	"github.com/sayeed1999/freelance-bangladesh/models"
)

type updateClientUseCase struct{}

func NewUpdateClientUseCase() *updateClientUseCase {
	return &updateClientUseCase{}
}

type UpdateClientCommand struct {
	ClientID   uuid.UUID `validate:"required,uuid"`
	IsVerified *bool     // Pointer will be null if not provided by user
}

func (uc *updateClientUseCase) Handler(ctx context.Context, claims middlewares.Claims, command UpdateClientCommand) error {
	db := database.DB.Db

	var validate = validator.New()
	err := validate.Struct(command)
	if err != nil {
		return fmt.Errorf("failed to validate command: %v", err)
	}

	var client models.Client

	if err = db.First(&client, "ID = ?", command.ClientID).Error; err != nil {
		return fmt.Errorf("failed to find client: %v", err)
	}

	// PATCH updates
	if command.IsVerified != nil {
		client.IsVerified = *command.IsVerified
	}

	if err = db.Save(client).Error; err != nil {
		return fmt.Errorf("failed to update client entity: %v", err)
	}

	return nil
}
