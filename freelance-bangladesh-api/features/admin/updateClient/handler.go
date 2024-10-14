package updateclient

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sayeed1999/freelance-bangladesh/api/middlewares"
)

type UpdateClientUseCase interface {
	Handler(ctx context.Context, claims middlewares.Claims, command UpdateClientCommand) error
}

func UpdateClientHandler(usecase UpdateClientUseCase) gin.HandlerFunc {
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

		var command UpdateClientCommand

		if err := c.ShouldBindJSON(&command); err != nil {
			c.JSON(400, gin.H{"error": errors.Wrap(err, "unable to parse incoming request").Error()})
			return
		}

		err := usecase.Handler(c.Request.Context(), userClaims, command)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(204, gin.H{"message": "account verification success!"})
	}
}
