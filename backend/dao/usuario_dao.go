package dao

import (
	"fmt"
	"proyecto-desarrollo-sw/backend/models"
)

type UsuarioDAO struct {
	usuarios []models.Usuario
}

var Usuarios []models.Usuario

func (d *UsuarioDAO) CrearUsuario(usuario models.Usuario) error {

	Usuarios = append(Usuarios, usuario)

	fmt.Println(usuario)

	return nil
}

func (d *UsuarioDAO) BuscarPorEmail(email string) (*models.Usuario, error) {

	for _, usuario := range Usuarios {

		if usuario.Email == email {
			return &usuario, nil
		}
	}

	return nil, fmt.Errorf("usuario no encontrado")
}
