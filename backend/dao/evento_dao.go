package dao

import (
	"database/sql"
	"proyecto-desarrollo-sw/backend/db"
	"proyecto-desarrollo-sw/backend/models"
)

type EventoDAO struct{}

func NewEventoDAO() *EventoDAO {
	return &EventoDAO{}
}

func (dao *EventoDAO) CrearEvento(evento models.Evento) error {
	query := `
		INSERT INTO eventos 
		(titulo, descripcion, fecha, horario, duracion, ubicacion, capacidad, precio, categoria, imagen_url, estado)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := db.DB.Exec(
		query,
		evento.Titulo,
		evento.Descripcion,
		evento.Fecha,
		evento.Horario,
		evento.Duracion,
		evento.Ubicacion,
		evento.Capacidad,
		evento.Precio,
		evento.Categoria,
		evento.ImagenURL,
		evento.Estado,
	)

	return err
}

func (dao *EventoDAO) ObtenerEventos() ([]models.Evento, error) {
	query := `
		SELECT id, titulo, descripcion, fecha, horario, duracion, ubicacion, capacidad, precio, categoria, imagen_url, estado
		FROM eventos
	`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var eventos []models.Evento

	for rows.Next() {
		var evento models.Evento

		err := rows.Scan(
			&evento.ID,
			&evento.Titulo,
			&evento.Descripcion,
			&evento.Fecha,
			&evento.Horario,
			&evento.Duracion,
			&evento.Ubicacion,
			&evento.Capacidad,
			&evento.Precio,
			&evento.Categoria,
			&evento.ImagenURL,
			&evento.Estado,
		)

		if err != nil {
			return nil, err
		}

		eventos = append(eventos, evento)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return eventos, nil
}

func (dao *EventoDAO) ObtenerEventoPorID(id int) (*models.Evento, error) {
	query := `
		SELECT id, titulo, descripcion, fecha, horario, duracion, ubicacion, capacidad, precio, categoria, imagen_url, estado
		FROM eventos
		WHERE id = ?
	`

	var evento models.Evento

	err := db.DB.QueryRow(query, id).Scan(
		&evento.ID,
		&evento.Titulo,
		&evento.Descripcion,
		&evento.Fecha,
		&evento.Horario,
		&evento.Duracion,
		&evento.Ubicacion,
		&evento.Capacidad,
		&evento.Precio,
		&evento.Categoria,
		&evento.ImagenURL,
		&evento.Estado,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &evento, nil
}

func (dao *EventoDAO) ActualizarEvento(id int, evento models.Evento) error {
	query := `
		UPDATE eventos
		SET titulo = ?, descripcion = ?, fecha = ?, horario = ?, duracion = ?, ubicacion = ?, capacidad = ?, precio = ?, categoria = ?, imagen_url = ?, estado = ?
		WHERE id = ?
	`

	_, err := db.DB.Exec(
		query,
		evento.Titulo,
		evento.Descripcion,
		evento.Fecha,
		evento.Horario,
		evento.Duracion,
		evento.Ubicacion,
		evento.Capacidad,
		evento.Precio,
		evento.Categoria,
		evento.ImagenURL,
		evento.Estado,
		id,
	)

	return err
}

func (dao *EventoDAO) EliminarEvento(id int) error {
	query := `
		DELETE FROM eventos
		WHERE id = ?
	`

	_, err := db.DB.Exec(query, id)

	return err
}
