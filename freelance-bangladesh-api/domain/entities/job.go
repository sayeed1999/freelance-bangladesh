package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Job struct {
	ID               uuid.UUID      `gorm:"type:uuid;primaryKey"`
	Title            string         `gorm:"size:255;not null"`
	Description      string         `gorm:"type:text;not null"`
	Budget           float32        `gorm:"not null"`
	Deadline         *time.Time     `gorm:"type:date"`
	ClientKeycloakID string         `gorm:"size:255;not null"` // Keycloak ID of the client
	Status           status         `gorm:"size:20;default:'active'"`
	CreatedAt        time.Time      `gorm:"autoCreateTime"`
	UpdatedAt        time.Time      `gorm:"autoUpdateTime"`
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}

func (job *Job) BeforeCreate(tx *gorm.DB) (err error) {
	job.ID = uuid.New() // Set unique UUID before creating the record
	job.CreatedAt = time.Now().UTC()
	return nil
}

type status string

const (
	DRAFT     status = "draft"
	ACTIVE    status = "active"
	ASSIGNED  status = "assigned"
	COMPLETED status = "completed"
	CLOSED    status = "closed"
)
