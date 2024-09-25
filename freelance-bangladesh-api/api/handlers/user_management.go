package handlers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sayeed1999/freelance-bangladesh/use_cases/usermgmtuc"
)

type RegisterUseCase interface {
	Register(context.Context, usermgmtuc.RegisterRequest) (*usermgmtuc.RegisterResponse, error)
}

func RegisterHandler(useCase RegisterUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx = c.Request.Context()
		var request usermgmtuc.RegisterRequest

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": errors.Wrap(err, "unable to parse incoming request").Error()})
			return
		}

		response, err := useCase.Register(ctx, request)
		if err != nil {
			c.JSON(400, gin.H{"error": errors.Wrap(err, "unable to register user").Error()})
			return
		}

		c.JSON(201, response)
	}
}
