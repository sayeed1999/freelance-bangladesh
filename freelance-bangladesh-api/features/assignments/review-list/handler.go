package reviewlist

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReviewListUseCase interface {
	GetReviewList(ctx context.Context, assignmentID string) ([]ReviewResponse, error)
}

func ReviewListHandler(useCase ReviewListUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		assignmentID := c.Param("assignmentid")
		if assignmentID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "assignment id is required"})
			return
		}

		// Call the use case to list bids for the specified job
		bids, err := useCase.GetReviewList(c.Request.Context(), assignmentID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"total":  len(bids),
			"result": bids,
		})
	}
}
