package controllers

import (
	"net/http"
	"strconv"

	"proyecto-desarrollo-sw/backend/dtos"
	"proyecto-desarrollo-sw/backend/services"

	"github.com/gin-gonic/gin"
)

type EventoController struct {
	EventoService *services.EventoService
}

func NewEventoController() *EventoController {
	return &EventoController{
		EventoService: services.NewEventoService(),
	}
}

func (controller *EventoController) CrearEvento(ctx *gin.Context) {
	var request dtos.EventoDTO

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "datos inválidos"})
		return
	}

	err := controller.EventoService.CrearEvento(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "evento creado correctamente"})
}

func (controller *EventoController) ObtenerEventos(ctx *gin.Context) {
	eventos, err := controller.EventoService.ObtenerEventos()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error al obtener eventos"})
		return
	}

	ctx.JSON(http.StatusOK, eventos)
}

func (controller *EventoController) ObtenerEventoPorID(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}

	evento, err := controller.EventoService.ObtenerEventoPorID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error al obtener evento"})
		return
	}

	if evento == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "evento no encontrado"})
		return
	}

	ctx.JSON(http.StatusOK, evento)
}

func (controller *EventoController) ActualizarEvento(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}

	var request dtos.EventoDTO

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "datos inválidos"})
		return
	}

	err = controller.EventoService.ActualizarEvento(id, request)
	if err != nil {
		if err.Error() == "evento no encontrado" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "evento actualizado correctamente"})
}

func (controller *EventoController) EliminarEvento(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}

	err = controller.EventoService.EliminarEvento(id)
	if err != nil {
		if err.Error() == "evento no encontrado" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error al eliminar evento"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "evento eliminado correctamente"})
}
