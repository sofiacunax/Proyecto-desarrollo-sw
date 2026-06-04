package controllers

import (
	"net/http"
	"strconv"

	"proyecto-desarrollo-sw/backend/dtos"
	"proyecto-desarrollo-sw/backend/services"

	"github.com/gin-gonic/gin"
)

type EntradaController struct {
	EntradaService *services.EntradaService
}

func NewEntradaController() *EntradaController {
	return &EntradaController{
		EntradaService: services.NewEntradaService(),
	}
}

func (controller *EntradaController) ComprarEntrada(ctx *gin.Context) {

	var request dtos.ComprarEntradaDTO

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "datos invalidos"})
		return
	}

	userIDValue, _ := ctx.Get("userID")
	userID := int(userIDValue.(float64))

	err := controller.EntradaService.ComprarEntrada(
		userID,
		request.EventoID,
	)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "entrada comprada correctamente",
	})
}

func (controller *EntradaController) ObtenerMisEntradas(ctx *gin.Context) {

	userIDValue, _ := ctx.Get("userID")
	userID := int(userIDValue.(float64))

	entradas, err := controller.EntradaService.ObtenerMisEntradas(userID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error al obtener entradas",
		})
		return
	}

	ctx.JSON(http.StatusOK, entradas)
}

func (controller *EntradaController) CancelarEntrada(ctx *gin.Context) {

	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "id invalido",
		})
		return
	}

	userIDValue, _ := ctx.Get("userID")
	userID := int(userIDValue.(float64))

	err = controller.EntradaService.CancelarEntrada(
		id,
		userID,
	)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "entrada cancelada correctamente",
	})
}

func (controller *EntradaController) TransferirEntrada(ctx *gin.Context) {

	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "id invalido",
		})
		return
	}

	var request dtos.TransferirEntradaDTO

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "datos invalidos",
		})
		return
	}

	userIDValue, _ := ctx.Get("userID")
	userID := int(userIDValue.(float64))

	err = controller.EntradaService.TransferirEntrada(
		id,
		userID,
		request.NuevoUsuarioID,
	)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "entrada transferida correctamente",
	})
}