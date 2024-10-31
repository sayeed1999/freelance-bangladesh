package assignmentlist

import (
	"context"
	"fmt"
	"slices"

	"github.com/google/uuid"
	"github.com/sayeed1999/freelance-bangladesh/api/middlewares"
	"github.com/sayeed1999/freelance-bangladesh/database"
	"github.com/sayeed1999/freelance-bangladesh/models"
	"github.com/sayeed1999/freelance-bangladesh/shared/enums"
)

type listAssignmentsUseCase struct{}

func NewAssignmentListUseCase() *listAssignmentsUseCase {
	return &listAssignmentsUseCase{}
}

type AssignmentResponse struct {
	AssignmentID  uuid.UUID `json:"assignment_id"`
	JobID         uuid.UUID `json:"job_id"`
	Amount        *float32  `json:"amount,omitempty"`
	SubmissionURL *string   `json:"submission_url,omitempty"`
	Status        string    `json:"status"`
}

func (uc *listAssignmentsUseCase) AssignmentList(ctx context.Context, claims middlewares.Claims) ([]AssignmentResponse, error) {
	db := database.DB.Db

	var results []AssignmentResponse

	if slices.Contains(claims.RealmAccess.Roles, string(enums.ROLE_TALENT)) {
		var talent models.Talent
		if err := db.First(&talent, "Email = ?", claims.Email).Error; err != nil {
			return nil, fmt.Errorf("failed to get talent: %v", err.Error())
		}

		if err := db.Table("assignments").
			Select("assignments.id AS assignment_id, assignments.job_id, assignments.budget AS amount, assignments.submission_url, assignments.status").
			Where("assignments.talent_id = ?", talent.ID).
			Scan(&results).Error; err != nil {
			return nil, fmt.Errorf("failed to list assignments for talent: %v", err)
		}
	} else if slices.Contains(claims.RealmAccess.Roles, string(enums.ROLE_CLIENT)) {
		var client models.Client
		if err := db.First(&client, "Email = ?", claims.Email).Error; err != nil {
			return nil, fmt.Errorf("failed to get client: %v", err.Error())
		}

		if err := db.Table("assignments").
			Select("assignments.id AS assignment_id, assignments.job_id, assignments.budget AS amount, assignments.submission_url, assignments.status").
			Joins("JOIN jobs on jobs.id = assignments.job_id").
			Where("jobs.client_id = ?", client.ID).
			Scan(&results).Error; err != nil {
			return nil, fmt.Errorf("failed to list assignments for client: %v", err)
		}
	} else if slices.Contains(claims.RealmAccess.Roles, string(enums.ROLE_ADMIN)) {
		if err := db.Table("assignments").
			Select("assignments.id AS assignment_id, assignments.job_id, assignments.budget AS amount, assignments.submission_url, assignments.status").
			Scan(&results).Error; err != nil {
			return nil, fmt.Errorf("failed to list assignments for admin: %v", err)
		}
	} else {
		return nil, fmt.Errorf("unauthorized access")
	}

	return results, nil
}
