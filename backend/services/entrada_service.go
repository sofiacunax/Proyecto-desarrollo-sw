package services

import (
	"errors"

	"proyecto-desarrollo-sw/backend/dao"
	"proyecto-desarrollo-sw/backend/dtos"
	"proyecto-desarrollo-sw/backend/models"
)

type EntradaService struct {
	EntradaDAO *dao.EntradaDAO
}

func NewEntradaService() *EntradaService {
	return &EntradaService{
		EntradaDAO: dao.NewEntradaDAO(),
	}
}

func (service *EntradaService) ComprarEntrada(usuarioID int, eventoID int) error {

	if eventoID <= 0 {
		return errors.New("evento invalido")
	}

	capacidad, err := service.EntradaDAO.ObtenerCapacidadEvento(eventoID)
	if err != nil {
		return err
	}

	ocupados, err := service.EntradaDAO.ContarEntradasActivas(eventoID)
	if err != nil {
		return err
	}

	if ocupados >= capacidad {
		return errors.New("no hay cupos disponibles")
	}

	entrada := models.Entrada{
		UsuarioID: usuarioID,
		EventoID:  eventoID,
		Estado:    "ACTIVA",
	}

	return service.EntradaDAO.CrearEntrada(entrada)
}

func (service *EntradaService) ObtenerMisEntradas(usuarioID int) ([]models.Entrada, error) {
	return service.EntradaDAO.ObtenerPorUsuario(usuarioID)
}
func (service *EntradaService) ObtenerMisEntradasDTO(usuarioID int) ([]dtos.MisEntradaDTO, error) {
	return service.EntradaDAO.ObtenerMisEntradasDTO(usuarioID)
}

func (service *EntradaService) CancelarEntrada(id int, usuarioID int) error {

	entrada, err := service.EntradaDAO.ObtenerPorID(id)

	if err != nil {
		return err
	}

	if entrada == nil {
		return errors.New("entrada no encontrada")
	}

	if entrada.UsuarioID != usuarioID {
		return errors.New("no puede cancelar esta entrada")
	}

	return service.EntradaDAO.CancelarEntrada(id)
}

func (service *EntradaService) TransferirEntrada(id int, usuarioID int, nuevoUsuarioID int) error {

	entrada, err := service.EntradaDAO.ObtenerPorID(id)

	if err != nil {
		return err
	}

	if entrada == nil {
		return errors.New("entrada no encontrada")
	}

	if entrada.UsuarioID != usuarioID {
		return errors.New("no puede transferir esta entrada")
	}
	if entrada.Estado == "CANCELADA" {
    return errors.New("no se puede transferir una entrada cancelada")
}

	return service.EntradaDAO.TransferirEntrada(id, nuevoUsuarioID)
}