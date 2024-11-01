package bidjob

import (
	"context"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/sayeed1999/freelance-bangladesh/api/middlewares"
	"github.com/sayeed1999/freelance-bangladesh/database"
	"github.com/sayeed1999/freelance-bangladesh/models"
)

type BidRequest struct {
	JobID   string   `json:"job_id" validate:"required,uuid"`  // Ensure it's a valid UUID
	Amount  *float32 `json:"amount" validate:"omitempty,gt=0"` // Optional, but must be greater than 0 if provided
	Message string   `json:"message" validate:"max=1000"`      // Message with a max length of 1000 characters
}

type BidResponse struct {
	Bid *models.Bid
}

type bidOnJobUseCase struct{}

func NewBidOnJobUseCase() *bidOnJobUseCase {
	return &bidOnJobUseCase{}
}

func (uc *bidOnJobUseCase) BidOnJob(ctx context.Context, claims middlewares.Claims, request BidRequest) (*BidResponse, error) {
	db := database.DB.Db

	// Validate the request
	var validate = validator.New()
	if err := validate.Struct(request); err != nil {
		return nil, fmt.Errorf("validation failed: %v", err.Error())
	}

	// Parse and validate JobID as a UUID
	jobID, err := uuid.Parse(request.JobID)
	if err != nil {
		return nil, fmt.Errorf("invalid JobID format: %v", err)
	}

	// Fetch the talent by email from the claims
	var talent models.Talent
	if err := db.First(&talent, "Email = ?", claims.Email).Error; err != nil {
		return nil, fmt.Errorf("failed to get talent: %v", err)
	}

	// Ensure the talent's account is verified
	if !talent.IsVerified {
		return nil, fmt.Errorf("failed to place bid: talent account is not verified")
	}

	// Fetch the job by UUID (JobID)
	var job models.Job
	if err := db.First(&job, "id = ?", jobID).Error; err != nil {
		return nil, fmt.Errorf("job not found: %v", err)
	}

	// Ensure the job is open for bidding
	if job.Status != models.ACTIVE {
		return nil, fmt.Errorf("job is not open for bidding")
	}

	// Check if the talent has already placed a bid on this job
	var existingBid models.Bid
	if err := db.First(&existingBid, "job_id = ? AND talent_id = ?", jobID, talent.ID).Error; err == nil {
		return nil, fmt.Errorf("you have already placed a bid on this job")
	}

	// Create a new bid
	bid := &models.Bid{
		JobID:    jobID,     // Use UUID for JobID
		TalentID: talent.ID, // TalentID is also a UUID
		Amount:   request.Amount,
		Message:  request.Message,
	}

	if err := db.Create(&bid).Error; err != nil {
		return nil, fmt.Errorf("failed to place bid: %v", err)
	}

	response := &BidResponse{Bid: bid}
	return response, nil
}
