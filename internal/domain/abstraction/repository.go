package abstraction

import (
	"github.com/ShindeSatish/bookstore/internal/models"
	"github.com/gin-gonic/gin"
)

type UserRepository interface {
	CreateUser(user models.User) error
	GetUserByEmail(email string) (models.User, error)
}

type BookRepository interface {
	GetBooks(ctx *gin.Context)
	GetBookPrices(bookIDs []uint) (map[uint]float64, error)
}

type OrderRepository interface {
	CreateOrder(order *models.Order) error
	GetOrdersByUserID(userID uint) ([]models.Order, error)
}
