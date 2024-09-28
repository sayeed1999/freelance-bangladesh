package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Job struct {
	ID               uuid.UUID       `gorm:"type:uuid;primaryKey"`
	Title            string          `gorm:"size:255;not null"`
	Description      string          `gorm:"type:text;not null"`
	Budget           float32         `gorm:"not null"`
	Deadline         *time.Time      `gorm:"type:date"`
	ClientKeycloakID string          `gorm:"size:255;not null"`        // Keycloak ID of the client
	ProgressStatus   progressStatus  `gorm:"size:20;default:'open'"`   // open, assigned, completed
	LifecycleStatus  lifecycleStatus `gorm:"size:20;default:'active'"` // draft, active, closed
	CreatedAt        time.Time       `gorm:"autoCreateTime"`
}

func (job *Job) BeforeCreate(tx *gorm.DB) (err error) {
	job.ID = uuid.New() // Set unique UUID before creating the record
	return nil
}

type progressStatus string

const (
	OPEN      progressStatus = "open"
	ASSIGNED  progressStatus = "assigned"
	COMPLETED progressStatus = "completed"
)

type lifecycleStatus string

const (
	DRAFT  lifecycleStatus = "draft"
	ACTIVE lifecycleStatus = "active"
	CLOSED lifecycleStatus = "closed"
)
