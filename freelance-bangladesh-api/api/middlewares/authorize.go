package middlewares

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sayeed1999/freelance-bangladesh/config"
)

// JWTAuthMiddleware checks the JWT token in the Authorization header and introspects it
func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// The token should be in the format: "Bearer token"
		tokenString := strings.Split(authHeader, "Bearer ")
		if len(tokenString) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format"})
			c.Abort()
			return
		}

		// Call Keycloak to introspect the token
		isActive, err := IntrospectToken(tokenString[1])
		if err != nil || !isActive {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// If token is valid, continue with the request
		c.Next()
	}
}

func IntrospectToken(token string) (bool, error) {
	cfg := config.GetConfig()

	// Prepare the request payload
	data := fmt.Sprintf("client_id=%s&client_secret=%s&token=%s", cfg.Keycloak.RestApi.ClientId, cfg.Keycloak.RestApi.ClientSecret, token)
	addr := fmt.Sprintf("%v:%v", cfg.ListenIP, cfg.ListenPort)
	req, err := http.NewRequest("POST", fmt.Sprintf("http://%v/realms/%v/protocol/openid-connect/token/introspect", addr, cfg.Dashboard.Realm), bytes.NewBuffer([]byte(data)))
	if err != nil {
		return false, err
	}

	// Set headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Make the request
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	// Read and parse the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	// Check if the token is active (Keycloak returns {"active": true/false})
	if strings.Contains(string(body), `"active":true`) {
		return true, nil
	}

	return false, nil
}
