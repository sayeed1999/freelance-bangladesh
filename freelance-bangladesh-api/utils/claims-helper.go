package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/sayeed1999/freelance-bangladesh/api/middlewares"
)

func GetUserClaims(c *gin.Context) (middlewares.Claims, error) {
	claims, exists := c.Get("userClaims")
	if !exists {
		return middlewares.Claims{}, errors.New("no claims found")
	}

	userClaims, ok := claims.(middlewares.Claims)
	if !ok {
		return middlewares.Claims{}, errors.New("failed to cast claims")
	}

	return userClaims, nil
}
