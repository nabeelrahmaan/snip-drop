package middleware

import (
	"codeDrop/internal/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var key = utils.AccessSecret
func AuthMiddleware() gin.HandlerFunc {
	return func (c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			c.JSON(401, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return 
		}

		parts := strings.Split(header, " ")
		if len(parts) != 2 {
			c.JSON(401, gin.H{"error": "invalid authorization format"})
			c.Abort()
			return 
		}

		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
			return key, nil
		})
		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "invalid token"})
			c.Abort()
			return 
		}

		claims := token.Claims.(jwt.MapClaims)

		c.Set("user_id", claims["user_id"])
		c.Set("role", claims["role"])

		c.Next()
	}
}