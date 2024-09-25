package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/sayeed1999/freelance-bangladesh/api/handlers"
	"github.com/sayeed1999/freelance-bangladesh/infrastructure/identity"
	jobsuc "github.com/sayeed1999/freelance-bangladesh/use_cases/jobs_uc"
	"github.com/sayeed1999/freelance-bangladesh/use_cases/usermgmtuc"
)

func InitPublicRoutes(app *fiber.App) {
	app.Get("/", func(ctx fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to Freelance Bangladesh API v1.0!"))
	})

	group := app.Group("/api/v1")

	identityManager := identity.NewIdentityManager()
	registerUseCase := usermgmtuc.NewRegisterUseCase(identityManager)

	group.Post("/users", handlers.RegisterHandler(registerUseCase))
}

func InitProtectedRoutes(app *fiber.App) {
	group := app.Group("/api/v1")

	createJobUseCase := jobsuc.NewCreateJobUseCase()
	getJobsUseCase := jobsuc.NewGetJobsUseCase()

	group.Post("/jobs", handlers.CreateJobHandler(createJobUseCase))
	group.Get("/jobs", handlers.GetJobsHandler(getJobsUseCase))
}
