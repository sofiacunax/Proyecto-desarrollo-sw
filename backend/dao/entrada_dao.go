package dao

import (
	"database/sql"

	"proyecto-desarrollo-sw/backend/db"
	"proyecto-desarrollo-sw/backend/models"
)

type EntradaDAO struct{}

func NewEntradaDAO() *EntradaDAO {
	return &EntradaDAO{}
}

func (dao *EntradaDAO) CrearEntrada(entrada models.Entrada) error {

	query := `
		INSERT INTO entradas
		(usuario_id, evento_id, estado)
		VALUES (?, ?, ?)
	`

	_, err := db.DB.Exec(
		query,
		entrada.UsuarioID,
		entrada.EventoID,
		entrada.Estado,
	)

	return err
}

func (dao *EntradaDAO) ObtenerPorUsuario(usuarioID int) ([]models.Entrada, error) {

	query := `
		SELECT id, usuario_id, evento_id, estado
		FROM entradas
		WHERE usuario_id = ?
	`

	rows, err := db.DB.Query(query, usuarioID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var entradas []models.Entrada

	for rows.Next() {

		var entrada models.Entrada

		err := rows.Scan(
			&entrada.ID,
			&entrada.UsuarioID,
			&entrada.EventoID,
			&entrada.Estado,
		)

		if err != nil {
			return nil, err
		}

		entradas = append(entradas, entrada)
	}

	return entradas, nil
}

func (dao *EntradaDAO) ObtenerPorID(id int) (*models.Entrada, error) {

	query := `
		SELECT id, usuario_id, evento_id, estado
		FROM entradas
		WHERE id = ?
	`

	var entrada models.Entrada

	err := db.DB.QueryRow(query, id).Scan(
		&entrada.ID,
		&entrada.UsuarioID,
		&entrada.EventoID,
		&entrada.Estado,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &entrada, nil
}

func (dao *EntradaDAO) CancelarEntrada(id int) error {

	query := `
		UPDATE entradas
		SET estado = 'CANCELADA'
		WHERE id = ?
	`

	_, err := db.DB.Exec(query, id)

	return err
}

func (dao *EntradaDAO) TransferirEntrada(id int, nuevoUsuarioID int) error {

	query := `
		UPDATE entradas
		SET usuario_id = ?
		WHERE id = ?
	`

	_, err := db.DB.Exec(
		query,
		nuevoUsuarioID,
		id,
	)

	return err
}

func (dao *EntradaDAO) ContarEntradasActivas(eventoID int) (int, error) {

	query := `
		SELECT COUNT(*)
		FROM entradas
		WHERE evento_id = ?
		AND estado = 'ACTIVA'
	`

	var cantidad int

	err := db.DB.QueryRow(query, eventoID).Scan(&cantidad)

	if err != nil {
		return 0, err
	}

	return cantidad, nil
}

func (dao *EntradaDAO) ObtenerCapacidadEvento(eventoID int) (int, error) {

	query := `
		SELECT capacidad
		FROM eventos
		WHERE id = ?
	`

	var capacidad int

	err := db.DB.QueryRow(query, eventoID).Scan(&capacidad)

	if err != nil {
		return 0, err
	}

	return capacidad, nil
}