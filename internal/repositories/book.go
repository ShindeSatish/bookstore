package repositories

import (
	"github.com/ShindeSatish/bookstore/internal/models"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) GetAllBooks() ([]models.Book, error) {
	var books []models.Book
	err := r.db.Find(&books).Error
	return books, err
}

// GetBookPrices fetches the prices of books with the given IDs
func (r *BookRepository) GetBookPrices(bookIDs []uint) (map[uint]float64, error) {
	var books []models.Book
	result := make(map[uint]float64)

	// Fetch books with the provided IDs
	if err := r.db.Where("id IN ?", bookIDs).Find(&books).Error; err != nil {
		return nil, err
	}

	// Map book IDs to their prices
	for _, book := range books {
		result[book.ID] = book.Price
	}

	return result, nil
}
