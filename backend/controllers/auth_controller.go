package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"proyecto-desarrollo-sw/backend/dtos"
	"proyecto-desarrollo-sw/backend/services"
)

type AuthController struct {
	AuthService services.AuthService
}

func (c *AuthController) Register(ctx *gin.Context) {

	var request dtos.RegisterRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "datos inválidos",
		})
		return
	}

	err := c.AuthService.Register(request)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "usuario registrado correctamente",
	})
}

func (c *AuthController) Login(ctx *gin.Context) {

	var request dtos.LoginRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "datos inválidos",
		})

		return
	}

	token, err := c.AuthService.Login(request)

	if err != nil {

		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "login correcto",
		"token":   token,
	})
}
