package main

import (
	"github.com/gin-gonic/gin"

	"proyecto-desarrollo-sw/backend/controllers"
	"proyecto-desarrollo-sw/backend/db"
	"proyecto-desarrollo-sw/backend/utils"
)

func main() {

	db.Connect()

	router := gin.Default()

	authController := controllers.AuthController{}

	router.POST("/auth/register", authController.Register)
	router.POST("/auth/login", authController.Login)

	private := router.Group("/private")
	private.Use(utils.AuthMiddleware())

	private.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "acceso permitido",
		})
	})

	router.Run(":8080")
}
