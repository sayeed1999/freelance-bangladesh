package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Bid struct {
	BaseModel
	JobID    uuid.UUID `json:"job_id"`
	Job      Job       `json:"job"`
	TalentID uuid.UUID `json:"talent_id"`
	Talent   Talent    `json:"talent"`
	Message  string    `json:"message" gorm:"type:text"` // Optional message from talent
	Amount   *float32  `json:"amount,omitempty" gorm:""` // Optional bid amount (pointer to float32)
}

func (u *Bid) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return nil
}
