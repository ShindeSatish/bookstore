package helpers

import (
	"github.com/ShindeSatish/bookstore/internal/dto"
	"github.com/ShindeSatish/bookstore/internal/models"
)

func FromOrderModelToOrderResponse(order models.Order) dto.OrderResponse {
	var orderDetails []dto.OrderDetailResponse

	for _, orderDetail := range order.OrderDetails {
		orderDetails = append(orderDetails, dto.OrderDetailResponse{
			BookID:     orderDetail.BookID,
			Quantity:   orderDetail.Quantity,
			Price:      orderDetail.Price,
			TotalPrice: orderDetail.TotalPrice,
		})
	}
	return dto.OrderResponse{
		ID:                order.ID,
		UserID:            order.UserID,
		TotalPrice:        order.TotalPrice,
		CreatedAt:         order.CreatedAt,
		UpdatedAt:         order.UpdatedAt,
		OrderDetails:      orderDetails,
		AdditionalCharges: order.AdditionalCharges,
	}
}
