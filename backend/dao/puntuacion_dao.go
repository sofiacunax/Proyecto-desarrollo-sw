package dao

import (
	"proyecto-desarrollo-sw/backend/db"
	"proyecto-desarrollo-sw/backend/dtos"
	"proyecto-desarrollo-sw/backend/models"
)

type PuntuacionDAO struct{}

func NewPuntuacionDAO() *PuntuacionDAO { return &PuntuacionDAO{} }

func (dao *PuntuacionDAO) CrearPuntuacion(puntuacion models.Puntuacion) error {
	return db.DB.Create(&puntuacion).Error
}

func (dao *PuntuacionDAO) ObtenerPuntuacionesPorEvento(eventoID int) ([]models.Puntuacion, error) {
	var puntuaciones []models.Puntuacion
	if err := db.DB.Where("evento_id = ?", eventoID).Find(&puntuaciones).Error; err != nil {
		return nil, err
	}
	return puntuaciones, nil
}

func (dao *PuntuacionDAO) ObtenerPromedioPorEvento(eventoID int) (float64, error) {
	var promedio float64
	err := db.DB.Model(&models.Puntuacion{}).
		Select("COALESCE(AVG(puntuacion), 0)").
		Where("evento_id = ?", eventoID).
		Scan(&promedio).Error
	return promedio, err
}

func (dao *PuntuacionDAO) ObtenerRankingEventos() ([]dtos.RankingDTO, error) {
	var ranking []dtos.RankingDTO
	err := db.DB.Model(&models.Evento{}).
		Select("eventos.id AS evento_id, eventos.titulo, COALESCE(AVG(puntuaciones.puntuacion), 0) AS promedio").
		Joins("LEFT JOIN puntuaciones ON eventos.id = puntuaciones.evento_id").
		Group("eventos.id, eventos.titulo").
		Order("promedio DESC").
		Scan(&ranking).Error
	if err != nil {
		return nil, err
	}
	return ranking, nil
}
