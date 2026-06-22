package dao

import (
	"errors"

	"proyecto-desarrollo-sw/backend/db"
	"proyecto-desarrollo-sw/backend/dtos"
	"proyecto-desarrollo-sw/backend/models"

	"gorm.io/gorm"
)

type EntradaDAO struct{}

func NewEntradaDAO() *EntradaDAO { return &EntradaDAO{} }

func (dao *EntradaDAO) CrearEntrada(entrada models.Entrada) error {
	return db.DB.Create(&entrada).Error
}

func (dao *EntradaDAO) ObtenerPorUsuario(usuarioID int) ([]models.Entrada, error) {
	var entradas []models.Entrada
	if err := db.DB.Where("usuario_id = ?", usuarioID).Find(&entradas).Error; err != nil {
		return nil, err
	}
	return entradas, nil
}

func (dao *EntradaDAO) ObtenerMisEntradasDTO(usuarioID int) ([]dtos.MisEntradaDTO, error) {
	var entradas []dtos.MisEntradaDTO
	err := db.DB.Model(&models.Entrada{}).
		Select("entradas.id, entradas.evento_id, eventos.titulo, eventos.fecha, eventos.ubicacion, entradas.estado").
		Joins("JOIN eventos ON entradas.evento_id = eventos.id").
		Where("entradas.usuario_id = ?", usuarioID).
		Scan(&entradas).Error
	if err != nil {
		return nil, err
	}
	return entradas, nil
}

func (dao *EntradaDAO) ObtenerPorID(id int) (*models.Entrada, error) {
	var entrada models.Entrada
	if err := db.DB.First(&entrada, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &entrada, nil
}

func (dao *EntradaDAO) CancelarEntrada(id int) error {
	return db.DB.Model(&models.Entrada{}).Where("id = ?", id).
		Update("estado", "CANCELADA").Error
}

func (dao *EntradaDAO) TransferirEntrada(id int, nuevoUsuarioID int) error {
	return db.DB.Model(&models.Entrada{}).Where("id = ?", id).
		Update("usuario_id", nuevoUsuarioID).Error
}

func (dao *EntradaDAO) ContarEntradasActivas(eventoID int) (int, error) {
	var cantidad int64
	err := db.DB.Model(&models.Entrada{}).
		Where("evento_id = ? AND estado = ?", eventoID, "ACTIVA").
		Count(&cantidad).Error
	return int(cantidad), err
}

func (dao *EntradaDAO) ObtenerCapacidadEvento(eventoID int) (int, error) {
	var evento models.Evento
	if err := db.DB.Select("capacidad").First(&evento, eventoID).Error; err != nil {
		return 0, err
	}
	return evento.Capacidad, nil
}

func (dao *EntradaDAO) ObtenerCompradoresPorEvento(
	eventoID int,
) ([]dtos.CompradorDTO, error) {

	var compradores []dtos.CompradorDTO

	err := db.DB.
		Table("entradas").
		Select(
			"usuarios.id, usuarios.nombre, usuarios.email",
		).
		Joins(
			"JOIN usuarios ON entradas.usuario_id = usuarios.id",
		).
		Where(
			"entradas.evento_id = ? AND entradas.estado = ?",
			eventoID,
			"ACTIVA",
		).
		Scan(&compradores).
		Error

	return compradores, err
}
