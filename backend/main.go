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
	eventoController := controllers.NewEventoController()

	router.POST("/auth/register", authController.Register)
	router.POST("/auth/login", authController.Login)

	router.POST("/eventos", eventoController.CrearEvento)
	router.GET("/eventos", eventoController.ObtenerEventos)
	router.GET("/eventos/:id", eventoController.ObtenerEventoPorID)
	router.PUT("/eventos/:id", eventoController.ActualizarEvento)
	router.DELETE("/eventos/:id", eventoController.EliminarEvento)

	private := router.Group("/private")
	private.Use(utils.AuthMiddleware())

	private.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "acceso permitido",
		})
	})

	admin := private.Group("/admin")
	admin.Use(utils.AdminMiddleware())

	admin.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "solo admin",
		})
	})

	router.Run(":8080")
}
