package services

import (
	"errors"

	"proyecto-desarrollo-sw/backend/dao"
	"proyecto-desarrollo-sw/backend/dtos"
	"proyecto-desarrollo-sw/backend/models"
)

type PuntuacionService struct {
	PuntuacionDAO *dao.PuntuacionDAO
	EventoDAO     *dao.EventoDAO
}

func NewPuntuacionService() *PuntuacionService {
	return &PuntuacionService{
		PuntuacionDAO: dao.NewPuntuacionDAO(),
		EventoDAO:     dao.NewEventoDAO(),
	}
}

func (service *PuntuacionService) CrearPuntuacion(usuarioID int, request dtos.PuntuacionDTO) error {

	if request.Puntuacion < 0 || request.Puntuacion > 5 {
		return errors.New("la puntuacion debe estar entre 0 y 5")
	}

	if usuarioID <= 0 {
		return errors.New("usuario invalido")
	}

	if request.EventoID <= 0 {
		return errors.New("evento invalido")
	}

	evento, err := service.EventoDAO.ObtenerEventoPorID(request.EventoID)
	if err != nil {
		return err
	}

	if evento == nil {
		return errors.New("evento no encontrado")
	}

	if evento.Estado != EstadoEventoActivo {
		return errors.New("No se puede puntuar un evento cancelado.")
	}

	puntuacion := models.Puntuacion{
		UsuarioID:  usuarioID,
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

func (service *PuntuacionService) ObtenerRankingEventos() ([]dtos.RankingDTO, error) {
	return service.PuntuacionDAO.ObtenerRankingEventos()
}
