package dao

import (
	"database/sql/driver"
	"errors"
	"testing"

	"proyecto-desarrollo-sw/backend/models"
)

func TestEventoDAO(t *testing.T) {
	dao := NewEventoDAO()

	t.Run("crear exitoso", func(t *testing.T) {
		useDaoDB(t, daoDBResult{})
		if err := dao.CrearEvento(models.Evento{Titulo: "Recital"}); err != nil {
			t.Fatalf("error inesperado: %v", err)
		}
	})

	t.Run("listar con busqueda", func(t *testing.T) {
		useDaoDB(t, daoDBResult{columns: eventoDAOColumns(), rows: [][]driver.Value{eventoDAORow()}})
		eventos, err := dao.ObtenerEventos("recital")
		if err != nil || len(eventos) != 1 || eventos[0].Titulo != "Recital" {
			t.Fatalf("resultado inesperado: %#v, %v", eventos, err)
		}
	})

	t.Run("listar error", func(t *testing.T) {
		useDaoDB(t, daoDBResult{err: errors.New("fallo listar")})
		eventos, err := dao.ObtenerEventos("")
		assertDAOError(t, err, "fallo listar")
		if eventos != nil {
			t.Fatalf("se esperaba nil, obtenido %#v", eventos)
		}
	})

	t.Run("obtener por id exitoso", func(t *testing.T) {
		useDaoDB(t, daoDBResult{columns: eventoDAOColumns(), rows: [][]driver.Value{eventoDAORow()}})
		evento, err := dao.ObtenerEventoPorID(1)
		if err != nil || evento == nil || evento.ID != 1 {
			t.Fatalf("resultado inesperado: %#v, %v", evento, err)
		}
	})

	t.Run("obtener por id inexistente", func(t *testing.T) {
		useDaoDB(t, daoDBResult{columns: eventoDAOColumns()})
		evento, err := dao.ObtenerEventoPorID(99)
		if err != nil || evento != nil {
			t.Fatalf("resultado inesperado: %#v, %v", evento, err)
		}
	})

	t.Run("obtener por id error", func(t *testing.T) {
		useDaoDB(t, daoDBResult{err: errors.New("fallo obtener")})
		evento, err := dao.ObtenerEventoPorID(99)
		assertDAOError(t, err, "fallo obtener")
		if evento != nil {
			t.Fatalf("se esperaba nil, obtenido %#v", evento)
		}
	})

	t.Run("actualizar exitoso", func(t *testing.T) {
		useDaoDB(t, daoDBResult{})
		if err := dao.ActualizarEvento(1, models.Evento{Titulo: "Nuevo", Estado: "ACTIVO"}); err != nil {
			t.Fatalf("error inesperado: %v", err)
		}
	})

	t.Run("actualizar error", func(t *testing.T) {
		useDaoDB(t, daoDBResult{err: errors.New("fallo update")})
		assertDAOError(t, dao.ActualizarEvento(1, models.Evento{Titulo: "Nuevo"}), "fallo update")
	})

	t.Run("eliminar exitoso", func(t *testing.T) {
		useDaoDB(t, daoDBResult{})
		if err := dao.EliminarEvento(1); err != nil {
			t.Fatalf("error inesperado: %v", err)
		}
	})

	t.Run("eliminar error", func(t *testing.T) {
		useDaoDB(t, daoDBResult{err: errors.New("fallo delete")})
		assertDAOError(t, dao.EliminarEvento(1), "fallo delete")
	})
}

