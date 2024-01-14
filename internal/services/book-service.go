package services

import (
	"github.com/ShindeSatish/bookstore/internal/dto"
	"github.com/ShindeSatish/bookstore/internal/repositories"
)

type BookService struct {
	repo *repositories.BookRepository
}

func NewBookService(repo *repositories.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) GetAllBooks() ([]dto.BookResponse, error) {
	var books []dto.BookResponse

	response, err := s.repo.GetAllBooks()
	if err != nil {
		return books, err
	}

	for _, book := range response {
		books = append(books, dto.BookResponse{
			ID:     book.ID,
			Title:  book.Title,
			Author: book.Author,
			Price:  book.Price,
		})
	}
	return books, nil
}

func (s *BookService) FetchBookPrices(bookIDs []uint) (map[uint]float64, error) {
	return s.repo.GetBookPrices(bookIDs)
}
