package createjob

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sayeed1999/freelance-bangladesh/api/handlers"
	"github.com/sayeed1999/freelance-bangladesh/api/middlewares"
)

type CreateJobUseCase interface {
	CreateJob(ctx context.Context, claims middlewares.Claims, request CreateJobRequest) (*CreateJobResponse, error)
}

func CreateJobHandler(useCase CreateJobUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		userClaims, err := handlers.ExtractUserClaims(c)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
		}

		var request CreateJobRequest

		// Bind the incoming JSON to the request struct
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": errors.Wrap(err, "unable to parse incoming request").Error()})
			return
		}

		// Create job using the use case
		response, err := useCase.CreateJob(c.Request.Context(), *userClaims, request)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// Respond with the created job
		c.JSON(201, response)
	}
}
