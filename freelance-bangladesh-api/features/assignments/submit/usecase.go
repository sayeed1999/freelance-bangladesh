package submit

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sayeed1999/freelance-bangladesh/api/middlewares"
	"github.com/sayeed1999/freelance-bangladesh/database"
	"github.com/sayeed1999/freelance-bangladesh/models"
)

type SubmitWorkRequest struct {
	AssignmentID  string `json:"assignment_id" validate:"required,uuid"`
	SubmissionURL string `json:"work_url" validate:"required,url"`
	Comments      string `json:"comments,omitempty"`
}

type SubmitResponse struct {
	Assignment *models.Assignment
}

type submitWorkUseCase struct{}

func NewSubmitWorkUseCase() *submitWorkUseCase {
	return &submitWorkUseCase{}
}

func (uc *submitWorkUseCase) SubmitWork(ctx context.Context, claims middlewares.Claims, request SubmitWorkRequest) (*SubmitResponse, error) {
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

	var talent models.Talent
	if err := db.First(&talent, "Email = ?", claims.Email).Error; err != nil {
		return nil, fmt.Errorf("failed to get talent: %v", err.Error())
	}

	// Ensure the talent making the request is the owner of the assignment
	if assignment.TalentID != talent.ID {
		return nil, fmt.Errorf("you do not have permission to submit work for this assignment")
	}

	// Ensure the assignment is not already submitted
	if assignment.Status != "PENDING" {
		return nil, fmt.Errorf("assignment has already been submitted or is closed")
	}

	// Update the assignment with the submitted work details
	submittedAt := time.Now().UTC()
	assignment.SubmissionURL = &request.SubmissionURL
	assignment.SubmittedAt = &submittedAt
	assignment.Status = models.SUBMITTED

	// Save the updated assignment
	if err := db.Save(&assignment).Error; err != nil {
		return nil, fmt.Errorf("failed to submit work: %v", err)
	}

	response := &SubmitResponse{Assignment: &assignment}
	return response, nil
}
