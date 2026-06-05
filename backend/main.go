package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"proyecto-desarrollo-sw/backend/controllers"
	"proyecto-desarrollo-sw/backend/db"
	"proyecto-desarrollo-sw/backend/utils"
)

func main() {

	db.Connect()

	router := gin.Default()
	router.Use(cors.Default())

	authController := controllers.AuthController{}
	eventoController := controllers.NewEventoController()
	puntuacionController := controllers.NewPuntuacionController()
	entradaController := controllers.NewEntradaController()

	router.POST("/auth/register", authController.Register)
	router.POST("/auth/login", authController.Login)

	router.GET("/eventos", eventoController.ObtenerEventos)
	router.GET("/eventos/:id", eventoController.ObtenerEventoPorID)

	router.GET("/eventos/:id/puntuaciones", puntuacionController.ObtenerPuntuacionesPorEvento)
	router.GET("/eventos/:id/promedio", puntuacionController.ObtenerPromedioPorEvento)
	router.GET("/eventos/ranking", puntuacionController.ObtenerRankingEventos)

	private := router.Group("/private")
	private.Use(utils.AuthMiddleware())
	private.POST("/puntuaciones", puntuacionController.CrearPuntuacion)
	private.POST("/entradas", entradaController.ComprarEntrada)
	private.GET("/mis-entradas", entradaController.ObtenerMisEntradas)
	private.PUT("/entradas/:id/cancelar", entradaController.CancelarEntrada)
	private.PUT("/entradas/:id/transferir", entradaController.TransferirEntrada)

	private.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "acceso permitido",
		})
	})

	admin := private.Group("/admin")
	admin.Use(utils.AdminMiddleware())

	admin.POST("/eventos", eventoController.CrearEvento)
	admin.PUT("/eventos/:id", eventoController.ActualizarEvento)
	admin.DELETE("/eventos/:id", eventoController.EliminarEvento)

	admin.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "solo admin",
		})
	})

	router.Run(":8080")
}
