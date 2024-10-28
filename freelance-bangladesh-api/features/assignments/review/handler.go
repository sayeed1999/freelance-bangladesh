package review

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sayeed1999/freelance-bangladesh/api/handlers"
	"github.com/sayeed1999/freelance-bangladesh/api/middlewares"
)

type ReviewWorkUseCase interface {
	ReviewWork(ctx context.Context, claims middlewares.Claims, request ReviewWorkRequest) (*ReviewResponse, error)
}

func ReviewWorkHandler(useCase ReviewWorkUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		userClaims, err := handlers.ExtractUserClaims(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Bind the JSON request to ReviewWorkRequest
		var request ReviewWorkRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "unable to parse request: " + err.Error()})
			return
		}

		// Call the use case
		response, err := useCase.ReviewWork(c.Request.Context(), *userClaims, request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, response)
	}
}
