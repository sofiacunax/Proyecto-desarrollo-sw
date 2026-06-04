package controllers

import (
	"net/http"
	"strconv"

	"proyecto-desarrollo-sw/backend/dtos"
	"proyecto-desarrollo-sw/backend/services"

	"github.com/gin-gonic/gin"
)

type PuntuacionController struct {
	PuntuacionService *services.PuntuacionService
}

func NewPuntuacionController() *PuntuacionController {
	return &PuntuacionController{
		PuntuacionService: services.NewPuntuacionService(),
	}
}

func (controller *PuntuacionController) CrearPuntuacion(ctx *gin.Context) {
	var request dtos.PuntuacionDTO

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "datos invalidos"})
		return
	}

	usuarioIDValue, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "usuario no autenticado"})
		return
	}

	usuarioIDFloat, ok := usuarioIDValue.(float64)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "usuario invalido"})
		return
	}

	usuarioID := int(usuarioIDFloat)

	err := controller.PuntuacionService.CrearPuntuacion(usuarioID, request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "puntuacion creada correctamente"})
}

func (controller *PuntuacionController) ObtenerPuntuacionesPorEvento(ctx *gin.Context) {
	idParam := ctx.Param("id")

	eventoID, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id invalido"})
		return
	}

	puntuaciones, err := controller.PuntuacionService.ObtenerPuntuacionesPorEvento(eventoID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error al obtener puntuaciones"})
		return
	}

	ctx.JSON(http.StatusOK, puntuaciones)
}

func (controller *PuntuacionController) ObtenerPromedioPorEvento(ctx *gin.Context) {
	idParam := ctx.Param("id")

	eventoID, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id invalido"})
		return
	}

	promedio, err := controller.PuntuacionService.ObtenerPromedioPorEvento(eventoID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error al obtener promedio"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"promedio": promedio})
}

func (controller *PuntuacionController) ObtenerRankingEventos(ctx *gin.Context) {
	ranking, err := controller.PuntuacionService.ObtenerRankingEventos()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error al obtener ranking"})
		return
	}

	ctx.JSON(http.StatusOK, ranking)
}
