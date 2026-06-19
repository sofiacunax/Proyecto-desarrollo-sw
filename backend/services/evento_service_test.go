package services

import (
	"testing"

	"proyecto-desarrollo-sw/backend/dtos"
)

func TestCrearEventoSinTitulo(t *testing.T) {
	service := NewEventoService()

	request := dtos.EventoDTO{
		Fecha:     "2026-01-01",
		Horario:   "20:00",
		Ubicacion: "Cordoba",
		Capacidad: 100,
		Precio:    1000,
	}

	err := service.CrearEvento(request)

	if err == nil {
		t.Error("se esperaba error por titulo vacio")
	}
}

func TestCrearEventoSinFecha(t *testing.T) {
	service := NewEventoService()

	request := dtos.EventoDTO{
		Titulo:    "Evento",
		Horario:   "20:00",
		Ubicacion: "Cordoba",
		Capacidad: 100,
		Precio:    1000,
	}

	err := service.CrearEvento(request)

	if err == nil {
		t.Error("se esperaba error por fecha vacia")
	}
}

func TestCrearEventoSinHorario(t *testing.T) {
	service := NewEventoService()

	request := dtos.EventoDTO{
		Titulo:    "Evento",
		Fecha:     "2026-01-01",
		Ubicacion: "Cordoba",
		Capacidad: 100,
		Precio:    1000,
	}

	err := service.CrearEvento(request)

	if err == nil {
		t.Error("se esperaba error por horario vacio")
	}
}

func TestCrearEventoSinUbicacion(t *testing.T) {
	service := NewEventoService()

	request := dtos.EventoDTO{
		Titulo:    "Evento",
		Fecha:     "2026-01-01",
		Horario:   "20:00",
		Capacidad: 100,
		Precio:    1000,
	}

	err := service.CrearEvento(request)

	if err == nil {
		t.Error("se esperaba error por ubicacion vacia")
	}
}

func TestCrearEventoCapacidadInvalida(t *testing.T) {
	service := NewEventoService()

	request := dtos.EventoDTO{
		Titulo:    "Evento",
		Fecha:     "2026-01-01",
		Horario:   "20:00",
		Ubicacion: "Cordoba",
		Capacidad: 0,
		Precio:    1000,
	}

	err := service.CrearEvento(request)

	if err == nil {
		t.Error("se esperaba error por capacidad invalida")
	}
}

func TestCrearEventoPrecioNegativo(t *testing.T) {
	service := NewEventoService()

	request := dtos.EventoDTO{
		Titulo:    "Evento",
		Fecha:     "2026-01-01",
		Horario:   "20:00",
		Ubicacion: "Cordoba",
		Capacidad: 100,
		Precio:    -1,
	}

	err := service.CrearEvento(request)

	if err == nil {
		t.Error("se esperaba error por precio negativo")
	}
}
