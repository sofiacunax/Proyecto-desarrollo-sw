package services

import (
	"errors"

	"proyecto-desarrollo-sw/backend/dao"
	"proyecto-desarrollo-sw/backend/dtos"
	"proyecto-desarrollo-sw/backend/models"
)

type PuntuacionService struct {
	PuntuacionDAO *dao.PuntuacionDAO
}

func NewPuntuacionService() *PuntuacionService {
	return &PuntuacionService{
		PuntuacionDAO: dao.NewPuntuacionDAO(),
	}
}

func (service *PuntuacionService) CrearPuntuacion(request dtos.PuntuacionDTO) error {

	if request.Puntuacion < 0 || request.Puntuacion > 5 {
		return errors.New("la puntuacion debe estar entre 0 y 5")
	}

	if request.UsuarioID <= 0 {
		return errors.New("usuario invalido")
	}

	if request.EventoID <= 0 {
		return errors.New("evento invalido")
	}

	puntuacion := models.Puntuacion{
		UsuarioID:  request.UsuarioID,
		EventoID:   request.EventoID,
		Puntuacion: request.Puntuacion,
	}

	return service.PuntuacionDAO.CrearPuntuacion(puntuacion)
}

func (service *PuntuacionService) ObtenerPuntuacionesPorEvento(eventoID int) ([]models.Puntuacion, error) {
	return service.PuntuacionDAO.ObtenerPuntuacionesPorEvento(eventoID)
}

func (service *PuntuacionService) ObtenerPromedioPorEvento(eventoID int) (float64, error) {
	return service.PuntuacionDAO.ObtenerPromedioPorEvento(eventoID)
}
