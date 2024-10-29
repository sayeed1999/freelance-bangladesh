package assignmentlist

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sayeed1999/freelance-bangladesh/api/middlewares"
	"github.com/sayeed1999/freelance-bangladesh/database"
	"github.com/sayeed1999/freelance-bangladesh/models"
)

type listAssignmentsUseCase struct{}

func NewAssignmentListUseCase() *listAssignmentsUseCase {
	return &listAssignmentsUseCase{}
}

type AssignmentResponse struct {
	AssignmentID uuid.UUID `json:"assignment_id"`
	JobID        uuid.UUID `json:"job_id"`
	Amount       *float32  `json:"amount,omitempty"`
	Message      string    `json:"message"`
	Status       string    `json:"status"`
}

func (uc *listAssignmentsUseCase) AssignmentList(ctx context.Context, claims middlewares.Claims, talentID string) ([]AssignmentResponse, error) {
	db := database.DB.Db

	// Parse and validate TalentID as a UUID
	parsedTalentID, err := uuid.Parse(talentID)
	if err != nil {
		return nil, fmt.Errorf("invalid TalentID format: %v", err)
	}

	var talent models.Talent
	if err := db.First(&talent, "Email = ?", claims.Email).Error; err != nil {
		return nil, fmt.Errorf("failed to get talent: %v", err.Error())
	}

	// Ensure the talent making the request is the owner of the assignment
	if talentID != talent.ID.String() {
		return nil, fmt.Errorf("you do not have permission to fetch this user's assignments")
	}

	// Define a slice to hold the result with assignment details only
	var results []AssignmentResponse

	// Perform the query to retrieve assignments by talent
	if err := db.Table("assignments").
		Select("assignments.id AS assignment_id, assignments.job_id, assignments.budget AS amount, assignments.message, assignments.status").
		Where("assignments.talent_id = ?", parsedTalentID).
		Scan(&results).Error; err != nil {
		return nil, fmt.Errorf("failed to list assignments: %v", err)
	}

	return results, nil
}
