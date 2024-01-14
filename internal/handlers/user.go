package handlers

import (
	"github.com/ShindeSatish/bookstore/internal/domain/abstraction"
	"github.com/ShindeSatish/bookstore/internal/dto"
	"github.com/ShindeSatish/bookstore/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	service abstraction.UserService
}

func NewUserHandler(service abstraction.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// @Summary Register new user
// @Description Register a new user with email and password
// @Tags users
// @Accept json
// @Produce json
// @Param user body dto.RegisterUserRequest true "Register User"
// @Success 200 {object} dto.APIResponse
// @Router /register [post]
func (h *UserHandler) Register(c *gin.Context) {
	request := &dto.RegisterUserRequest{}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewErrorResponse("Invalid request body", err.Error()))
		return
	}

	serviceResponse := h.service.RegisterUser(request)
	if serviceResponse.Code != http.StatusOK {
		c.JSON(serviceResponse.Code, utils.NewErrorResponse(serviceResponse.Message, serviceResponse.Error))
		return
	}

	c.JSON(http.StatusOK, utils.NewSuccessResponse("User registered successfully", nil))
}

// @Summary Login user
// @Description Login a user with email and password
// @Tags users
// @Accept json
// @Produce json
// @Param user body dto.LoginRequest true "Login User"
// @Success 200 {object} dto.APIResponse
// @Router /login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	serviceResponse := h.service.Authenticate(req.Email, req.Password)
	if serviceResponse.Code != http.StatusOK {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, utils.NewSuccessResponse("User logged in successfully", serviceResponse.Data))
}
