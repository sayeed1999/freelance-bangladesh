package pendingreviewlist

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sayeed1999/freelance-bangladesh/api/handlers"
	"github.com/sayeed1999/freelance-bangladesh/api/middlewares"
)

type PendingReviewListUseCase interface {
	PendingReviewList(ctx context.Context, claims middlewares.Claims, jobID string) ([]PendingReviewResponse, error)
}

func PendingReviewListHandler(useCase PendingReviewListUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract user claims from the context
		userClaims, err := handlers.ExtractUserClaims(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Get the job ID from the URL parameters
		jobID := c.Param("jobID")
		if jobID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "jobID is required"})
			return
		}

		// Call the use case to fetch pending reviews
		response, err := useCase.PendingReviewList(c.Request.Context(), *userClaims, jobID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Return the response
		c.JSON(http.StatusOK, response)
	}
}
