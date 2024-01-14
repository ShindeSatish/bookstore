package main

import (
	"github.com/ShindeSatish/bookstore/app"
	"github.com/ShindeSatish/bookstore/config"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Initialize database connection

	db := app.InitDatabase()

	// After initializing the database connection
	app.MigrateAndSeedDatabase(db)

	// Initialize Gin
	r := gin.Default()

	// Set up repositories, services, and handlers
	handlers := app.InitAllComponents(db)

	// Set up routes
	app.SetupRoutes(r, handlers)

	// Start server
	port := config.GetPort()
	r.Run(port)
}
