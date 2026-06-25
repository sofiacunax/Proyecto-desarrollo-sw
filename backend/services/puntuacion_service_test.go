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
		results   []databaseResult
		wantErr   string
	}{
		{name: "menor que cero", usuarioID: 1, request: dtos.PuntuacionDTO{EventoID: 1, Puntuacion: -1}, wantErr: "puntuacion"},
		{name: "mayor que cinco", usuarioID: 1, request: dtos.PuntuacionDTO{EventoID: 1, Puntuacion: 6}, wantErr: "puntuacion"},
		{name: "usuario invalido", usuarioID: 0, request: dtos.PuntuacionDTO{EventoID: 1, Puntuacion: 4}, wantErr: "usuario invalido"},
		{name: "evento invalido", usuarioID: 1, request: dtos.PuntuacionDTO{EventoID: 0, Puntuacion: 4}, wantErr: "evento invalido"},
		{name: "evento cancelado", usuarioID: 1, request: dtos.PuntuacionDTO{EventoID: 2, Puntuacion: 5}, results: []databaseResult{eventoCompraResult("CANCELADO", 20)}, wantErr: "evento cancelado"},
		{name: "valida", usuarioID: 1, request: dtos.PuntuacionDTO{EventoID: 2, Puntuacion: 5}, results: []databaseResult{eventoCompraResult("ACTIVO", 20), {}}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if len(test.results) > 0 {
				useDatabaseStub(t, test.results...)
			}
			err := NewPuntuacionService().CrearPuntuacion(test.usuarioID, test.request)
			if test.wantErr == "" && err != nil {
				t.Fatalf("error inesperado: %v", err)
			}
			if test.wantErr != "" && (err == nil || !stringsContains(err.Error(), test.wantErr)) {
				t.Fatalf("se esperaba error %q, se obtuvo %v", test.wantErr, err)
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
