package entities

import (
	"time"
)

type Job struct {
	BaseModel
	Title            string     `gorm:"size:255;not null"`
	Description      string     `gorm:"type:text;not null"`
	Budget           float32    `gorm:"not null"`
	Deadline         *time.Time `gorm:"type:date"`
	ClientKeycloakID string     `gorm:"size:255;not null"` // Keycloak ID of the client
	Status           jobStatus  `gorm:"size:20;default:'active'"`
}

type jobStatus string

const (
	DRAFT     jobStatus = "draft"
	ACTIVE    jobStatus = "active"
	ASSIGNED  jobStatus = "assigned"
	COMPLETED jobStatus = "completed"
	CLOSED    jobStatus = "closed"
)
