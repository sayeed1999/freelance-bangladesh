package entities

import (
	"github.com/google/uuid"
)

type Bid struct {
	BaseModel
	JobID    uuid.UUID
	Job      Job
	TalentID uuid.UUID
	Talent   Talent
	Message  string   `gorm:"type:text"` // Optional message from talent
	Amount   *float32 `gorm:""`          // Optional bid amount (pointer to float32)
}
