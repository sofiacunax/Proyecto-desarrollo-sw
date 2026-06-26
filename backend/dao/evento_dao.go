package dao

import (
	"errors"

	"proyecto-desarrollo-sw/backend/db"
	"proyecto-desarrollo-sw/backend/models"

	"gorm.io/gorm"
)

type EventoDAO struct{}

func NewEventoDAO() *EventoDAO { return &EventoDAO{} }

func (dao *EventoDAO) CrearEvento(evento models.Evento) error {
	return db.DB.Create(&evento).Error
}

func (dao *EventoDAO) ObtenerEventos(busqueda string) ([]models.Evento, error) {
	var eventos []models.Evento
	consulta := db.DB.Where("estado = ?", "ACTIVO")
	if busqueda != "" {
		filtro := "%" + busqueda + "%"
		consulta = consulta.Where(
			"titulo LIKE ? OR ubicacion LIKE ? OR categoria LIKE ?",
			filtro, filtro, filtro,
		)
	}
	if err := consulta.Find(&eventos).Error; err != nil {
		return nil, err
	}
	return eventos, nil
}

func (dao *EventoDAO) ObtenerTodosEventos(busqueda string) ([]models.Evento, error) {
	var eventos []models.Evento
	consulta := db.DB
	if busqueda != "" {
		filtro := "%" + busqueda + "%"
		consulta = consulta.Where(
			"titulo LIKE ? OR ubicacion LIKE ? OR categoria LIKE ?",
			filtro, filtro, filtro,
		)
	}
	if err := consulta.Find(&eventos).Error; err != nil {
		return nil, err
	}
	return eventos, nil
}

func (dao *EventoDAO) ObtenerEventoPorID(id int) (*models.Evento, error) {
	var evento models.Evento
	if err := db.DB.First(&evento, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &evento, nil
}

func (dao *EventoDAO) ActualizarEvento(id int, evento models.Evento) error {
	return db.DB.Model(&models.Evento{}).Where("id = ?", id).Updates(map[string]any{
		"titulo": evento.Titulo, "descripcion": evento.Descripcion,
		"fecha": evento.Fecha, "horario": evento.Horario,
		"duracion": evento.Duracion, "ubicacion": evento.Ubicacion,
		"capacidad": evento.Capacidad, "precio": evento.Precio,
		"categoria": evento.Categoria, "imagen_url": evento.ImagenURL,
		"estado": evento.Estado,
	}).Error
}

func (dao *EventoDAO) ActualizarEstadoEvento(id int, estado string) error {
	return db.DB.Model(&models.Evento{}).
		Where("id = ?", id).
		Update("estado", estado).Error
}

func (dao *EventoDAO) TieneActividad(id int) (bool, error) {
	var entradas int64
	if err := db.DB.Model(&models.Entrada{}).
		Where("evento_id = ?", id).
		Count(&entradas).Error; err != nil {
		return false, err
	}

	if entradas > 0 {
		return true, nil
	}

	var puntuaciones int64
	if err := db.DB.Model(&models.Puntuacion{}).
		Where("evento_id = ?", id).
		Count(&puntuaciones).Error; err != nil {
		return false, err
	}

	return puntuaciones > 0, nil
}

func (dao *EventoDAO) EliminarEvento(id int) error {
	return db.DB.Delete(&models.Evento{}, id).Error
}
