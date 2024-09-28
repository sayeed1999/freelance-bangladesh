package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sayeed1999/freelance-bangladesh/api/handlers"
	"github.com/sayeed1999/freelance-bangladesh/api/middlewares"
	"github.com/sayeed1999/freelance-bangladesh/infrastructure/identity"
	"github.com/sayeed1999/freelance-bangladesh/shared/enums"
	jobsuc "github.com/sayeed1999/freelance-bangladesh/use_cases/jobs_uc"
	"github.com/sayeed1999/freelance-bangladesh/use_cases/usermgmtuc"
)

// InitRoutes initializes all the routes
func InitRoutes(app *gin.Engine) {
	identityManager := identity.NewIdentityManager()
	registerUseCase := usermgmtuc.NewRegisterUseCase(identityManager)
	createJobUseCase := jobsuc.NewCreateJobUseCase()
	getJobsUseCase := jobsuc.NewGetJobsUseCase()

	// Grouping API v1 routes
	apiV1 := app.Group("/api/v1")
	{
		// User management routes
		users := apiV1.Group("/users")
		{
			users.POST("", handlers.RegisterHandler(registerUseCase))
		}

		// Jobs routes
		jobs := apiV1.Group("/jobs")
		{
			jobs.POST(
				"",
				middlewares.Authorize(string(enums.ROLE_ADMIN), string(enums.ROLE_CLIENT)),
				handlers.CreateJobHandler(createJobUseCase),
			)
			jobs.GET(
				"",
				middlewares.Authorize(string(enums.ROLE_ADMIN), string(enums.ROLE_TALENT)),
				handlers.GetJobsHandler(getJobsUseCase),
			)
		}
	}

	// Homepage route
	app.GET("/", homePage)
}

// homePage handles the root route
func homePage(c *gin.Context) {
	c.String(200, "Welcome to Freelance Bangladesh API v1.0!")
}
