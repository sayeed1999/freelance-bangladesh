package auth

import (
	"context"
	"errors"
	"fmt"
	"slices"

	"github.com/sayeed1999/freelance-bangladesh/api/middlewares"
	"github.com/sayeed1999/freelance-bangladesh/database"
	"github.com/sayeed1999/freelance-bangladesh/models"
	"github.com/sayeed1999/freelance-bangladesh/shared/enums"
	"gorm.io/gorm"
)

type SyncUserResponse struct {
	Success bool `json:"success"`
}

type syncUserUseCase struct{}

func NewSyncUserUseCase() *syncUserUseCase {
	return &syncUserUseCase{}
}

func (uc *syncUserUseCase) SyncUser(ctx context.Context, claims middlewares.Claims) (*SyncUserResponse, error) {
	db := database.DB.Db

	fmt.Println(claims.RealmAccess.Roles)

	isAdmin := slices.Contains(claims.RealmAccess.Roles, string(enums.ROLE_ADMIN))
	isClient := slices.Contains(claims.RealmAccess.Roles, string(enums.ROLE_CLIENT))

	// Note: Ensure to let only users with no role insert into 'else' block.
	if isAdmin {
		// do nothing
	} else if isClient {
		var client models.Client

		if err := db.First(&client, "Email = ?", claims.Email).Error; err != nil {

			if errors.Is(err, gorm.ErrRecordNotFound) {
				client = models.Client{
					Email:      claims.Email,
					Name:       claims.FirstName + " " + claims.LastName,
					IsVerified: true,
					// KeycloakUserID: ,
				}

				if err := db.Create(&client).Error; err != nil {
					return nil, fmt.Errorf("failed to create client: %v", err)
				}
			} else {
				return nil, fmt.Errorf("failed to get client: %v", err)
			}
		}
	} else {
		var talent models.Talent

		if err := db.First(&talent, "Email = ?", claims.Email).Error; err != nil {

			if errors.Is(err, gorm.ErrRecordNotFound) {
				talent = models.Talent{
					Email:      claims.Email,
					Name:       claims.FirstName + " " + claims.LastName,
					IsVerified: true,
					// KeycloakUserID: ,
				}

				if err := db.Create(&talent).Error; err != nil {
					return nil, fmt.Errorf("failed to create talent: %v", err)
				}
			} else {
				return nil, fmt.Errorf("failed to get talent: %v", err)
			}
		}
	}

	response := &SyncUserResponse{Success: true}
	return response, nil
}