func TestEntradaDAO(t *testing.T) {
	dao := NewEntradaDAO()

	t.Run("crear exitoso", func(t *testing.T) {
		useDaoDB(t, daoDBResult{})
		if err := dao.CrearEntrada(models.Entrada{UsuarioID: 1, EventoID: 2, Estado: "ACTIVA"}); err != nil {
			t.Fatalf("error inesperado: %v", err)
		}
	})

	t.Run("obtener por usuario", func(t *testing.T) {
		useDaoDB(t, daoDBResult{columns: entradaDAOColumns(), rows: [][]driver.Value{{int64(1), int64(1), int64(2), "ACTIVA"}}})
		entradas, err := dao.ObtenerPorUsuario(1)
		if err != nil || len(entradas) != 1 {
			t.Fatalf("resultado inesperado: %#v, %v", entradas, err)
		}
	})

	t.Run("obtener por usuario error", func(t *testing.T) {
		useDaoDB(t, daoDBResult{err: errors.New("fallo usuario")})
		entradas, err := dao.ObtenerPorUsuario(1)
		assertDAOError(t, err, "fallo usuario")
		if entradas != nil {
			t.Fatalf("se esperaba nil, obtenido %#v", entradas)
		}
	})

	t.Run("obtener mis entradas dto", func(t *testing.T) {
		useDaoDB(t, daoDBResult{
			columns: []string{"id", "evento_id", "titulo", "fecha", "ubicacion", "estado"},
			rows:    [][]driver.Value{{int64(1), int64(2), "Recital", "2026-01-01", "Cordoba", "ACTIVA"}},
		})
		entradas, err := dao.ObtenerMisEntradasDTO(1)
		if err != nil || len(entradas) != 1 || entradas[0].Titulo != "Recital" {
			t.Fatalf("resultado inesperado: %#v, %v", entradas, err)
		}
	})

	t.Run("obtener mis entradas dto error", func(t *testing.T) {
		useDaoDB(t, daoDBResult{err: errors.New("fallo dto")})
		entradas, err := dao.ObtenerMisEntradasDTO(1)
		assertDAOError(t, err, "fallo dto")
		if entradas != nil {
			t.Fatalf("se esperaba nil, obtenido %#v", entradas)
		}
	})

	t.Run("obtener por id inexistente", func(t *testing.T) {
		useDaoDB(t, daoDBResult{columns: entradaDAOColumns()})
		entrada, err := dao.ObtenerPorID(99)
		if err != nil || entrada != nil {
			t.Fatalf("resultado inesperado: %#v, %v", entrada, err)
		}
	})

	t.Run("obtener por id exitoso", func(t *testing.T) {
		useDaoDB(t, daoDBResult{
			columns: entradaDAOColumns(),
			rows:    [][]driver.Value{{int64(1), int64(2), int64(3), "ACTIVA"}},
		})
		entrada, err := dao.ObtenerPorID(1)
		if err != nil || entrada == nil || entrada.ID != 1 {
			t.Fatalf("resultado inesperado: %#v, %v", entrada, err)
		}
	})

	t.Run("obtener por id error", func(t *testing.T) {
		useDaoDB(t, daoDBResult{err: errors.New("fallo entrada")})
		entrada, err := dao.ObtenerPorID(99)
		assertDAOError(t, err, "fallo entrada")
		if entrada != nil {
			t.Fatalf("se esperaba nil, obtenido %#v", entrada)
		}
	})

	t.Run("cancelar y transferir", func(t *testing.T) {
		useDaoDB(t, daoDBResult{}, daoDBResult{})
		if err := dao.CancelarEntrada(1); err != nil {
			t.Fatalf("error al cancelar: %v", err)
		}
		if err := dao.TransferirEntrada(1, 3); err != nil {
			t.Fatalf("error al transferir: %v", err)
		}
	})

	t.Run("capacidad y ocupacion", func(t *testing.T) {
		useDaoDB(t,
			daoDBResult{columns: []string{"cantidad"}, rows: [][]driver.Value{{int64(3)}}},
			daoDBResult{columns: []string{"capacidad"}, rows: [][]driver.Value{{int64(10)}}},
		)
		ocupados, err := dao.ContarEntradasActivas(2)
		if err != nil || ocupados != 3 {
			t.Fatalf("ocupacion inesperada: %d, %v", ocupados, err)
		}
		capacidad, err := dao.ObtenerCapacidadEvento(2)
		if err != nil || capacidad != 10 {
			t.Fatalf("capacidad inesperada: %d, %v", capacidad, err)
		}
	})

	t.Run("capacidad error", func(t *testing.T) {
		useDaoDB(t, daoDBResult{err: errors.New("fallo capacidad")})
		capacidad, err := dao.ObtenerCapacidadEvento(2)
		assertDAOError(t, err, "fallo capacidad")
		if capacidad != 0 {
			t.Fatalf("se esperaba capacidad cero ante error, obtenido %d", capacidad)
		}
	})

	t.Run("compradores por evento", func(t *testing.T) {
		useDaoDB(t, daoDBResult{
			columns: []string{"id", "nombre", "email"},
			rows:    [][]driver.Value{{int64(1), "Ana", "ana@example.com"}},
		})
		compradores, err := dao.ObtenerCompradoresPorEvento(2)
		if err != nil || len(compradores) != 1 || compradores[0].Email != "ana@example.com" {
			t.Fatalf("compradores inesperados: %#v, %v", compradores, err)
		}
	})
}

