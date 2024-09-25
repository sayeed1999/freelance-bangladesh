package middlewares

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/requestid"
	"github.com/sayeed1999/freelance-bangladesh/shared/enums"
)

func InitFiberMiddlewares(app *fiber.App,
	initPublicRoutes func(app *fiber.App),
	initProtectedRoutes func(app *fiber.App)) {
	app.Use(requestid.New())
	app.Use(logger.New())

	app.Use(func(c fiber.Ctx) error {
		// get the request id that was added by the requestid middleware
		var requestid = c.Locals("requestid")

		// create a new context and add the requestid to it
		var ctx = context.WithValue(context.Background(), enums.ContextKeyRequestId, requestid)
		c.SetUserContext((ctx))

		return c.Next()
	})

	// routes that don't require a JWT token
	initPublicRoutes(app)

	// routes that require authentication/authorization
	initProtectedRoutes(app)

	log.Println("fiber middlewares initialized")
}
