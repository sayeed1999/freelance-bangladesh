package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Client struct {
	BaseModel
	// KeycloakUserID *uuid.UUID `json:"keycloak_user_id" gorm:"uniqueIndex;"`
	Email      string `json:"email" gorm:"type:varchar(100);uniqueIndex;not null"`
	Name       string `json:"name" gorm:"type:varchar(50);not null"`
	Phone      string `json:"phone" gorm:"type:varchar(20);not null"`
	IsVerified bool   `json:"is_verified" gorm:"default:false"`
}

func (u *Client) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return nil
}
