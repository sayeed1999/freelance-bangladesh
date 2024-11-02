package middlewares

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"slices"
	"strings"
	"time"

	oidc "github.com/coreos/go-oidc"
	"github.com/gin-gonic/gin"
	"github.com/sayeed1999/freelance-bangladesh/config"
	"github.com/sayeed1999/freelance-bangladesh/shared/enums"
)

type Res401Struct struct {
	Status   string `json:"status" example:"FAILED"`
	HTTPCode int    `json:"httpCode" example:"401"`
	Message  string `json:"message" example:"authorisation failed"`
}

type Claims struct {
	Email         string      `json:"email"`
	EmailVerified bool        `json:"email_verified"`
	FirstName     string      `json:"given_name"`
	LastName      string      `json:"family_name"`
	Username      string      `json:"preferred_username"`
	RealmAccess   realmAccess `json:"realm_access,omitempty"`
	JTI           string      `json:"jti,omitempty"`
}

type realmAccess struct {
	Roles []string `json:"roles,omitempty"`
}

// Authorize middleware for JWT role-based access control in Gin Gonic
func Authorize(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		cfg := config.GetConfig()
		issuerURL := fmt.Sprintf("%v/realms/%v", cfg.Keycloak.BaseUrl, cfg.Keycloak.Realm)

		// Extract token from the Authorization header
		rawAccessToken := c.GetHeader("Authorization")
		if rawAccessToken == "" || !strings.HasPrefix(rawAccessToken, "Bearer ") {
			authorizationFailed("please check authorization header", c)
			c.Abort()
			return
		}

		token := strings.TrimPrefix(rawAccessToken, "Bearer ")

		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // TODO:- DON'T USE IN PRODUCTION!
		}

		client := &http.Client{
			Timeout:   time.Duration(30000) * time.Second,
			Transport: tr,
		}

		ctx := oidc.ClientContext(context.Background(), client)

		provider, err := oidc.NewProvider(ctx, issuerURL)
		if err != nil {
			authorizationFailed("failed to get provider: "+err.Error(), c)
			c.Abort()
			return
		}

		oidcConfig := &oidc.Config{
			// Note:- Error with clientId: oidc: expected audience "backend-api" got ["account"]
			ClientID: "account", // cfg.Keycloak.RestApi.ClientId,

		}

		verifier := provider.Verifier(oidcConfig)
		idToken, err := verifier.Verify(ctx, token)
		if err != nil {
			fmt.Println("Failed to verify token", err)
			authorizationFailed("failed to verify token: "+err.Error(), c)
			c.Abort()
			return
		}

		// Extract claims
		var IDTokenClaims Claims
		if err := idToken.Claims(&IDTokenClaims); err != nil {
			authorizationFailed("failed to parse claims: "+err.Error(), c)
			c.Abort()
			return
		}

		// Check user roles against required roles
		userRoles := IDTokenClaims.RealmAccess.Roles

		// LOGIC: if the user dont have any other application role,
		// we provide him 'talent' scope.
		if !slices.Contains(userRoles, string(enums.ROLE_ADMIN)) &&
			!slices.Contains(userRoles, string(enums.ROLE_CLIENT)) {
			userRoles = append(userRoles, string(enums.ROLE_TALENT))
			IDTokenClaims.RealmAccess.Roles = userRoles
		}

		if !hasRequiredRole(userRoles, roles) {
			// Authorization failed if no roles matched
			authorizationFailed("user not allowed to access this API", c)
			c.Abort()
			return
		}

		c.Set("userClaims", IDTokenClaims)

		c.Next()
	}
}

// Helper function to check if the user has one of the required roles
func hasRequiredRole(userRoles, requiredRoles []string) bool {
	if len(requiredRoles) == 0 {
		return true
	}

	/// User is a talent if:
	/// - user is not an admin user
	/// - user is not a client
	/// - if user has no roles at all!

	for _, role := range requiredRoles {
		for _, userRole := range userRoles {
			if userRole == role {
				return true
			}
		}
	}
	return false
}

// Respond with 401 Unauthorized and message
func authorizationFailed(message string, c *gin.Context) {
	c.JSON(http.StatusUnauthorized, Res401Struct{
		Status:   "FAILED",
		HTTPCode: http.StatusUnauthorized,
		Message:  message,
	})
}

func PrintIDTokenInDebugLevel(idToken *oidc.IDToken) {
	// Marshal the token's claims into a JSON string
	claims := map[string]interface{}{}
	if err := idToken.Claims(&claims); err != nil {
		return
	}

	// Convert the claims to a JSON string
	claimsJSON, err := json.MarshalIndent(claims, "", "  ")
	if err != nil {
		return
	}

	// Log the token's claims
	log.Println("ID Token claims:", string(claimsJSON))
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
