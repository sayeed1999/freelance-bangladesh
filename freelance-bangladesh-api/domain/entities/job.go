package entities

import (
	"time"

	"github.com/google/uuid"
)

type Job struct {
	BaseModel
	Title       string     `gorm:"size:255;not null"`
	Description string     `gorm:"type:text;not null"`
	Budget      float32    `gorm:"not null"`
	Deadline    *time.Time `gorm:"type:date"`
	ClientID    uuid.UUID
	Client      Client
	Status      jobStatus `gorm:"size:20;default:'active'"`
}

type jobStatus string

const (
	DRAFT     jobStatus = "draft"
	ACTIVE    jobStatus = "active"
	ASSIGNED  jobStatus = "assigned"
	COMPLETED jobStatus = "completed"
	CLOSED    jobStatus = "closed"
)