func TestPuntuacionDAO(t *testing.T) {
	dao := NewPuntuacionDAO()

	t.Run("crear exitoso", func(t *testing.T) {
		useDaoDB(t, daoDBResult{})
		if err := dao.CrearPuntuacion(models.Puntuacion{UsuarioID: 1, EventoID: 2, Puntuacion: 5}); err != nil {
			t.Fatalf("error inesperado: %v", err)
		}
	})

	t.Run("listar por evento", func(t *testing.T) {
		useDaoDB(t, daoDBResult{columns: []string{"id", "usuario_id", "evento_id", "puntuacion"}, rows: [][]driver.Value{{int64(1), int64(1), int64(2), int64(5)}}})
		puntuaciones, err := dao.ObtenerPuntuacionesPorEvento(2)
		if err != nil || len(puntuaciones) != 1 || puntuaciones[0].Puntuacion != 5 {
			t.Fatalf("resultado inesperado: %#v, %v", puntuaciones, err)
		}
	})

	t.Run("listar por evento error", func(t *testing.T) {
		useDaoDB(t, daoDBResult{err: errors.New("fallo listar")})
		puntuaciones, err := dao.ObtenerPuntuacionesPorEvento(2)
		assertDAOError(t, err, "fallo listar")
		if puntuaciones != nil {
			t.Fatalf("se esperaba nil, obtenido %#v", puntuaciones)
		}
	})

	t.Run("promedio", func(t *testing.T) {
		useDaoDB(t, daoDBResult{columns: []string{"promedio"}, rows: [][]driver.Value{{4.5}}})
		promedio, err := dao.ObtenerPromedioPorEvento(2)
		if err != nil || promedio != 4.5 {
			t.Fatalf("promedio inesperado: %f, %v", promedio, err)
		}
	})

	t.Run("ranking", func(t *testing.T) {
		useDaoDB(t, daoDBResult{columns: []string{"evento_id", "titulo", "promedio"}, rows: [][]driver.Value{{int64(2), "Recital", 4.5}}})
		ranking, err := dao.ObtenerRankingEventos()
		if err != nil || len(ranking) != 1 || ranking[0].Titulo != "Recital" {
			t.Fatalf("ranking inesperado: %#v, %v", ranking, err)
		}
	})

	t.Run("ranking error", func(t *testing.T) {
		useDaoDB(t, daoDBResult{err: errors.New("fallo ranking")})
		ranking, err := dao.ObtenerRankingEventos()
		assertDAOError(t, err, "fallo ranking")
		if ranking != nil {
			t.Fatalf("se esperaba nil, obtenido %#v", ranking)
		}
	})
}

func TestUsuarioDAO(t *testing.T) {
	dao := NewUsuarioDAO()

	t.Run("crear exitoso", func(t *testing.T) {
		useDaoDB(t, daoDBResult{})
		if err := dao.CrearUsuario(models.Usuario{Nombre: "Ana", Email: "ana@example.com", PasswordHash: "hash", Rol: "CLIENTE"}); err != nil {
			t.Fatalf("error inesperado: %v", err)
		}
	})

	t.Run("buscar por email", func(t *testing.T) {
		useDaoDB(t, usuarioDAOResult(1, "Ana", "ana@example.com", "CLIENTE"))
		usuario, err := dao.BuscarPorEmail("ana@example.com")
		if err != nil || usuario == nil || usuario.Email != "ana@example.com" {
			t.Fatalf("usuario inesperado: %#v, %v", usuario, err)
		}
	})

	t.Run("buscar por email error", func(t *testing.T) {
		useDaoDB(t, daoDBResult{err: errors.New("fallo buscar")})
		usuario, err := dao.BuscarPorEmail("ana@example.com")
		assertDAOError(t, err, "fallo buscar")
		if usuario != nil {
			t.Fatalf("se esperaba nil, obtenido %#v", usuario)
		}
	})

	t.Run("obtener usuarios", func(t *testing.T) {
		useDaoDB(t, usuarioDAOResult(1, "Ana", "ana@example.com", "CLIENTE"))
		usuarios, err := dao.ObtenerUsuarios()
		if err != nil || len(usuarios) != 1 {
			t.Fatalf("usuarios inesperados: %#v, %v", usuarios, err)
		}
	})

	t.Run("actualizar rol", func(t *testing.T) {
		useDaoDB(t, daoDBResult{})
		if err := dao.ActualizarRol(1, "ADMIN"); err != nil {
			t.Fatalf("error inesperado: %v", err)
		}
	})
}

func eventoDAOColumns() []string {
	return []string{"id", "titulo", "descripcion", "fecha", "horario", "duracion", "ubicacion", "capacidad", "precio", "categoria", "imagen_url", "estado"}
}

func eventoDAORow() []driver.Value {
	return []driver.Value{int64(1), "Recital", "Musica", "2026-01-01", "20:00", int64(120), "Cordoba", int64(100), 1000.0, "Musica", "img", "ACTIVO"}
}

func entradaDAOColumns() []string {
	return []string{"id", "usuario_id", "evento_id", "estado"}
}

func usuarioDAOResult(id int, nombre string, email string, rol string) daoDBResult {
	return daoDBResult{
		columns: []string{"id", "nombre", "email", "password_hash", "rol"},
		rows:    [][]driver.Value{{int64(id), nombre, email, "hash", rol}},
	}
}

func assertDAOError(t *testing.T, err error, want string) {
	t.Helper()
	if err == nil || !containsText(err.Error(), want) {
		t.Fatalf("se esperaba error %q, se obtuvo %v", want, err)
	}
}

func containsText(value, part string) bool {
	for i := 0; i+len(part) <= len(value); i++ {
		if value[i:i+len(part)] == part {
			return true
		}
	}
	return false
}
