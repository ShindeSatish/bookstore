package middleware

import (
	"fmt"
	"github.com/ShindeSatish/bookstore/internal/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		claims := &utils.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return utils.JwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access! Please login to use the API."})
			c.Abort()
			return
		}

		fmt.Println(claims.UserID)
		// Token is valid; you can now set the user ID in the context if needed
		c.Set("userID", claims.UserID)

		c.Next()
	}
}
