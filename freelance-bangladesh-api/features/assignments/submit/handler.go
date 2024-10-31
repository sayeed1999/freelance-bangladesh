package submit

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sayeed1999/freelance-bangladesh/api/handlers"
	"github.com/sayeed1999/freelance-bangladesh/api/middlewares"
)

type SubmitWorkUseCase interface {
	SubmitWork(ctx context.Context, claims middlewares.Claims, request SubmitWorkRequest) (*SubmitResponse, error)
}

func SubmitWorkHandler(useCase SubmitWorkUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		userClaims, err := handlers.ExtractUserClaims(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		assignmentID := c.Param("assignmentid")
		if assignmentID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "assignment id is required"})
			return
		}

		// Bind the JSON request to SubmitWorkRequest
		var request SubmitWorkRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "unable to parse request: " + err.Error()})
			return
		}

		request.AssignmentID = assignmentID // id on the route param gets higher priority

		// Call the use case
		response, err := useCase.SubmitWork(c.Request.Context(), *userClaims, request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, response)
	}
}
