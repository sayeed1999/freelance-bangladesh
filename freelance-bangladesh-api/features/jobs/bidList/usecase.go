package bidlist

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sayeed1999/freelance-bangladesh/database"
)

type listBidsUseCase struct{}

func NewBidListUseCase() *listBidsUseCase {
	return &listBidsUseCase{}
}

type BidResponse struct {
	BidID       uuid.UUID `json:"bid_id"`
	JobID       uuid.UUID `json:"job_id"`
	Amount      *float32  `json:"amount,omitempty"`
	Message     string    `json:"message"`
	TalentID    uuid.UUID `json:"talent_id"`
	TalentName  string    `json:"talent_name"`
	TalentEmail string    `json:"talent_email"`
}

func (uc *listBidsUseCase) BidList(ctx context.Context, jobID string) ([]BidResponse, error) {
	db := database.DB.Db

	// Parse and validate JobID as a UUID
	parsedJobID, err := uuid.Parse(jobID)
	if err != nil {
		return nil, fmt.Errorf("invalid JobID format: %v", err)
	}

	// Define a slice to hold the result with bid and talent details
	var results []BidResponse

	// Perform the join query
	if err := db.Table("bids").
		Select("bids.id AS bid_id, bids.job_id, bids.amount, bids.message, bids.talent_id, talents.name AS talent_name, talents.email AS talent_email").
		Joins("JOIN talents ON bids.talent_id = talents.id").
		Where("bids.job_id = ?", parsedJobID).
		Scan(&results).Error; err != nil {
		return nil, fmt.Errorf("failed to list bids: %v", err)
	}

	return results, nil
}
