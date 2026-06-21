package services

import (
	"database/sql/driver"
	"errors"
	"testing"

	"proyecto-desarrollo-sw/backend/models"
)

func TestEntradaServiceComprarEntrada(t *testing.T) {
	tests := []struct {
		name    string
		evento  int
		results []databaseResult
		wantErr string
	}{
		{name: "evento invalido", evento: 0, wantErr: "evento invalido"},
		{name: "error capacidad", evento: 1, results: []databaseResult{{err: errors.New("fallo capacidad")}}, wantErr: "fallo capacidad"},
		{name: "error ocupacion", evento: 1, results: []databaseResult{{columns: []string{"capacidad"}, rows: [][]driver.Value{{int64(10)}}}, {err: errors.New("fallo ocupacion")}}, wantErr: "fallo ocupacion"},
		{name: "sin cupo", evento: 1, results: []databaseResult{{columns: []string{"capacidad"}, rows: [][]driver.Value{{int64(2)}}}, {columns: []string{"cantidad"}, rows: [][]driver.Value{{int64(2)}}}}, wantErr: "no hay cupos"},
		{name: "compra exitosa", evento: 1, results: []databaseResult{{columns: []string{"capacidad"}, rows: [][]driver.Value{{int64(2)}}}, {columns: []string{"cantidad"}, rows: [][]driver.Value{{int64(1)}}}, {}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if len(test.results) > 0 {
				useDatabaseStub(t, test.results...)
			}
			err := NewEntradaService().ComprarEntrada(5, test.evento)
			if test.wantErr == "" && err != nil {
				t.Fatalf("error inesperado: %v", err)
			}
			if test.wantErr != "" && (err == nil || !stringsContains(err.Error(), test.wantErr)) {
				t.Fatalf("se esperaba error %q, se obtuvo %v", test.wantErr, err)
			}
		})
	}
}

func TestEntradaServiceConsultas(t *testing.T) {
	t.Run("modelos", func(t *testing.T) {
		useDatabaseStub(t, databaseResult{
			columns: []string{"id", "usuario_id", "evento_id", "estado"},
			rows:    [][]driver.Value{{int64(1), int64(2), int64(3), "ACTIVA"}},
		})
		items, err := NewEntradaService().ObtenerMisEntradas(2)
		if err != nil || len(items) != 1 {
			t.Fatalf("resultado inesperado: %#v, %v", items, err)
		}
	})

	t.Run("dto", func(t *testing.T) {
		useDatabaseStub(t, databaseResult{
			columns: []string{"id", "evento_id", "titulo", "fecha", "ubicacion", "estado"},
			rows:    [][]driver.Value{{int64(1), int64(3), "Recital", "2026-01-01", "Cordoba", "ACTIVA"}},
		})
		items, err := NewEntradaService().ObtenerMisEntradasDTO(2)
		if err != nil || len(items) != 1 || items[0].Titulo != "Recital" {
			t.Fatalf("resultado inesperado: %#v, %v", items, err)
		}
	})
}

func TestEntradaServiceCancelarEntrada(t *testing.T) {
	tests := []struct {
		name    string
		results []databaseResult
		wantErr string
	}{
		{name: "error consulta", results: []databaseResult{{err: errors.New("fallo consulta")}}, wantErr: "fallo consulta"},
		{name: "inexistente", results: []databaseResult{{columns: []string{"id", "usuario_id", "evento_id", "estado"}}}, wantErr: "entrada no encontrada"},
		{name: "otro propietario", results: []databaseResult{entradaResult(models.Entrada{ID: 1, UsuarioID: 9, EventoID: 2, Estado: "ACTIVA"})}, wantErr: "no puede cancelar"},
		{name: "exitosa", results: []databaseResult{entradaResult(models.Entrada{ID: 1, UsuarioID: 5, EventoID: 2, Estado: "ACTIVA"}), {}}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			useDatabaseStub(t, test.results...)
			err := NewEntradaService().CancelarEntrada(1, 5)
			assertServiceError(t, err, test.wantErr)
		})
	}
}

func TestEntradaServiceTransferirEntrada(t *testing.T) {
	tests := []struct {
		name    string
		results []databaseResult
		wantErr string
	}{
		{name: "inexistente", results: []databaseResult{{columns: []string{"id", "usuario_id", "evento_id", "estado"}}}, wantErr: "entrada no encontrada"},
		{name: "otro propietario", results: []databaseResult{entradaResult(models.Entrada{ID: 1, UsuarioID: 9, Estado: "ACTIVA"})}, wantErr: "no puede transferir"},
		{name: "cancelada", results: []databaseResult{entradaResult(models.Entrada{ID: 1, UsuarioID: 5, Estado: "CANCELADA"})}, wantErr: "cancelada"},
		{name: "exitosa", results: []databaseResult{entradaResult(models.Entrada{ID: 1, UsuarioID: 5, Estado: "ACTIVA"}), {}}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			useDatabaseStub(t, test.results...)
			err := NewEntradaService().TransferirEntrada(1, 5, 8)
			assertServiceError(t, err, test.wantErr)
		})
	}
}

func entradaResult(item models.Entrada) databaseResult {
	return databaseResult{
		columns: []string{"id", "usuario_id", "evento_id", "estado"},
		rows:    [][]driver.Value{{int64(item.ID), int64(item.UsuarioID), int64(item.EventoID), item.Estado}},
	}
}

func assertServiceError(t *testing.T, err error, want string) {
	t.Helper()
	if want == "" && err != nil {
		t.Fatalf("error inesperado: %v", err)
	}
	if want != "" && (err == nil || !stringsContains(err.Error(), want)) {
		t.Fatalf("se esperaba error %q, se obtuvo %v", want, err)
	}
}

func stringsContains(value, part string) bool {
	for i := 0; i+len(part) <= len(value); i++ {
		if value[i:i+len(part)] == part {
			return true
		}
	}
	return false
}
