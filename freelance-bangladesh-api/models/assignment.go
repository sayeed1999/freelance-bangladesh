package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Assignment struct {
	BaseModel
	TalentID      uuid.UUID        `json:"talent_id"`
	Talent        Talent           `json:"talent"`
	JobID         uuid.UUID        `json:"job_id"`
	Job           Job              `json:"job"`
	Budget        float32          `json:"budget" gorm:"not null"`
	AssignedAt    time.Time        `json:"assigned_at" gorm:"type:date"`
	SubmittedAt   *time.Time       `json:"submitted_at,omitempty" gorm:"type:date"`
	SubmissionURL *string          `json:"submission_url,omitempty" gorm:"type:text"`
	Status        submissionStatus `json:"status" gorm:"size:20;default:'pending'"`
}

type submissionStatus string

const (
	PENDING   submissionStatus = "pending"
	SUBMITTED submissionStatus = "submitted"
	APPROVED  submissionStatus = "approved"
	REJECTED  submissionStatus = "rejected"
)

func (u *Assignment) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return nil
}
