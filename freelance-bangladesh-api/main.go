package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sayeed1999/freelance-bangladesh/api/routes"
	"github.com/sayeed1999/freelance-bangladesh/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Initialize Gin engine
	app := gin.Default()

	// Initialize routes
	routes.InitRoutes(app)

	addr := fmt.Sprintf("%v:%v", cfg.ListenIP, cfg.ListenPort)
	log.Printf("%v api will listen on %v", cfg.Dashboard.Title, addr)

	err = app.Run(addr)
	log.Fatal(err)
}
