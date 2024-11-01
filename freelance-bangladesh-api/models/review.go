package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Review struct {
	BaseModel
	AssignmentID uuid.UUID  `json:"assignment_id"`
	Assignment   Assignment `json:"assignment"`
	Comments     string     `json:"comments" gorm:"type:text"`
}

func (u *Review) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return nil
}
