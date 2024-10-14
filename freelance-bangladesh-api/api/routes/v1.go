package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sayeed1999/freelance-bangladesh/api/handlers"
	"github.com/sayeed1999/freelance-bangladesh/api/middlewares"
	bidjob "github.com/sayeed1999/freelance-bangladesh/features/jobs/bidJob"
	createjob "github.com/sayeed1999/freelance-bangladesh/features/jobs/createJob"
	getjobs "github.com/sayeed1999/freelance-bangladesh/features/jobs/getJobs"
	"github.com/sayeed1999/freelance-bangladesh/infrastructure/identity"
	"github.com/sayeed1999/freelance-bangladesh/shared/enums"
	admindashboarduc "github.com/sayeed1999/freelance-bangladesh/use_cases/admin_dashboard_uc"
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
	updateClientUseCase := admindashboarduc.NewUpdateClientUseCase()
	updateTalentUseCase := admindashboarduc.NewUpdateTalentUseCase()

	adminRoutes := rg.Group("/admin-dashboard")
	{
		adminRoutes.Use(middlewares.Authorize(string((enums.ROLE_ADMIN))))

		adminRoutes.GET("/clients",
			handlers.GetClientsHandler(getClientsUseCase))

		adminRoutes.GET("/talents",
			handlers.GetTalentsHandler(getTalentsUseCase))

		adminRoutes.POST("/clients",
			handlers.UpdateClientHandler(updateClientUseCase))

		adminRoutes.POST("/talents",
			handlers.UpdateTalentHandler(updateTalentUseCase))

	}

	return adminRoutes
}

func RegisterJobRoutes(rg *gin.RouterGroup) *gin.RouterGroup {
	createJobUseCase := createjob.NewCreateJobUseCase()
	getJobsUseCase := getjobs.NewGetJobsUseCase()
	bidJobUseCase := bidjob.NewBidOnJobUseCase()

	jobs := rg.Group("/jobs")
	{
		jobs.POST(
			"",
			middlewares.Authorize(string(enums.ROLE_CLIENT)),
			createjob.CreateJobHandler(createJobUseCase),
		)

		jobs.GET(
			"",
			middlewares.Authorize(
				string(enums.ROLE_ADMIN),
				string(enums.ROLE_CLIENT),
				string(enums.ROLE_TALENT)),
			getjobs.GetJobsHandler(getJobsUseCase),
		)

		jobs.POST(
			"/:jobid/bids",
			middlewares.Authorize(string(enums.ROLE_TALENT)),
			bidjob.BidOnJobHandler(bidJobUseCase),
		)
	}

	return jobs
}
