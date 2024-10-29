package pendingreviewlist

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sayeed1999/freelance-bangladesh/api/middlewares"
	"github.com/sayeed1999/freelance-bangladesh/database"
	"github.com/sayeed1999/freelance-bangladesh/models"
)

type listPendingReviewsUseCase struct{}

func NewPendingReviewListUseCase() *listPendingReviewsUseCase {
	return &listPendingReviewsUseCase{}
}

type PendingReviewResponse struct {
	AssignmentID  uuid.UUID  `json:"assignment_id"`
	JobID         uuid.UUID  `json:"job_id"`
	SubmissionURL *string    `json:"submission_url,omitempty"`
	SubmittedAt   *time.Time `json:"submitted_at,omitempty"`
	Message       string     `json:"message"`
	Status        string     `json:"status"`
}

func (uc *listPendingReviewsUseCase) PendingReviewList(ctx context.Context, claims middlewares.Claims, jobID string) ([]PendingReviewResponse, error) {
	db := database.DB.Db

	// Parse and validate job ID as a UUID
	parsedJobID, err := uuid.Parse(jobID)
	if err != nil {
		return nil, fmt.Errorf("invalid jobID format: %v", err)
	}

	// Retrieve the job by job ID
	var job models.Job
	if err := db.First(&job, "id = ?", parsedJobID).Error; err != nil {
		return nil, fmt.Errorf("job not found: %v", err)
	}

	// Retrieve client information using claims
	var client models.Client
	if err := db.First(&client, "email = ?", claims.Email).Error; err != nil {
		return nil, fmt.Errorf("failed to get client: %v", err)
	}

	// Ensure the client is the owner of the job
	if job.ClientID != client.ID {
		return nil, fmt.Errorf("you do not have permission to fetch pending reviews for this job")
	}

	// Define a slice to hold the result
	var results []PendingReviewResponse

	// Query to retrieve assignments with "SUBMITTED" status for the given job
	if err := db.Table("assignments").
		Select("assignments.id AS assignment_id, assignments.job_id, assignments.submission_url, assignments.submitted_at, assignments.message, assignments.status").
		Where("assignments.job_id = ? AND assignments.status = ?", job.ID, models.SUBMITTED).
		Scan(&results).Error; err != nil {
		return nil, fmt.Errorf("failed to list assignments: %v", err)
	}

	return results, nil
}
