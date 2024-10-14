package getjobs

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sayeed1999/freelance-bangladesh/api/middlewares"
	"github.com/sayeed1999/freelance-bangladesh/domain/entities"
	"github.com/sayeed1999/freelance-bangladesh/utils"
)

type GetJobsUseCase interface {
	GetJobs(ctx context.Context, userClaims middlewares.Claims) ([]entities.Job, error)
}

func GetJobsHandler(useCase GetJobsUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		userClaims, err := utils.GetUserClaims(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		Jobs, err := useCase.GetJobs(c.Request.Context(), userClaims)
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
