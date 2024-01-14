package models

import (
	"time"
)

type Order struct {
	ID                uint `gorm:"primaryKey"`
	UserID            uint // Foreign key for User
	TotalPrice        float64
	AdditionalCharges float64
	DiscountAmount    float64
	CreatedAt         time.Time
	UpdatedAt         time.Time
	OrderDetails      []OrderDetail `gorm:"foreignKey:OrderID"`
}

type OrderDetail struct {
	ID         uint `gorm:"primaryKey"`
	OrderID    uint // Foreign key for Order
	BookID     uint // Foreign key for Book
	Price      float64
	Quantity   int
	TotalPrice float64
}
