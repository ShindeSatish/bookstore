package dto

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type RegisterUserRequest struct {
	Email     string `json:"email" validate:"required,email" example:"shindesatishsss@gmail.com"`
	Password  string `json:"password" validate:"required,min=8" example:"StrongPassword"`
	FirstName string `json:"first_name" validate:"required,alpha" example:"Satish"`
	LastName  string `json:"last_name" validate:"required,alpha" example:"Shinde"`
	Phone     string `json:"phone" example:"1234567890"`
}

// Use this function to validate the RegisterUserRequest struct

func (r *RegisterUserRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)

}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AuthenticateUserResponse struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Token  string `json:"token"`
}

type ServiceResponse struct {
	Code    int
	Message string
	Data    interface{}
	Error   interface{}
}

type NewOrderRequest struct {
	AdditionalCharges float64            `json:"additional_charges"`
	DiscountAmount    float64            `json:"discount_amount"`
	Items             []OrderItemRequest `json:"items"`
}

type OrderItemRequest struct {
	BookID   uint `json:"book_id"`
	Quantity int  `json:"quantity"`
}
type OrderResponse struct {
	ID                uint                  `json:"id"`
	UserID            uint                  `json:"user_id"`
	TotalPrice        float64               `json:"total_price"`
	AdditionalCharges float64               `json:"additional_charges"`
	DiscountAmount    float64               `json:"discount_amount"`
	CreatedAt         time.Time             `json:"created_at"`
	UpdatedAt         time.Time             `json:"updated_at"`
	OrderDetails      []OrderDetailResponse `json:"order_details"`
}

type OrderDetailResponse struct {
	BookID     uint `json:"book_id"`
	Quantity   int  `json:"quantity"`
	Price      float64
	TotalPrice float64
}

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

type BookResponse struct {
	ID     uint    `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}
