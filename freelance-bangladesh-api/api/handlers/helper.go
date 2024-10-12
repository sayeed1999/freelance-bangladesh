package handlers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/sayeed1999/freelance-bangladesh/api/middlewares"
)

func ExtractUserClaims(c *gin.Context) (*middlewares.Claims, error) {
	claims, exists := c.Get("userClaims")
	if !exists {
		return nil, errors.New("no claims found")
	}

	userClaims, ok := claims.(middlewares.Claims)
	if !ok {
		return nil, errors.New("failed to cast claims")
	}

	return &userClaims, nil
}
