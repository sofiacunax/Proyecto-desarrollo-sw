package utils

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {

			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "token requerido",
			})

			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(
			authHeader,
			"Bearer ",
		)

		token, err := jwt.Parse(
			tokenString,
			func(token *jwt.Token) (interface{}, error) {
				return SecretKey, nil
			},
		)

		if err != nil || !token.Valid {

			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "token inválido",
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
