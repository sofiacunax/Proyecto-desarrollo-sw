package services

import (
	"database/sql/driver"
	"errors"
	"testing"

	"proyecto-desarrollo-sw/backend/dtos"
)

func validEventoRequest() dtos.EventoDTO {
	return dtos.EventoDTO{
		Titulo: "Recital", Fecha: "2026-01-01", Horario: "20:00",
		Ubicacion: "Cordoba", Capacidad: 100, Precio: 1500,
	}
}

func TestEventoServiceCrearEventoExitoso(t *testing.T) {
	useDatabaseStub(t, databaseResult{})
	if err := NewEventoService().CrearEvento(validEventoRequest()); err != nil {
		t.Fatalf("error inesperado: %v", err)
	}
}

func TestEventoServiceConsultas(t *testing.T) {
	columns := []string{"id", "titulo", "descripcion", "fecha", "horario", "duracion", "ubicacion", "capacidad", "precio", "categoria", "imagen_url", "estado"}
	row := []driver.Value{int64(1), "Recital", "Musica", "2026-01-01", "20:00", int64(120), "Cordoba", int64(100), 1500.0, "Musica", "img", "ACTIVO"}

	t.Run("listar", func(t *testing.T) {
		useDatabaseStub(t, databaseResult{columns: columns, rows: [][]driver.Value{row}})
		items, err := NewEventoService().ObtenerEventos()
		if err != nil || len(items) != 1 {
			t.Fatalf("resultado inesperado: %#v, %v", items, err)
		}
	})

	t.Run("por id", func(t *testing.T) {
		useDatabaseStub(t, databaseResult{columns: columns, rows: [][]driver.Value{row}})
		item, err := NewEventoService().ObtenerEventoPorID(1)
		if err != nil || item == nil || item.Titulo != "Recital" {
			t.Fatalf("resultado inesperado: %#v, %v", item, err)
		}
	})
}

func TestEventoServiceActualizarEvento(t *testing.T) {
	columns := []string{"id", "titulo", "descripcion", "fecha", "horario", "duracion", "ubicacion", "capacidad", "precio", "categoria", "imagen_url", "estado"}
	row := []driver.Value{int64(1), "Viejo", "", "2026-01-01", "20:00", int64(60), "Cordoba", int64(10), 10.0, "", "", "ACTIVO"}

	t.Run("inexistente", func(t *testing.T) {
		useDatabaseStub(t, databaseResult{columns: columns})
		if err := NewEventoService().ActualizarEvento(1, validEventoRequest()); err == nil {
			t.Fatal("se esperaba evento no encontrado")
		}
	})

	t.Run("error consulta", func(t *testing.T) {
		useDatabaseStub(t, databaseResult{err: errors.New("fallo consulta")})
		if err := NewEventoService().ActualizarEvento(1, validEventoRequest()); err == nil {
			t.Fatal("se esperaba el error de consulta")
		}
	})

	t.Run("validaciones", func(t *testing.T) {
		requests := []dtos.EventoDTO{
			{Fecha: "x", Horario: "x", Ubicacion: "x", Capacidad: 1},
			{Titulo: "x", Horario: "x", Ubicacion: "x", Capacidad: 1},
			{Titulo: "x", Fecha: "x", Ubicacion: "x", Capacidad: 1},
			{Titulo: "x", Fecha: "x", Horario: "x", Capacidad: 1},
			{Titulo: "x", Fecha: "x", Horario: "x", Ubicacion: "x", Capacidad: 0},
			{Titulo: "x", Fecha: "x", Horario: "x", Ubicacion: "x", Capacidad: 1, Precio: -1},
		}
		results := make([]databaseResult, len(requests))
		for index := range results {
			results[index] = databaseResult{columns: columns, rows: [][]driver.Value{row}}
		}
		useDatabaseStub(t, results...)
		for index, request := range requests {
			if err := NewEventoService().ActualizarEvento(1, request); err == nil {
				t.Fatalf("caso %d: se esperaba error", index)
			}
		}
	})

	t.Run("exitosa", func(t *testing.T) {
		useDatabaseStub(t, databaseResult{columns: columns, rows: [][]driver.Value{row}}, databaseResult{})
		if err := NewEventoService().ActualizarEvento(1, validEventoRequest()); err != nil {
			t.Fatalf("error inesperado: %v", err)
		}
	})
}

func TestEventoServiceEliminarEvento(t *testing.T) {
	columns := []string{"id", "titulo", "descripcion", "fecha", "horario", "duracion", "ubicacion", "capacidad", "precio", "categoria", "imagen_url", "estado"}
	row := []driver.Value{int64(1), "Recital", "", "2026-01-01", "20:00", int64(60), "Cordoba", int64(10), 10.0, "", "", "ACTIVO"}

	t.Run("inexistente", func(t *testing.T) {
		useDatabaseStub(t, databaseResult{columns: columns})
		if err := NewEventoService().EliminarEvento(1); err == nil {
			t.Fatal("se esperaba evento no encontrado")
		}
	})

	t.Run("exitoso", func(t *testing.T) {
		useDatabaseStub(t, databaseResult{columns: columns, rows: [][]driver.Value{row}}, databaseResult{})
		if err := NewEventoService().EliminarEvento(1); err != nil {
			t.Fatalf("error inesperado: %v", err)
		}
	})
}
