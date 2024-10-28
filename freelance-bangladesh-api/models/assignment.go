package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Assignment struct {
	BaseModel
	TalentID      uuid.UUID
	Talent        Talent
	JobID         uuid.UUID
	Job           Job
	Budget        float32           `gorm:"not null"`
	AssignedAt    time.Time         `gorm:"type:date"`
	SubmittedAt   *time.Time        `gorm:"type:date"`
	SubmissionURL *string           `gorm:"type:text"`
	Status        *submissionStatus `gorm:"size:20"`
}

type submissionStatus string

const (
	SUBMITTED submissionStatus = "submitted"
	APPROVED  submissionStatus = "approved"
	REJECTED  submissionStatus = "rejected"
)

func (u *Assignment) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return nil
}
