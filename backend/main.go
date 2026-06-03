package main

import (
	"github.com/gin-gonic/gin"

	"proyecto-desarrollo-sw/backend/controllers"
	"proyecto-desarrollo-sw/backend/db"
)

func main() {

	db.Connect()

	router := gin.Default()

	authController := controllers.AuthController{}

	router.POST("/auth/register", authController.Register)
	router.POST("/auth/login", authController.Login)

	router.Run(":8080")
}
