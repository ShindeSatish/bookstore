package services

import (
	"github.com/ShindeSatish/bookstore/internal/domain/abstraction"
	"github.com/ShindeSatish/bookstore/internal/dto"
	"github.com/ShindeSatish/bookstore/internal/helpers"
	"github.com/ShindeSatish/bookstore/internal/utils"
	"net/http"
)

type userService struct {
	repo abstraction.UserRepository
}

func NewUserService(repo abstraction.UserRepository) abstraction.UserService {
	return &userService{repo: repo}
}

func (s *userService) RegisterUser(request *dto.RegisterUserRequest) dto.ServiceResponse {
	// Validate the request before move ahead
	err := request.Validate()
	if err != nil {
		return dto.ServiceResponse{Code: http.StatusBadRequest, Message: err.Error()}
	}

	// Map the request to the user model
	user, err := helpers.UserFromRegisterRequest(request)
	if err != nil {
		return dto.ServiceResponse{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	// Before creating a user, check if the user already exists
	userExists, _ := s.repo.GetUserByEmail(user.Email)
	if userExists.ID != 0 {
		return dto.ServiceResponse{Code: http.StatusConflict, Message: "User already exists with this email"}
	}

	// Save the user to the database
	err = s.repo.CreateUser(user)
	if err != nil {
		return dto.ServiceResponse{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	return dto.ServiceResponse{
		Code:    http.StatusOK,
		Message: "User registered successfully",
	}
}

//

func (s *userService) Authenticate(email, password string) dto.ServiceResponse {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return dto.ServiceResponse{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return dto.ServiceResponse{Code: http.StatusUnauthorized, Message: "Invalid credentials"}
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return dto.ServiceResponse{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	return dto.ServiceResponse{
		Code:    http.StatusOK,
		Message: "User authenticated successfully",
		Data: dto.AuthenticateUserResponse{
			Token:  token,
			UserID: user.ID,
			Email:  user.Email,
		},
	}
}
