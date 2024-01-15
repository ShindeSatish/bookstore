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
		{Title: "1984", Author: "George Orwell", Price: 9.99},
		{Title: "To Kill a Mockingbird", Author: "Harper Lee", Price: 7.99},
		{Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Price: 8.99},
		{Title: "Pride and Prejudice", Author: "Jane Austen", Price: 6.99},
		{Title: "Brave New World", Author: "Aldous Huxley", Price: 9.99},
		{Title: "The Catcher in the Rye", Author: "J.D. Salinger", Price: 8.99},
		{Title: "Animal Farm", Author: "George Orwell", Price: 7.99},
		{Title: "The Hobbit", Author: "J.R.R. Tolkien", Price: 10.99},
		{Title: "The Alchemist", Author: "Paulo Coelho", Price: 9.99},
		{Title: "Sapiens: A Brief History of Humankind", Author: "Yuval Noah Harari", Price: 12.99},
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
