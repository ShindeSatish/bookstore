package app

import (
	_ "github.com/ShindeSatish/bookstore/docs"
	"github.com/ShindeSatish/bookstore/internal/handlers"
	"github.com/ShindeSatish/bookstore/internal/middleware"
	"github.com/ShindeSatish/bookstore/internal/repositories"
	"github.com/ShindeSatish/bookstore/internal/services"
	"github.com/ShindeSatish/bookstore/migrations"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func InitDatabase() *gorm.DB {
	dsn := os.Getenv("MYSQL_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}

func MigrateAndSeedDatabase(db *gorm.DB) {
	if err := migrations.MigrateAndSeed(db); err != nil {
		log.Fatalf("Failed to migrate or seed: %v", err)
	}
}

type Handlers struct {
	UserHandler  *handlers.UserHandler
	BookHandler  *handlers.BookHandler
	OrderHandler *handlers.OrderHandler
}

func InitAllComponents(db *gorm.DB) *Handlers {
	// Initialize User dependency components
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Initialize Book dependency components
	bookRepo := repositories.NewBookRepository(db)
	bookService := services.NewBookService(bookRepo)
	bookHandler := handlers.NewBookHandler(bookService)

	// Initialize Order dependency components

	orderRepo := repositories.NewOrderRepository(db)
	orderService := services.NewOrderService(*orderRepo, bookService)
	orderHandler := handlers.NewOrderHandler(orderService)

	return &Handlers{
		UserHandler:  userHandler,
		BookHandler:  bookHandler,
		OrderHandler: orderHandler,
	}

}

func SetupRoutes(router *gin.Engine, handlers *Handlers) {
	router.POST("/register", handlers.UserHandler.Register)
	router.POST("/login", handlers.UserHandler.Login)
	router.GET("/books", handlers.BookHandler.GetBooks)

	// Create order against login user
	router.POST("/order", middleware.TokenAuthMiddleware(), handlers.OrderHandler.CreateOrder)
	// send all the order of login user
	router.GET("/orders", middleware.TokenAuthMiddleware(), handlers.OrderHandler.GetOrdersByUserID)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
