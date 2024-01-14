package tests

import (
	"fmt"
	"github.com/ShindeSatish/bookstore/app"
	"github.com/ShindeSatish/bookstore/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http/httptest"
	"os"
	"testing"
)

var (
	testServer *httptest.Server
	db         *gorm.DB
)
var dsn = "mysql_user:mysql_password@tcp(127.0.0.1:3308)/test_bookstore?charset=utf8mb4&parseTime=True&loc=Local"

func TestMain(m *testing.M) {
	config.LoadEnv()

	// Set up the test database
	db = app.InitDatabase()

	// Run migrations
	app.MigrateAndSeedDatabase(db)

	// Set up your application routes
	// Initialize Gin
	r := gin.Default()

	// Set up repositories, services, and handlers
	handlers := app.InitAllComponents(db)

	// Set up routes
	app.SetupRoutes(r, handlers)

	// Create a test HTTP server
	testServer = httptest.NewServer(r)

	go func() {
		err := r.Run(":8081")
		if err != nil {
			panic(err)
		}
	}()
	// Run tests
	exitVal := m.Run()

	// Clean up after tests
	testServer.Close()
	tearDownTestDatabase(db) // Implement this function
	os.Exit(exitVal)
}

func tearDownTestDatabase(d *gorm.DB) {
	// Remove all tables from the test database
	fmt.Println("Removing all tables from the test database...")
}
