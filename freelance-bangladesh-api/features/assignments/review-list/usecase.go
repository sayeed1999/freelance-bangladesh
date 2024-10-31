package reviewlist

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sayeed1999/freelance-bangladesh/database"
)

type reviewListUseCase struct{}

func NewReviewListUseCase() *reviewListUseCase {
	return &reviewListUseCase{}
}

type ReviewResponse struct {
	ReviewID uuid.UUID `json:"review_id"`
	Comments string    `json:"comments"`
}

func (uc *reviewListUseCase) GetReviewList(ctx context.Context, assignmentID string) ([]ReviewResponse, error) {
	db := database.DB.Db

	// Parse and validate JobID as a UUID
	parsedAssignmentID, err := uuid.Parse(assignmentID)
	if err != nil {
		return nil, fmt.Errorf("invalid AssignmentID format: %v", err)
	}

	var results []ReviewResponse

	if err := db.Table("reviews").
		Select("reviews.id as review_id, reviews.comments").
		Where("reviews.assignment_id = ?", parsedAssignmentID).
		Scan(&results).Error; err != nil {
		return nil, fmt.Errorf("failed to list reviews: %v", err)
	}

	return results, nil
}
