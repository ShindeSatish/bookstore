package helpers

import (
	"github.com/ShindeSatish/bookstore/internal/dto"
	"github.com/ShindeSatish/bookstore/internal/models"
	"github.com/ShindeSatish/bookstore/internal/utils"
)

func UserFromRegisterRequest(request *dto.RegisterUserRequest) (models.User, error) {
	// We need to store the password in a hashed format
	password, err := utils.HashPassword(request.Password)
	if err != nil {
		return models.User{}, err
	}

	return models.User{
		Email:     request.Email,
		Password:  password,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Phone:     request.Phone,
	}, nil
}
