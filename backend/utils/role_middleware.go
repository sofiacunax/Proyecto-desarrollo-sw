package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		rol := c.GetString("rol")

		if rol != "ADMIN" {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "acceso denegado",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
