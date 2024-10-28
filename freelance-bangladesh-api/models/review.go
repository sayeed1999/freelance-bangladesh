package models

import "github.com/google/uuid"

type Review struct {
	BaseModel
	AssignmentID uuid.UUID
	Assignment
	Comments string `gorm:"type:text"`
}
