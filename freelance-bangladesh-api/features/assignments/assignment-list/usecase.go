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
	Status       string    `json:"status"`
}

func (uc *listAssignmentsUseCase) AssignmentList(ctx context.Context, claims middlewares.Claims) ([]AssignmentResponse, error) {
	db := database.DB.Db

	// the talent is defined by the user claims, no extra params needed!
	var talent models.Talent
	if err := db.First(&talent, "Email = ?", claims.Email).Error; err != nil {
		return nil, fmt.Errorf("failed to get talent: %v", err.Error())
	}

	// Define a slice to hold the result with assignment details only
	var results []AssignmentResponse

	// Perform the query to retrieve assignments by talent
	if err := db.Table("assignments").
		Select("assignments.id AS assignment_id, assignments.job_id, assignments.budget AS amount, assignments.status").
		Where("assignments.talent_id = ?", talent.ID).
		Scan(&results).Error; err != nil {
		return nil, fmt.Errorf("failed to list assignments: %v", err)
	}

	return results, nil
}
