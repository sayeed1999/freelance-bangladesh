package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
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

func (u *Bid) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return nil
}
