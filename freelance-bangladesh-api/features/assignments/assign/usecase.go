package assign

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sayeed1999/freelance-bangladesh/api/middlewares"
	"github.com/sayeed1999/freelance-bangladesh/database"
	"github.com/sayeed1999/freelance-bangladesh/models"
)

type AssignRequest struct {
	JobID    string  `json:"job_id" validate:"required,uuid"`
	TalentID string  `json:"talent_id" validate:"required,uuid"`
	Amount   float32 `json:"amount" validate:"required,gt=0"`
}

type AssignResponse struct {
	Job *models.Job
}

type assignTalentUseCase struct{}

func NewAssignTalentUseCase() *assignTalentUseCase {
	return &assignTalentUseCase{}
}

func (uc *assignTalentUseCase) AssignTalent(ctx context.Context, claims middlewares.Claims, request AssignRequest) (*AssignResponse, error) {
	db := database.DB.Db

	// Validate and parse jobID and talentID
	parsedJobID, err := uuid.Parse(request.JobID)
	if err != nil {
		return nil, fmt.Errorf("invalid JobID format: %v", err)
	}
	parsedTalentID, err := uuid.Parse(request.TalentID)
	if err != nil {
		return nil, fmt.Errorf("invalid TalentID format: %v", err)
	}

	// Fetch job by ID and check if it exists
	var job models.Job
	if err := db.First(&job, "id = ?", parsedJobID).Error; err != nil {
		return nil, fmt.Errorf("job not found: %v", err)
	}

	var client models.Client
	if err := db.First(&client, "Email = ?", claims.Email).Error; err != nil {
		return nil, fmt.Errorf("failed to get client: %v", err.Error())
	}

	// Ensure the user making the request is the job's owner
	if job.ClientID != client.ID {
		return nil, fmt.Errorf("you do not have permission to assign talent to this job")
	}

	// Check if the job is still active and can be assigned
	if job.Status != models.ACTIVE {
		return nil, fmt.Errorf("job is not open for assignment")
	}

	// Fetch the talent by ID and verify they exist and are verified
	var talent models.Talent
	if err := db.First(&talent, "id = ?", parsedTalentID).Error; err != nil {
		return nil, fmt.Errorf("talent not found: %v", err)
	}
	if !talent.IsVerified {
		return nil, fmt.Errorf("cannot assign an unverified talent")
	}

	// Update the job's status to assigned and set the talent
	job.Status = models.ASSIGNED
	if err := db.Save(&job).Error; err != nil {
		return nil, fmt.Errorf("failed to assign talent: %v", err)
	}

	// Create the assignment record, including the assigned amount
	assignment := &models.Assignment{
		JobID:      parsedJobID,
		TalentID:   parsedTalentID,
		Budget:     request.Amount, // Assign the provided amount
		AssignedAt: time.Now().UTC(),
	}
	if err := db.Create(&assignment).Error; err != nil {
		return nil, fmt.Errorf("failed to create assignment: %v", err)
	}

	response := &AssignResponse{Job: &job}
	return response, nil
}
