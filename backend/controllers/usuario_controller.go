package controllers

import (
	"net/http"
	"strconv"

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

func (controller *UsuarioController) CambiarRol(
	ctx *gin.Context,
) {

	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "id inválido",
			},
		)
		return
	}

	var request struct {
		Rol string `json:"rol"`
	}

	if err := ctx.ShouldBindJSON(
		&request,
	); err != nil {

		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "datos inválidos",
			},
		)
		return
	}

	err = controller.
		UsuarioService.
		CambiarRol(
			id,
			request.Rol,
		)

	if err != nil {

		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"message": "rol actualizado correctamente",
		},
	)
}
