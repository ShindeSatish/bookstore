package repositories

import (
	"github.com/ShindeSatish/bookstore/internal/domain/abstraction"
	"github.com/ShindeSatish/bookstore/internal/models"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) abstraction.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(user models.User) error {
	return r.db.Create(&user).Error
}

func (r *userRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return user, err
}
