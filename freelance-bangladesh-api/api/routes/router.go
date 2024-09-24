package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/sayeed1999/freelance-bangladesh/api/handlers"
	"github.com/sayeed1999/freelance-bangladesh/infrastructure/identity"
	"github.com/sayeed1999/freelance-bangladesh/use_cases/usermgmtuc"
)

func InitPublicRoutes(app *fiber.App) {
	app.Get("/", func(ctx fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to Freelance Bangladesh API v1.0!"))
	})

	group := app.Group("/api/v1")

	identityManager := identity.NewIdentityManager()
	registerUseCase := usermgmtuc.NewRegisterUseCase(identityManager)

	group.Post("users", handlers.RegisterHandler(registerUseCase))
}
