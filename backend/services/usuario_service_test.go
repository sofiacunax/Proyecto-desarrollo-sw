package services

import (
	"database/sql/driver"
	"errors"
	"testing"
)

func TestUsuarioServiceObtenerUsuarios(t *testing.T) {
	t.Run("exitoso", func(t *testing.T) {
		useDatabaseStub(t, databaseResult{
			columns: []string{"id", "nombre", "email", "password_hash", "rol"},
			rows: [][]driver.Value{
				{int64(1), "Ana", "ana@example.com", "hash", "CLIENTE"},
			},
		})

		usuarios, err := NewUsuarioService().ObtenerUsuarios()
		if err != nil {
			t.Fatalf("error inesperado: %v", err)
		}
		if len(usuarios) != 1 || usuarios[0].Email != "ana@example.com" {
			t.Fatalf("usuarios inesperados: %#v", usuarios)
		}
	})

	t.Run("error dao", func(t *testing.T) {
		useDatabaseStub(t, databaseResult{err: errors.New("fallo usuarios")})

		usuarios, err := NewUsuarioService().ObtenerUsuarios()
		if err == nil || !stringsContains(err.Error(), "fallo usuarios") {
			t.Fatalf("se esperaba error de DAO, obtenido %#v, %v", usuarios, err)
		}
	})
}

func TestUsuarioServiceCambiarRol(t *testing.T) {
	tests := []struct {
		name    string
		rol     string
		results []databaseResult
		wantErr string
	}{
		{name: "rol invalido", rol: "SUPERADMIN", wantErr: "rol"},
		{name: "admin exitoso", rol: "ADMIN", results: []databaseResult{{}}},
		{name: "cliente exitoso", rol: "CLIENTE", results: []databaseResult{{}}},
		{name: "error persistencia", rol: "ADMIN", results: []databaseResult{{err: errors.New("fallo update")}}, wantErr: "fallo update"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if len(test.results) > 0 {
				useDatabaseStub(t, test.results...)
			}

			err := NewUsuarioService().CambiarRol(1, test.rol)
			assertServiceError(t, err, test.wantErr)
		})
	}
}

func TestEntradaServiceObtenerReporteEvento(t *testing.T) {
	t.Run("evento existente sin compradores", func(t *testing.T) {
		useDatabaseStub(t,
			eventoReporteResult(10, "Concierto", 100),
			databaseResult{columns: []string{"cantidad"}, rows: [][]driver.Value{{int64(0)}}},
			databaseResult{columns: []string{"id", "nombre", "email"}},
		)

		reporte, err := NewEntradaService().ObtenerReporteEvento(10)
		if err != nil {
			t.Fatalf("error inesperado: %v", err)
		}
		if reporte.EventoID != 10 || reporte.EntradasVendidas != 0 || reporte.PorcentajeOcupado != 0 || len(reporte.Compradores) != 0 {
			t.Fatalf("reporte inesperado: %#v", reporte)
		}
	})

	t.Run("evento existente con compradores", func(t *testing.T) {
		useDatabaseStub(t,
			eventoReporteResult(10, "Concierto", 4),
			databaseResult{columns: []string{"cantidad"}, rows: [][]driver.Value{{int64(2)}}},
			databaseResult{
				columns: []string{"id", "nombre", "email"},
				rows: [][]driver.Value{
					{int64(1), "Ana", "ana@example.com"},
					{int64(2), "Bruno", "bruno@example.com"},
				},
			},
		)

		reporte, err := NewEntradaService().ObtenerReporteEvento(10)
		if err != nil {
			t.Fatalf("error inesperado: %v", err)
		}
		if reporte.EntradasVendidas != 2 || reporte.PorcentajeOcupado != 50 || len(reporte.Compradores) != 2 {
			t.Fatalf("reporte inesperado: %#v", reporte)
		}
	})

	t.Run("evento con capacidad cero", func(t *testing.T) {
		useDatabaseStub(t,
			eventoReporteResult(10, "Concierto", 0),
			databaseResult{columns: []string{"cantidad"}, rows: [][]driver.Value{{int64(0)}}},
			databaseResult{columns: []string{"id", "nombre", "email"}},
		)

		reporte, err := NewEntradaService().ObtenerReporteEvento(10)
		if err != nil {
			t.Fatalf("error inesperado: %v", err)
		}
		if reporte.PorcentajeOcupado != 0 || reporte.Capacidad != 0 {
			t.Fatalf("reporte inesperado para capacidad cero: %#v", reporte)
		}
	})

	t.Run("evento inexistente", func(t *testing.T) {
		useDatabaseStub(t, databaseResult{columns: eventoReporteColumns()})

		reporte, err := NewEntradaService().ObtenerReporteEvento(99)
		if err == nil || !stringsContains(err.Error(), "evento no encontrado") || reporte != nil {
			t.Fatalf("resultado inesperado: %#v, %v", reporte, err)
		}
	})

	t.Run("error consulta evento", func(t *testing.T) {
		useDatabaseStub(t, databaseResult{err: errors.New("fallo evento")})

		reporte, err := NewEntradaService().ObtenerReporteEvento(10)
		if err == nil || !stringsContains(err.Error(), "fallo evento") || reporte != nil {
			t.Fatalf("resultado inesperado: %#v, %v", reporte, err)
		}
	})

	t.Run("error al contar entradas", func(t *testing.T) {
		useDatabaseStub(t,
			eventoReporteResult(10, "Concierto", 100),
			databaseResult{err: errors.New("fallo conteo")},
		)

		reporte, err := NewEntradaService().ObtenerReporteEvento(10)
		if err == nil || !stringsContains(err.Error(), "fallo conteo") || reporte != nil {
			t.Fatalf("resultado inesperado: %#v, %v", reporte, err)
		}
	})

	t.Run("error al obtener compradores", func(t *testing.T) {
		useDatabaseStub(t,
			eventoReporteResult(10, "Concierto", 100),
			databaseResult{columns: []string{"cantidad"}, rows: [][]driver.Value{{int64(1)}}},
			databaseResult{err: errors.New("fallo compradores")},
		)

		reporte, err := NewEntradaService().ObtenerReporteEvento(10)
		if err == nil || !stringsContains(err.Error(), "fallo compradores") || reporte != nil {
			t.Fatalf("resultado inesperado: %#v, %v", reporte, err)
		}
	})
}

func eventoReporteResult(id int, titulo string, capacidad int) databaseResult {
	return databaseResult{
		columns: eventoReporteColumns(),
		rows: [][]driver.Value{{
			int64(id), titulo, "Desc", "2026-01-01", "20:00", int64(120), "Cordoba", int64(capacidad), 1000.0, "Musica", "img", "ACTIVO",
		}},
	}
}

func eventoReporteColumns() []string {
	return []string{"id", "titulo", "descripcion", "fecha", "horario", "duracion", "ubicacion", "capacidad", "precio", "categoria", "imagen_url", "estado"}
}
