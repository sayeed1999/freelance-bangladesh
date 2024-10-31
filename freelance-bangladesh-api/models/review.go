package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Review struct {
	BaseModel
	AssignmentID uuid.UUID
	Assignment   Assignment
	Comments     string `gorm:"type:text"`
}

func (u *Review) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return nil
}
