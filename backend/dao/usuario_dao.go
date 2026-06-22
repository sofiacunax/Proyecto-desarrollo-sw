package dao

import (
	"proyecto-desarrollo-sw/backend/db"
	"proyecto-desarrollo-sw/backend/models"
)

type UsuarioDAO struct{}

func NewUsuarioDAO() *UsuarioDAO {
	return &UsuarioDAO{}
}

func (d *UsuarioDAO) CrearUsuario(usuario models.Usuario) error {
	return db.DB.Create(&usuario).Error
}

func (d *UsuarioDAO) BuscarPorEmail(email string) (*models.Usuario, error) {
	var usuario models.Usuario

	if err := db.DB.Where("email = ?", email).
		First(&usuario).Error; err != nil {
		return nil, err
	}

	return &usuario, nil
}

func (d *UsuarioDAO) ObtenerUsuarios() ([]models.Usuario, error) {

	var usuarios []models.Usuario

	err := db.DB.Find(&usuarios).Error

	return usuarios, err
}

func (d *UsuarioDAO) ActualizarRol(
	id int,
	rol string,
) error {

	return db.DB.Model(&models.Usuario{}).
		Where("id = ?", id).
		Update("rol", rol).Error
}
