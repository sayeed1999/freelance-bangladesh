package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Talent struct {
	BaseModel
	KeycloakUserID uuid.UUID `gorm:"uniqueIndex;not null"`
	Email          string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	Name           string    `gorm:"type:varchar(50);not null"`
	Phone          string    `gorm:"type:varchar(20);not null"`
	IsVerified     bool      `gorm:"default:false"`
}

func (u *Talent) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return nil
}
