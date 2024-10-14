package gettalents

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sayeed1999/freelance-bangladesh/domain/entities"
)

type GetTalentsUseCase interface {
	Handler(ctx context.Context) ([]entities.Talent, error)
}

func GetTalentsHandler(useCase GetTalentsUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {

		talents, err := useCase.Handler(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"total":  len(talents),
			"result": talents,
		})
	}
}
