package auth

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/sayeed1999/freelance-bangladesh/api/handlers"
	"github.com/sayeed1999/freelance-bangladesh/api/middlewares"
)

type SyncUserUseCase interface {
	SyncUser(ctx context.Context, claims middlewares.Claims) (*SyncUserResponse, error)
}

func SyncUserHandler(useCase SyncUserUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		userClaims, err := handlers.ExtractUserClaims(c)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
		}

		response, err := useCase.SyncUser(c.Request.Context(), *userClaims)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(201, response)
	}
}
