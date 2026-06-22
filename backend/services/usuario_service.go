package services

import (
	"errors"

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

func (service *UsuarioService) CambiarRol(
	id int,
	rol string,
) error {

	if rol != "ADMIN" &&
		rol != "CLIENTE" {

		return errors.New(
			"rol inválido",
		)
	}

	return service.UsuarioDAO.
		ActualizarRol(id, rol)
}
