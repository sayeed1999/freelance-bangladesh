package bidjob

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sayeed1999/freelance-bangladesh/api/handlers"
	"github.com/sayeed1999/freelance-bangladesh/api/middlewares"
)

type BidOnJobUseCase interface {
	BidOnJob(ctx context.Context, claims middlewares.Claims, request BidRequest) (*BidResponse, error)
}

func BidOnJobHandler(useCase BidOnJobUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		userClaims, err := handlers.ExtractUserClaims(c)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
		}

		// Extract jobid from URL parameter
		jobID := c.Param("jobid")
		if jobID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "job id is required"})
			return
		}

		// Bind the incoming JSON to the BidRequest struct
		var request BidRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": "unable to parse incoming request: " + err.Error()})
			return
		}

		// set job id on request body
		request.JobID = jobID

		response, err := useCase.BidOnJob(c.Request.Context(), *userClaims, request)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(201, response)
	}
}
