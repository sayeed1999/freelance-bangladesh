package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sayeed1999/freelance-bangladesh/features/admin"
	"github.com/sayeed1999/freelance-bangladesh/features/assignments"
	"github.com/sayeed1999/freelance-bangladesh/features/auth"
	"github.com/sayeed1999/freelance-bangladesh/features/jobs"
)

// InitRoutes initializes all the routes
func InitRoutes(app *gin.Engine) {

	// Grouping API v1 routes
	apiV1 := app.Group("/api/v1")
	{
		auth.RegisterUserManagementRoutes(apiV1)
		admin.RegisterAdminRoutes(apiV1)
		jobs.RegisterJobRoutes(apiV1)
		assignments.RegisterAssignmentRoutes(apiV1)
	}

	// Homepage route
	app.GET("/", homePage)
}

// homePage handles the root route
func homePage(c *gin.Context) {
	c.String(200, "Welcome to Freelance Bangladesh API v1.0!")
}
