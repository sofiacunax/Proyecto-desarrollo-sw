package services

import (
	"database/sql/driver"
	"testing"

	"proyecto-desarrollo-sw/backend/dtos"
)

func TestPuntuacionServiceCrearPuntuacion(t *testing.T) {
	tests := []struct {
		name      string
		usuarioID int
		request   dtos.PuntuacionDTO
		valid     bool
	}{
		{name: "menor que cero", usuarioID: 1, request: dtos.PuntuacionDTO{EventoID: 1, Puntuacion: -1}},
		{name: "mayor que cinco", usuarioID: 1, request: dtos.PuntuacionDTO{EventoID: 1, Puntuacion: 6}},
		{name: "usuario invalido", usuarioID: 0, request: dtos.PuntuacionDTO{EventoID: 1, Puntuacion: 4}},
		{name: "evento invalido", usuarioID: 1, request: dtos.PuntuacionDTO{EventoID: 0, Puntuacion: 4}},
		{name: "valida", usuarioID: 1, request: dtos.PuntuacionDTO{EventoID: 2, Puntuacion: 5}, valid: true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.valid {
				useDatabaseStub(t, databaseResult{})
			}
			err := NewPuntuacionService().CrearPuntuacion(test.usuarioID, test.request)
			if test.valid && err != nil {
				t.Fatalf("error inesperado: %v", err)
			}
			if !test.valid && err == nil {
				t.Fatal("se esperaba error de validacion")
			}
		})
	}
}

func TestPuntuacionServiceConsultas(t *testing.T) {
	t.Run("puntuaciones", func(t *testing.T) {
		useDatabaseStub(t, databaseResult{
			columns: []string{"id", "usuario_id", "evento_id", "puntuacion"},
			rows:    [][]driver.Value{{int64(1), int64(2), int64(3), int64(5)}},
		})
		items, err := NewPuntuacionService().ObtenerPuntuacionesPorEvento(3)
		if err != nil || len(items) != 1 || items[0].Puntuacion != 5 {
			t.Fatalf("resultado inesperado: %#v, %v", items, err)
		}
	})

	t.Run("promedio", func(t *testing.T) {
		useDatabaseStub(t, databaseResult{columns: []string{"promedio"}, rows: [][]driver.Value{{4.5}}})
		value, err := NewPuntuacionService().ObtenerPromedioPorEvento(3)
		if err != nil || value != 4.5 {
			t.Fatalf("resultado inesperado: %v, %v", value, err)
		}
	})

	t.Run("ranking", func(t *testing.T) {
		useDatabaseStub(t, databaseResult{
			columns: []string{"evento_id", "titulo", "promedio"},
			rows:    [][]driver.Value{{int64(3), "Recital", 4.5}},
		})
		items, err := NewPuntuacionService().ObtenerRankingEventos()
		if err != nil || len(items) != 1 || items[0].Titulo != "Recital" {
			t.Fatalf("resultado inesperado: %#v, %v", items, err)
		}
	})
}
