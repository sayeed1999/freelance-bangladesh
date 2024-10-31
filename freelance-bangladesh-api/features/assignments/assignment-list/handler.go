package assignmentlist

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sayeed1999/freelance-bangladesh/api/handlers"
	"github.com/sayeed1999/freelance-bangladesh/api/middlewares"
)

type AssignmentListUseCase interface {
	AssignmentList(ctx context.Context, userClaims middlewares.Claims) ([]AssignmentResponse, error)
}

func AssignmentListHandler(useCase AssignmentListUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		userClaims, err := handlers.ExtractUserClaims(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Call the use case to list assignments for the specified talent
		assignments, err := useCase.AssignmentList(c.Request.Context(), *userClaims)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"total":  len(assignments),
			"result": assignments,
		})
	}
}
