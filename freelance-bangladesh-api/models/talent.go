package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Talent struct {
	BaseModel
	KeycloakUserID uuid.UUID `gorm:"uniqueIndex;not null" json:"keycloak_user_id"`
	Email          string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	Name           string    `gorm:"type:varchar(50);not null" json:"name"`
	Phone          string    `gorm:"type:varchar(20);not null" json:"phone"`
	IsVerified     bool      `gorm:"default:false" json:"is_verified"`
}

func (u *Talent) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return nil
}
