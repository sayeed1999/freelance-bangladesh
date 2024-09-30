package entities

import (
	"github.com/google/uuid"
)

type Bid struct {
	BaseModel
	JobID            uuid.UUID `gorm:"type:uuid;not null"` // Foreign key to Job
	TalentKeycloakID string    `gorm:"size:255;not null"`  // Keycloak ID of the talent
	Message          string    `gorm:"type:text"`          // Optional message from talent
	Amount           *float32  `gorm:""`                   // Optional bid amount (pointer to float32)
}
