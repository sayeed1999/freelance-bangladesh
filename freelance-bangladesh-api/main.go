package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sayeed1999/freelance-bangladesh/api/routes"
	"github.com/sayeed1999/freelance-bangladesh/config"
	"github.com/sayeed1999/freelance-bangladesh/database"
)

func main() {
	cfg := config.GetConfig()

	// Connect to database instance
	database.Connect()

	// Initialize Gin engine
	app := gin.Default()

	// Initialize routes
	routes.InitRoutes(app)

	addr := fmt.Sprintf("%v:%v", cfg.ListenIP, cfg.ListenPort)
	log.Printf("%v api will listen on %v", cfg.Dashboard.Title, addr)

	err := app.Run(addr)
	log.Fatal(err)
}
