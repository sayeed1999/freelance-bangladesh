package auth

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sayeed1999/freelance-bangladesh/shared/enums"
)

type RegisterUseCase interface {
	Register(context.Context, RegisterRequest) (*RegisterResponse, error)
}

func RegisterClientHandler(useCase RegisterUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx = c.Request.Context()
		var request RegisterRequest

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": errors.Wrap(err, "unable to parse incoming request").Error()})
			return
		}

		// decide role from endpoint, not user!!
		request.Role = string(enums.ROLE_CLIENT)

		response, err := useCase.Register(ctx, request)
		if err != nil {
			c.JSON(400, gin.H{"error": errors.Wrap(err, "unable to register user").Error()})
			return
		}

		c.JSON(201, response)
	}
}

func RegisterTalentHandler(useCase RegisterUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx = c.Request.Context()
		var request RegisterRequest

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": errors.Wrap(err, "unable to parse incoming request").Error()})
			return
		}

		// decide role from endpoint, not user!!
		request.Role = string(enums.ROLE_TALENT)

		response, err := useCase.Register(ctx, request)
		if err != nil {
			c.JSON(400, gin.H{"error": errors.Wrap(err, "unable to register user").Error()})
			return
		}

		c.JSON(201, response)
	}
}
