package services

import (
	"github.com/ShindeSatish/bookstore/internal/domain/abstraction"
	"github.com/ShindeSatish/bookstore/internal/dto"
	"github.com/ShindeSatish/bookstore/internal/helpers"
	"github.com/ShindeSatish/bookstore/internal/models"
	"github.com/ShindeSatish/bookstore/internal/repositories"
	"net/http"
)

type orderService struct {
	orderRepo        repositories.OrderRepository
	BookPriceFetcher abstraction.BookPriceFetcher
}

func NewOrderService(orderRepo repositories.OrderRepository, bookPriceFetcher abstraction.BookPriceFetcher) abstraction.OrderService {
	return &orderService{
		orderRepo:        orderRepo,
		BookPriceFetcher: bookPriceFetcher,
	}
}

// Assume we have a function to calculate the total price of each item
func calculateItemTotalPrice(price float64, quantity int) float64 {
	return price * float64(quantity)
}

func (s *orderService) CreateOrder(userID uint, req dto.NewOrderRequest) dto.ServiceResponse {
	var orderItems []models.OrderDetail
	var totalOrderPrice float64 = 0

	// Process each item in the order
	var bookIds []uint
	for _, item := range req.Items {
		bookIds = append(bookIds, item.BookID)
	}

	// Fetch the prices of the books
	bookPrices, err := s.BookPriceFetcher.FetchBookPrices(bookIds)
	if err != nil {
		return dto.ServiceResponse{Code: http.StatusInternalServerError, Message: "Failed to fetch book prices", Error: err}
	}

	for _, item := range req.Items {
		price := bookPrices[item.BookID]
		itemTotalPrice := calculateItemTotalPrice(price, item.Quantity)
		totalOrderPrice += itemTotalPrice

		orderItems = append(orderItems, models.OrderDetail{
			BookID:     item.BookID,
			Quantity:   item.Quantity,
			Price:      price,
			TotalPrice: itemTotalPrice,
		})
	}

	// Calculate final order price with additional charges and discount
	finalOrderPrice := totalOrderPrice + req.AdditionalCharges - req.DiscountAmount

	// Create the order model
	order := models.Order{
		UserID:            userID,
		TotalPrice:        finalOrderPrice,
		AdditionalCharges: req.AdditionalCharges,
		DiscountAmount:    req.DiscountAmount,
		OrderDetails:      orderItems,
	}

	// Persist the order using the order repository
	err = s.orderRepo.CreateOrder(&order)
	if err != nil {
		return dto.ServiceResponse{Code: http.StatusInternalServerError, Message: "Failed to create order", Error: err}
	}

	orderResponse := helpers.FromOrderModelToOrderResponse(order)
	return dto.ServiceResponse{
		Code:    http.StatusOK,
		Message: "Order created successfully",
		Data:    orderResponse,
	}
}

func (s *orderService) GetOrdersByUserID(userID uint) ([]dto.OrderResponse, error) {
	orders, err := s.orderRepo.GetOrdersByUserID(userID)
	if err != nil {
		return []dto.OrderResponse{}, err
	}
	var orderResponses []dto.OrderResponse
	for _, order := range orders {
		orderResponse := helpers.FromOrderModelToOrderResponse(order)
		orderResponses = append(orderResponses, orderResponse)
	}

	return orderResponses, nil
}
