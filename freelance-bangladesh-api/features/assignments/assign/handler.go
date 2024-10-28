package assign

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sayeed1999/freelance-bangladesh/api/handlers"
	"github.com/sayeed1999/freelance-bangladesh/api/middlewares"
)

type AssignTalentUseCase interface {
	AssignTalent(ctx context.Context, claims middlewares.Claims, request AssignRequest) (*AssignResponse, error)
}

func AssignTalentHandler(useCase AssignTalentUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		userClaims, err := handlers.ExtractUserClaims(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Bind the JSON request to AssignRequest
		var request AssignRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "unable to parse request: " + err.Error()})
			return
		}

		// Call the use case
		response, err := useCase.AssignTalent(c.Request.Context(), *userClaims, request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, response)
	}
}
