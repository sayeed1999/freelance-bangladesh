package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/sayeed1999/freelance-bangladesh/api/middlewares"
	"github.com/sayeed1999/freelance-bangladesh/api/routes"
	"github.com/sayeed1999/freelance-bangladesh/config"
	"github.com/spf13/viper"
)

func main() {
	config.LoadConfig()
	appTitle := viper.GetString("Dashboard.Title")

	app := fiber.New(fiber.Config{
		AppName:      appTitle,
		ServerHeader: "Fiber",
	})

	middlewares.InitFiberMiddlewares(app, routes.InitPublicRoutes, routes.InitProtectedRoutes)

	var listenIp = viper.GetString("ListenIP")
	var listenPort = viper.GetString("ListenPort")

	log.Printf("api will listen on %v:%v", listenIp, listenPort)

	err := app.Listen(fmt.Sprintf("%v:%v", listenIp, listenPort))
	log.Fatal(err)
}
