package controllers

import (
	"net/http"

	"proyecto-desarrollo-sw/backend/services"

	"github.com/gin-gonic/gin"
)

type UsuarioController struct {
	UsuarioService *services.UsuarioService
}

func NewUsuarioController() *UsuarioController {
	return &UsuarioController{
		UsuarioService: services.NewUsuarioService(),
	}
}

func (controller *UsuarioController) ObtenerUsuarios(ctx *gin.Context) {

	usuarios, err := controller.UsuarioService.ObtenerUsuarios()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{"error": "error al obtener usuarios"})
		return
	}

	ctx.JSON(http.StatusOK, usuarios)
}
