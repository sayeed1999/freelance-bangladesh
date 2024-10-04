package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sayeed1999/freelance-bangladesh/api/handlers"
	"github.com/sayeed1999/freelance-bangladesh/api/middlewares"
	"github.com/sayeed1999/freelance-bangladesh/infrastructure/identity"
	"github.com/sayeed1999/freelance-bangladesh/shared/enums"
	admindashboarduc "github.com/sayeed1999/freelance-bangladesh/use_cases/admin_dashboard_uc"
	jobsuc "github.com/sayeed1999/freelance-bangladesh/use_cases/jobs_uc"
	"github.com/sayeed1999/freelance-bangladesh/use_cases/usermgmtuc"
)

// InitRoutes initializes all the routes
func InitRoutes(app *gin.Engine) {

	// Grouping API v1 routes
	apiV1 := app.Group("/api/v1")
	{
		RegisterUserManagementRoutes(apiV1)
		RegisterAdminRoutes(apiV1)
		RegisterJobRoutes(apiV1)
	}

	// Homepage route
	app.GET("/", homePage)
}

// homePage handles the root route
func homePage(c *gin.Context) {
	c.String(200, "Welcome to Freelance Bangladesh API v1.0!")
}

func RegisterUserManagementRoutes(rg *gin.RouterGroup) *gin.RouterGroup {
	identityManager := identity.NewIdentityManager()
	registerUseCase := usermgmtuc.NewRegisterUseCase(identityManager)

	users := rg.Group("/users")
	{
		// N.B: client sigup considered admin route!
		users.POST("/client-signup",
			middlewares.Authorize(string(enums.ROLE_ADMIN)),
			handlers.RegisterClientHandler(registerUseCase))

		users.POST("/talent-signup",
			handlers.RegisterTalentHandler(registerUseCase))
	}

	return users
}

func RegisterAdminRoutes(rg *gin.RouterGroup) *gin.RouterGroup {
	getClientsUseCase := admindashboarduc.NewGetClientsUseCase()
	getTalentsUseCase := admindashboarduc.NewGetTalentsUseCase()
	verifyClientUseCase := admindashboarduc.NewVerifyClientUseCase()
	verifyTalentUseCase := admindashboarduc.NewVerifyTalentUseCase()

	adminRoutes := rg.Group("/admin-dashboard")
	{
		adminRoutes.Use(middlewares.Authorize(string((enums.ROLE_ADMIN))))

		adminRoutes.GET("/clients",
			handlers.GetClientsHandler(getClientsUseCase))

		adminRoutes.GET("/talents",
			handlers.GetTalentsHandler(getTalentsUseCase))

		adminRoutes.POST("/clients/verify",
			handlers.VerifyClientHandler(verifyClientUseCase))

		adminRoutes.POST("/talents/verify",
			handlers.VerifyTalentHandler(verifyTalentUseCase))

	}

	return adminRoutes
}

func RegisterJobRoutes(rg *gin.RouterGroup) *gin.RouterGroup {
	createJobUseCase := jobsuc.NewCreateJobUseCase()
	getJobsUseCase := jobsuc.NewGetJobsUseCase()

	jobs := rg.Group("/jobs")
	{
		jobs.POST(
			"",
			middlewares.Authorize(string(enums.ROLE_CLIENT)),
			handlers.CreateJobHandler(createJobUseCase),
		)

		jobs.GET(
			"",
			middlewares.Authorize(
				string(enums.ROLE_ADMIN),
				string(enums.ROLE_CLIENT),
				string(enums.ROLE_TALENT)),
			handlers.GetJobsHandler(getJobsUseCase),
		)
	}

	return jobs
}
