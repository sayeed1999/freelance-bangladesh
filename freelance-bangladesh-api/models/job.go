package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Job struct {
	BaseModel
	Title       string     `json:"title" gorm:"size:255;not null"`
	Description string     `json:"description" gorm:"type:text;not null"`
	Budget      float32    `json:"budget" gorm:"not null"`
	Deadline    *time.Time `json:"deadline,omitempty" gorm:"type:date"`
	ClientID    uuid.UUID  `json:"client_id"`
	Client      Client     `json:"client"`
	Status      jobStatus  `json:"status" gorm:"size:20;default:'active'"`
}

type jobStatus string

const (
	DRAFT     jobStatus = "draft"
	ACTIVE    jobStatus = "active"
	ASSIGNED  jobStatus = "assigned"
	COMPLETED jobStatus = "completed"
	CLOSED    jobStatus = "closed"
)

func (u *Job) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return nil
}
