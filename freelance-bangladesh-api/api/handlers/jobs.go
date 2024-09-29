package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sayeed1999/freelance-bangladesh/api/middlewares"
	"github.com/sayeed1999/freelance-bangladesh/domain/entities"
	jobsuc "github.com/sayeed1999/freelance-bangladesh/use_cases/jobs_uc"
)

// CreateJobUseCase interface for creating a job
type CreateJobUseCase interface {
	CreateJob(ctx context.Context, claims middlewares.Claims, request jobsuc.CreateJobRequest) (*jobsuc.CreateJobResponse, error)
}

// GetJobsUseCase interface for getting jobs
type GetJobsUseCase interface {
	GetJobs(ctx context.Context) ([]entities.Job, error)
}

type GetActiveJobsUseCase interface {
	GetActiveJobs(ctx context.Context, userClaims middlewares.Claims) ([]entities.Job, error)
}

func CreateJobHandler(useCase CreateJobUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, exists := c.Get("userClaims")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "no claims found"})
		}

		userClaims, ok := claims.(middlewares.Claims)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to cast claims"})
			return
		}

		var request jobsuc.CreateJobRequest

		// Bind the incoming JSON to the request struct
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": errors.Wrap(err, "unable to parse incoming request").Error()})
			return
		}

		// Create job using the use case
		response, err := useCase.CreateJob(c.Request.Context(), userClaims, request)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// Respond with the created job
		c.JSON(201, response)
	}
}

func GetJobsHandler(useCase GetJobsUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		Jobs, err := useCase.GetJobs(c.Request.Context()) // Capture potential error
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"total":  len(Jobs),
			"result": Jobs,
		})
	}
}

func GetActiveJobsHandler(useCase GetActiveJobsUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, exists := c.Get("userClaims")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "no claims found"})
		}

		userClaims, ok := claims.(middlewares.Claims)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to cast claims"})
			return
		}

		Jobs, err := useCase.GetActiveJobs(c.Request.Context(), userClaims) // Capture potential error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"total":  len(Jobs),
			"result": Jobs,
		})
	}
}
