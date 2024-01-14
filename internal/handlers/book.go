package handlers

import (
	"github.com/ShindeSatish/bookstore/internal/services"
	"github.com/ShindeSatish/bookstore/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BookHandler struct {
	service *services.BookService
}

func NewBookHandler(service *services.BookService) *BookHandler {
	return &BookHandler{service: service}
}

// @summary Get all books
// @description Get all books
// @tags books
// @accept json
// @produce json
// @success 200 {object} dto.APIResponse
// @router /books [get]
func (h *BookHandler) GetBooks(c *gin.Context) {
	// Fetch all books
	books, err := h.service.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewErrorResponse("Something went wrong!", err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.NewSuccessResponse("Books retrieved successfully", books))
}
