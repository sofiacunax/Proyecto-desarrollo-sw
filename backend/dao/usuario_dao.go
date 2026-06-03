package dao

import (
	"proyecto-desarrollo-sw/backend/db"
	"proyecto-desarrollo-sw/backend/models"
)

type UsuarioDAO struct {
	usuarios []models.Usuario
}

var Usuarios []models.Usuario

func (d *UsuarioDAO) CrearUsuario(usuario models.Usuario) error {

	query := `
        INSERT INTO usuarios
        (nombre, email, password_hash, rol)
        VALUES (?, ?, ?, ?)
    `

	_, err := db.DB.Exec(
		query,
		usuario.Nombre,
		usuario.Email,
		usuario.PasswordHash,
		usuario.Rol,
	)

	return err
}

func (d *UsuarioDAO) BuscarPorEmail(email string) (*models.Usuario, error) {

	query := `
		SELECT id, nombre, email, password_hash, rol
		FROM usuarios
		WHERE email = ?
	`

	var usuario models.Usuario

	err := db.DB.QueryRow(query, email).Scan(
		&usuario.ID,
		&usuario.Nombre,
		&usuario.Email,
		&usuario.PasswordHash,
		&usuario.Rol,
	)

	if err != nil {
		return nil, err
	}

	return &usuario, nil
}
