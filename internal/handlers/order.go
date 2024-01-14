package handlers

import (
	"github.com/ShindeSatish/bookstore/internal/domain/abstraction"
	"github.com/ShindeSatish/bookstore/internal/dto"
	"github.com/ShindeSatish/bookstore/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderHandler struct {
	service abstraction.OrderService
}

func NewOrderHandler(service abstraction.OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

// @summary Create a new order
// @description Create a new order for a login user (this API requires a valid Authentication token)
// @tags orders
// @accept json
// @produce json
// @Param Authorization header string true "Provide Authorization token you get after login"
// @param order body dto.NewOrderRequest true "New Order"
// @success 200 {object} dto.APIResponse
// @router /order [post]
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var orderRequest dto.NewOrderRequest
	if err := c.BindJSON(&orderRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId, ok := GetUserIdFromCtx(c)
	if !ok {
		// handle the error: the assertion failed
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get userID from Authorization"})
		return
	}

	// Call the service layer to create the order
	serviceResponse := h.service.CreateOrder(userId, orderRequest)
	if serviceResponse.Code != http.StatusOK {
		c.JSON(http.StatusInternalServerError, utils.NewErrorResponse(serviceResponse.Message, serviceResponse.Error))
		return
	}

	c.JSON(http.StatusOK, utils.NewSuccessResponse("Order created successfully", serviceResponse.Data))
}

// @summary Get orders by user ID
// @description Get orders by user ID (this API requires a valid Authentication token)
// @tags orders
// @Param Authorization header string true "Provide Authorization token you get after login"
// @success 200 {object} dto.APIResponse
// @router /orders [get]
func (h *OrderHandler) GetOrdersByUserID(c *gin.Context) {
	userId, ok := GetUserIdFromCtx(c)
	if !ok {
		// handle the error: the assertion failed
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get userID from Authorization"})
		return
	}

	// Call the service layer to get the orders
	orders, err := h.service.GetOrdersByUserID(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewErrorResponse("Something went wrong!", err.Error()))
		return
	}

	// if not orders present for this user return success response with no orders
	if len(orders) == 0 {
		c.JSON(http.StatusOK, utils.NewSuccessResponse("No orders found for this user", orders))
		return
	}

	c.JSON(http.StatusOK, utils.NewSuccessResponse("Orders retrieved successfully", orders))
}

func GetUserIdFromCtx(c *gin.Context) (uint, bool) {
	userID, exists := c.Get("userID")
	if !exists {
		return 0, false
	}
	userIDInt, ok := userID.(uint)
	if !ok {
		return 0, false
	}
	return userIDInt, true
}
