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

func TestEventoServiceCrearEventoConEstadoExplicito(t *testing.T) {
	useDatabaseStub(t, databaseResult{})
	request := validEventoRequest()
	request.Estado = "CANCELADO"
	if err := NewEventoService().CrearEvento(request); err != nil {
		t.Fatalf("error inesperado: %v", err)
	}
}

func TestEventoServiceConsultas(t *testing.T) {
	columns := []string{"id", "titulo", "descripcion", "fecha", "horario", "duracion", "ubicacion", "capacidad", "precio", "categoria", "imagen_url", "estado"}
	row := []driver.Value{int64(1), "Recital", "Musica", "2026-01-01", "20:00", int64(120), "Cordoba", int64(100), 1500.0, "Musica", "img", "ACTIVO"}

	t.Run("listar", func(t *testing.T) {
		useDatabaseStub(t, databaseResult{columns: columns, rows: [][]driver.Value{row}})
		items, err := NewEventoService().ObtenerEventos("")
		if err != nil || len(items) != 1 {
			t.Fatalf("resultado inesperado: %#v, %v", items, err)
		}
	})

	t.Run("buscar por texto", func(t *testing.T) {
		useDatabaseStub(t, databaseResult{columns: columns, rows: [][]driver.Value{row}})
		items, err := NewEventoService().ObtenerEventos("Recital")
		if err != nil || len(items) != 1 || items[0].Titulo != "Recital" {
			t.Fatalf("resultado de busqueda inesperado: %#v, %v", items, err)
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

	t.Run("exitosa con estado explicito", func(t *testing.T) {
		request := validEventoRequest()
		request.Estado = "CANCELADO"
		useDatabaseStub(t, databaseResult{columns: columns, rows: [][]driver.Value{row}}, databaseResult{})
		if err := NewEventoService().ActualizarEvento(1, request); err != nil {
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
		useDatabaseStub(t,
			databaseResult{columns: columns, rows: [][]driver.Value{row}},
			databaseResult{columns: []string{"count"}, rows: [][]driver.Value{{int64(0)}}},
			databaseResult{columns: []string{"count"}, rows: [][]driver.Value{{int64(0)}}},
			databaseResult{},
		)
		if err := NewEventoService().EliminarEvento(1); err != nil {
			t.Fatalf("error inesperado: %v", err)
		}
	})

	t.Run("error consulta", func(t *testing.T) {
		useDatabaseStub(t, databaseResult{err: errors.New("fallo consulta")})
		if err := NewEventoService().EliminarEvento(1); err == nil {
			t.Fatal("se esperaba error de consulta")
		}
	})

	t.Run("error eliminacion", func(t *testing.T) {
		useDatabaseStub(t,
			databaseResult{columns: columns, rows: [][]driver.Value{row}},
			databaseResult{columns: []string{"count"}, rows: [][]driver.Value{{int64(0)}}},
			databaseResult{columns: []string{"count"}, rows: [][]driver.Value{{int64(0)}}},
			databaseResult{err: errors.New("fallo eliminacion")},
		)
		if err := NewEventoService().EliminarEvento(1); err == nil {
			t.Fatal("se esperaba error de eliminacion")
		}
	})

	t.Run("con actividad por entradas", func(t *testing.T) {
		useDatabaseStub(t,
			databaseResult{columns: columns, rows: [][]driver.Value{row}},
			databaseResult{columns: []string{"count"}, rows: [][]driver.Value{{int64(1)}}},
		)
		err := NewEventoService().EliminarEvento(1)
		if err == nil || err.Error() != MensajeEventoConActividad {
			t.Fatalf("error inesperado: %v", err)
		}
	})

	t.Run("con actividad por puntuaciones", func(t *testing.T) {
		useDatabaseStub(t,
			databaseResult{columns: columns, rows: [][]driver.Value{row}},
			databaseResult{columns: []string{"count"}, rows: [][]driver.Value{{int64(0)}}},
			databaseResult{columns: []string{"count"}, rows: [][]driver.Value{{int64(2)}}},
		)
		err := NewEventoService().EliminarEvento(1)
		if err == nil || err.Error() != MensajeEventoConActividad {
			t.Fatalf("error inesperado: %v", err)
		}
	})
}

func TestEventoServiceCambiarEstadoEvento(t *testing.T) {
	columns := []string{"id", "titulo", "descripcion", "fecha", "horario", "duracion", "ubicacion", "capacidad", "precio", "categoria", "imagen_url", "estado"}
	rowActivo := []driver.Value{int64(1), "Recital", "", "2026-01-01", "20:00", int64(60), "Cordoba", int64(10), 10.0, "", "", "ACTIVO"}
	rowCancelado := []driver.Value{int64(1), "Recital", "", "2026-01-01", "20:00", int64(60), "Cordoba", int64(10), 10.0, "", "", "CANCELADO"}

	t.Run("activo a cancelado", func(t *testing.T) {
		useDatabaseStub(t, databaseResult{columns: columns, rows: [][]driver.Value{rowActivo}}, databaseResult{})
		if err := NewEventoService().CambiarEstadoEvento(1, "CANCELADO"); err != nil {
			t.Fatalf("error inesperado: %v", err)
		}
	})

	t.Run("cancelado a activo", func(t *testing.T) {
		useDatabaseStub(t, databaseResult{columns: columns, rows: [][]driver.Value{rowCancelado}}, databaseResult{})
		if err := NewEventoService().CambiarEstadoEvento(1, "ACTIVO"); err != nil {
			t.Fatalf("error inesperado: %v", err)
		}
	})

	t.Run("estado invalido", func(t *testing.T) {
		useDatabaseStub(t, databaseResult{columns: columns, rows: [][]driver.Value{rowActivo}})
		if err := NewEventoService().CambiarEstadoEvento(1, "PAUSADO"); err == nil {
			t.Fatal("se esperaba error por estado invalido")
		}
	})

	t.Run("mismo estado", func(t *testing.T) {
		useDatabaseStub(t, databaseResult{columns: columns, rows: [][]driver.Value{rowActivo}})
		if err := NewEventoService().CambiarEstadoEvento(1, "ACTIVO"); err == nil {
			t.Fatal("se esperaba error por mismo estado")
		}
	})

	t.Run("inexistente", func(t *testing.T) {
		useDatabaseStub(t, databaseResult{columns: columns})
		if err := NewEventoService().CambiarEstadoEvento(1, "ACTIVO"); err == nil {
			t.Fatal("se esperaba evento no encontrado")
		}
	})
}
