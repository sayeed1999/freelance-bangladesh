package review

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sayeed1999/freelance-bangladesh/api/middlewares"
	"github.com/sayeed1999/freelance-bangladesh/database"
	"github.com/sayeed1999/freelance-bangladesh/models"
)

type ReviewWorkRequest struct {
	AssignmentID string `json:"assignment_id" validate:"required,uuid"`
	Action       string `json:"action" validate:"required,oneof=approve change-request reject"`
	Comments     string `json:"comments,omitempty"`
}

type ReviewResponse struct {
	Assignment *models.Assignment
}

type reviewWorkUseCase struct{}

func NewReviewWorkUseCase() *reviewWorkUseCase {
	return &reviewWorkUseCase{}
}

func (uc *reviewWorkUseCase) ReviewWork(ctx context.Context, claims middlewares.Claims, request ReviewWorkRequest) (*ReviewResponse, error) {
	db := database.DB.Db

	// Parse and validate assignment ID
	parsedAssignmentID, err := uuid.Parse(request.AssignmentID)
	if err != nil {
		return nil, fmt.Errorf("invalid AssignmentID format: %v", err)
	}

	// Fetch the assignment record by ID
	var assignment models.Assignment
	if err := db.First(&assignment, "id = ?", parsedAssignmentID).Error; err != nil {
		return nil, fmt.Errorf("assignment not found: %v", err)
	}

	// Fetch the job by ID to verify the client
	var job models.Job
	if err := db.First(&job, "id = ?", assignment.JobID).Error; err != nil {
		return nil, fmt.Errorf("job not found: %v", err)
	}

	var client models.Client
	if err := db.First(&client, "Email = ?", claims.Email).Error; err != nil {
		return nil, fmt.Errorf("failed to get client: %v", err.Error())
	}

	// Ensure the client making the request is the owner of the job
	if job.ClientID != client.ID {
		return nil, fmt.Errorf("you do not have permission to review this assignment")
	}

	// Ensure the assignment has been submitted and is awaiting review
	if assignment.Status != models.SUBMITTED {
		return nil, fmt.Errorf("assignment is not in a reviewable state")
	}

	// Update the assignment's review status
	if request.Action == "change-request" {
		assignment.Status = models.PENDING
	} else if request.Action == "approve" {
		assignment.Status = models.APPROVED
	} else if request.Action == "reject" {
		assignment.Status = models.REJECTED
	}

	// Save the updated assignment
	if err := db.Save(&assignment).Error; err != nil {
		return nil, fmt.Errorf("failed to review assignment: %v", err)
	}

	// Add a review for the submitted assignment
	review := &models.Review{
		AssignmentID: assignment.ID,
		Comments:     request.Comments,
	}

	if err := db.Create(&review).Error; err != nil {
		return nil, fmt.Errorf("failed to add review: %v", err)
	}

	response := &ReviewResponse{Assignment: &assignment}
	return response, nil
}
