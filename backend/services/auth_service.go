package services

import (
	"fmt"

	"proyecto-desarrollo-sw/backend/dao"
	"proyecto-desarrollo-sw/backend/dtos"
	"proyecto-desarrollo-sw/backend/models"
	"proyecto-desarrollo-sw/backend/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UsuarioDAO dao.UsuarioDAO
}

func (s *AuthService) Register(request dtos.RegisterRequest) error {

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	usuario := models.Usuario{
		Nombre:       request.Nombre,
		Email:        request.Email,
		PasswordHash: string(hash),
		Rol:          "CLIENTE",
	}

	return s.UsuarioDAO.CrearUsuario(usuario)
}

func (s *AuthService) Login(request dtos.LoginRequest) (string, error) {

	usuario, err := s.UsuarioDAO.BuscarPorEmail(request.Email)

	if err != nil {
		return "", fmt.Errorf("credenciales inválidas")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(usuario.PasswordHash),
		[]byte(request.Password),
	)

	if err != nil {
		return "", fmt.Errorf("credenciales inválidas")
	}

	token, err := utils.GenerateToken(
		usuario.ID,
		usuario.Email,
		usuario.Rol,
	)

	if err != nil {
		return "", err
	}

	return token, nil
}
