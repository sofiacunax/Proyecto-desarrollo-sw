package services

import (
	"errors"
	"proyecto-desarrollo-sw/backend/dao"
	"proyecto-desarrollo-sw/backend/dtos"
	"proyecto-desarrollo-sw/backend/models"
	"strings"
)

const (
	EstadoEventoActivo    = "ACTIVO"
	EstadoEventoCancelado = "CANCELADO"

	MensajeEventoConActividad = "No se puede eliminar un evento que posee actividad. Si desea que deje de estar disponible, cambie su estado a CANCELADO."
)

type EventoService struct {
	EventoDAO *dao.EventoDAO
}

func NewEventoService() *EventoService {
	return &EventoService{
		EventoDAO: dao.NewEventoDAO(),
	}
}

func (service *EventoService) CrearEvento(request dtos.EventoDTO) error {
	if request.Titulo == "" {
		return errors.New("el titulo es obligatorio")
	}

	if request.Fecha == "" {
		return errors.New("la fecha es obligatoria")
	}

	if request.Horario == "" {
		return errors.New("el horario es obligatorio")
	}

	if request.Ubicacion == "" {
		return errors.New("la ubicacion es obligatoria")
	}

	if request.Capacidad <= 0 {
		return errors.New("la capacidad debe ser mayor a cero")
	}

	if request.Precio < 0 {
		return errors.New("el precio no puede ser negativo")
	}

	estado := strings.ToUpper(strings.TrimSpace(request.Estado))
	if estado == "" {
		estado = EstadoEventoActivo
	}
	if !estadoEventoValido(estado) {
		return errors.New("estado invalido. Debe ser ACTIVO o CANCELADO")
	}

	evento := models.Evento{
		Titulo:      request.Titulo,
		Descripcion: request.Descripcion,
		Fecha:       request.Fecha,
		Horario:     request.Horario,
		Duracion:    request.Duracion,
		Ubicacion:   request.Ubicacion,
		Capacidad:   request.Capacidad,
		Precio:      request.Precio,
		Categoria:   request.Categoria,
		ImagenURL:   request.ImagenURL,
		Estado:      estado,
	}

	return service.EventoDAO.CrearEvento(evento)
}

func (service *EventoService) ObtenerEventos(busqueda string) ([]models.Evento, error) {
	return service.EventoDAO.ObtenerEventos(busqueda)
}

func (service *EventoService) ObtenerTodosEventos(busqueda string) ([]models.Evento, error) {
	return service.EventoDAO.ObtenerTodosEventos(busqueda)
}

func (service *EventoService) ObtenerEventoPorID(id int) (*models.Evento, error) {
	return service.EventoDAO.ObtenerEventoPorID(id)
}

func (service *EventoService) ActualizarEvento(id int, request dtos.EventoDTO) error {
	eventoExistente, err := service.EventoDAO.ObtenerEventoPorID(id)
	if err != nil {
		return err
	}

	if eventoExistente == nil {
		return errors.New("evento no encontrado")
	}

	if request.Titulo == "" {
		return errors.New("el titulo es obligatorio")
	}

	if request.Fecha == "" {
		return errors.New("la fecha es obligatoria")
	}

	if request.Horario == "" {
		return errors.New("el horario es obligatorio")
	}

	if request.Ubicacion == "" {
		return errors.New("la ubicacion es obligatoria")
	}

	if request.Capacidad <= 0 {
		return errors.New("la capacidad debe ser mayor a cero")
	}

	if request.Precio < 0 {
		return errors.New("el precio no puede ser negativo")
	}

	estado := strings.ToUpper(strings.TrimSpace(request.Estado))
	if estado == "" {
		estado = EstadoEventoActivo
	}
	if !estadoEventoValido(estado) {
		return errors.New("estado invalido. Debe ser ACTIVO o CANCELADO")
	}

	evento := models.Evento{
		Titulo:      request.Titulo,
		Descripcion: request.Descripcion,
		Fecha:       request.Fecha,
		Horario:     request.Horario,
		Duracion:    request.Duracion,
		Ubicacion:   request.Ubicacion,
		Capacidad:   request.Capacidad,
		Precio:      request.Precio,
		Categoria:   request.Categoria,
		ImagenURL:   request.ImagenURL,
		Estado:      estado,
	}

	return service.EventoDAO.ActualizarEvento(id, evento)
}

func (service *EventoService) CambiarEstadoEvento(id int, estado string) error {
	eventoExistente, err := service.EventoDAO.ObtenerEventoPorID(id)
	if err != nil {
		return err
	}

	if eventoExistente == nil {
		return errors.New("evento no encontrado")
	}

	estado = strings.ToUpper(strings.TrimSpace(estado))
	if !estadoEventoValido(estado) {
		return errors.New("estado invalido. Debe ser ACTIVO o CANCELADO")
	}

	if eventoExistente.Estado == estado {
		return errors.New("el evento ya se encuentra en estado " + estado)
	}

	return service.EventoDAO.ActualizarEstadoEvento(id, estado)
}

func (service *EventoService) EliminarEvento(id int) error {
	eventoExistente, err := service.EventoDAO.ObtenerEventoPorID(id)
	if err != nil {
		return err
	}

	if eventoExistente == nil {
		return errors.New("evento no encontrado")
	}

	tieneActividad, err := service.EventoDAO.TieneActividad(id)
	if err != nil {
		return err
	}

	if tieneActividad {
		return errors.New(MensajeEventoConActividad)
	}

	return service.EventoDAO.EliminarEvento(id)
}

func estadoEventoValido(estado string) bool {
	return estado == EstadoEventoActivo || estado == EstadoEventoCancelado
}
