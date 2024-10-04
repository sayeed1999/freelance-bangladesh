package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sayeed1999/freelance-bangladesh/api/middlewares"
	"github.com/sayeed1999/freelance-bangladesh/domain/entities"
	admindashboarduc "github.com/sayeed1999/freelance-bangladesh/use_cases/admin_dashboard_uc"
)

type GetClientsUseCase interface {
	Handler(ctx context.Context) ([]entities.Client, error)
}

type GetTalentsUseCase interface {
	Handler(ctx context.Context) ([]entities.Talent, error)
}

type UpdateClientUseCase interface {
	Handler(ctx context.Context, claims middlewares.Claims, command admindashboarduc.UpdateClientCommand) error
}

type UpdateTalentUseCase interface {
	Handler(ctx context.Context, claims middlewares.Claims, command admindashboarduc.UpdateTalentCommand) error
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

		var command admindashboarduc.UpdateClientCommand

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

func UpdateTalentHandler(usecase UpdateTalentUseCase) gin.HandlerFunc {
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

		var command admindashboarduc.UpdateTalentCommand

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
