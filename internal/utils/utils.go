package utils

import (
	"fmt"
	"github.com/ShindeSatish/bookstore/internal/dto"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"os"
	"time"
)

func NewSuccessResponse(message string, data interface{}) dto.APIResponse {
	return dto.APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
}

func NewErrorResponse(message string, err interface{}) dto.APIResponse {
	return dto.APIResponse{
		Success: false,
		Message: message,
		Error:   err,
	}
}

var JwtKey = []byte(os.Getenv("JWT_KEY")) // this should be in a secure place like environment variables

type Claims struct {
	UserID uint
	jwt.StandardClaims
}

func GenerateToken(userID uint) (string, error) {
	expirationTime := time.Now().Add(12 * time.Hour)
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)

	return tokenString, err
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12) // 12 is the cost for hashing the password
	return string(bytes), err
}

// ComparePassword checks if the provided password is correct

func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func randomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	rand.Seed(time.Now().UnixNano())
	s := make([]byte, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func RandomEmail() string {
	return fmt.Sprintf("%s@example.com", randomString(10))
}
