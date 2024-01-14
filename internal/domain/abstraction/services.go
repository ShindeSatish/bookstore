package abstraction

import (
	"github.com/ShindeSatish/bookstore/internal/dto"
)

type UserService interface {
	RegisterUser(request *dto.RegisterUserRequest) dto.ServiceResponse
	Authenticate(email, password string) dto.ServiceResponse
}

type BookService interface {
	GetAllBooks() ([]dto.BookResponse, error)
	FetchBookPrices(bookIDs []uint) (map[uint]float64, error)
}

type OrderService interface {
	CreateOrder(userID uint, request dto.NewOrderRequest) dto.ServiceResponse
	GetOrdersByUserID(userID uint) ([]dto.OrderResponse, error)
}

type BookPriceFetcher interface {
	FetchBookPrices(bookIDs []uint) (map[uint]float64, error)
}
