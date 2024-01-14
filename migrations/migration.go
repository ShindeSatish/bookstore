package migrations

import (
	"github.com/ShindeSatish/bookstore/internal/models"
	"gorm.io/gorm"
)

func MigrateAndSeed(db *gorm.DB) error {
	// Migrate the schema
	err := db.AutoMigrate(&models.User{}, &models.Book{}, &models.Order{}, &models.OrderDetail{})
	if err != nil {
		return err
	}

	// Seed data
	return seedData(db)
}

func seedData(db *gorm.DB) error {
	// Example: Seed books
	books := []models.Book{
		{Title: "The Go Programming Language", Author: "Alan A. A. Donovan", Price: 44.99},
		{Title: "GORM Guide", Author: "Jinzhu", Price: 35.99},
		// Add more books as needed
	}

	for _, book := range books {
		// Check if the book already exists to avoid duplicate seeding
		if err := db.Where(models.Book{Title: book.Title}).FirstOrCreate(&book).Error; err != nil {
			return err
		}
	}

	// Add more seeding for other tables if necessary

	return nil
}
