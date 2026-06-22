package services

import (
	"proyecto-desarrollo-sw/backend/dao"
	"proyecto-desarrollo-sw/backend/models"
)

type UsuarioService struct {
	UsuarioDAO *dao.UsuarioDAO
}

func NewUsuarioService() *UsuarioService {
	return &UsuarioService{
		UsuarioDAO: dao.NewUsuarioDAO(),
	}
}

func (service *UsuarioService) ObtenerUsuarios() ([]models.Usuario, error) {
	return service.UsuarioDAO.ObtenerUsuarios()
}
