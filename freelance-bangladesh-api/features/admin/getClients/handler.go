package getclients

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sayeed1999/freelance-bangladesh/models"
)

type GetClientsUseCase interface {
	Handler(ctx context.Context) ([]models.Client, error)
}

func GetClientsHandler(useCase GetClientsUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {

		clients, err := useCase.Handler(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"total":  len(clients),
			"result": clients,
		})
	}
}
