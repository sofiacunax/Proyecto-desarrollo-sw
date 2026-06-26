package controllers

import (
	"database/sql/driver"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"proyecto-desarrollo-sw/backend/services"
	"proyecto-desarrollo-sw/backend/utils"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func performRequest(method, path, body string, handler gin.HandlerFunc, userID any) *httptest.ResponseRecorder {
	router := gin.New()
	if userID != nil {
		router.Use(func(ctx *gin.Context) {
			ctx.Set("userID", userID)
			ctx.Next()
		})
	}
	router.Handle(method, path, handler)
	request := httptest.NewRequest(method, path, strings.NewReader(body))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	return response
}

func assertStatus(t *testing.T, response *httptest.ResponseRecorder, expected int) {
	t.Helper()
	if response.Code != expected {
		t.Fatalf("se esperaba status %d, se obtuvo %d: %s", expected, response.Code, response.Body.String())
	}
}

func TestAuthControllerRegister(t *testing.T) {
	controller := &AuthController{AuthService: services.AuthService{}}

	t.Run("json invalido", func(t *testing.T) {
		assertStatus(t, performRequest(http.MethodPost, "/auth/register", "{", controller.Register, nil), http.StatusBadRequest)
	})

	t.Run("exitoso", func(t *testing.T) {
		useControllerDB(t, controllerDBResult{})
		body := `{"nombre":"Ana","email":"ana@example.com","password":"secreto"}`
		assertStatus(t, performRequest(http.MethodPost, "/auth/register", body, controller.Register, nil), http.StatusCreated)
	})

	t.Run("error persistencia", func(t *testing.T) {
		useControllerDB(t, controllerDBResult{err: errors.New("email duplicado")})
		body := `{"nombre":"Ana","email":"ana@example.com","password":"secreto"}`
		assertStatus(t, performRequest(http.MethodPost, "/auth/register", body, controller.Register, nil), http.StatusInternalServerError)
	})
}

func TestAuthControllerLogin(t *testing.T) {
	controller := &AuthController{AuthService: services.AuthService{}}

	t.Run("json invalido", func(t *testing.T) {
		assertStatus(t, performRequest(http.MethodPost, "/auth/login", "{", controller.Login, nil), http.StatusBadRequest)
	})

	t.Run("credenciales invalidas", func(t *testing.T) {
		useControllerDB(t, controllerDBResult{err: errors.New("sin usuario")})
		assertStatus(t, performRequest(http.MethodPost, "/auth/login", `{"email":"x@y.com","password":"x"}`, controller.Login, nil), http.StatusUnauthorized)
	})

	t.Run("exitoso", func(t *testing.T) {
		hash, _ := bcrypt.GenerateFromPassword([]byte("secreto"), bcrypt.MinCost)
		useControllerDB(t, controllerDBResult{
			columns: []string{"id", "nombre", "email", "password_hash", "rol"},
			rows:    [][]driver.Value{{int64(1), "Ana", "ana@example.com", string(hash), "CLIENTE"}},
		})
		assertStatus(t, performRequest(http.MethodPost, "/auth/login", `{"email":"ana@example.com","password":"secreto"}`, controller.Login, nil), http.StatusOK)
	})
}

func eventoColumns() []string {
	return []string{"id", "titulo", "descripcion", "fecha", "horario", "duracion", "ubicacion", "capacidad", "precio", "categoria", "imagen_url", "estado"}
}

func eventoRow() []driver.Value {
	return []driver.Value{int64(1), "Recital", "Musica", "2026-01-01", "20:00", int64(120), "Cordoba", int64(100), 1000.0, "Musica", "img", "ACTIVO"}
}

func eventoJSON() string {
	return `{"titulo":"Recital","fecha":"2026-01-01","horario":"20:00","ubicacion":"Cordoba","capacidad":100,"precio":1000}`
}

func TestEventoControllerCaminosPrincipales(t *testing.T) {
	controller := NewEventoController()

	t.Run("crear validacion", func(t *testing.T) {
		assertStatus(t, performRequest(http.MethodPost, "/eventos", `{}`, controller.CrearEvento, nil), http.StatusBadRequest)
	})
	t.Run("crear exitoso", func(t *testing.T) {
		useControllerDB(t, controllerDBResult{})
		assertStatus(t, performRequest(http.MethodPost, "/eventos", eventoJSON(), controller.CrearEvento, nil), http.StatusCreated)
	})
	t.Run("listar exitoso", func(t *testing.T) {
		useControllerDB(t, controllerDBResult{columns: eventoColumns(), rows: [][]driver.Value{eventoRow()}})
		assertStatus(t, performRequest(http.MethodGet, "/eventos", "", controller.ObtenerEventos, nil), http.StatusOK)
	})
	t.Run("listar error", func(t *testing.T) {
		useControllerDB(t, controllerDBResult{err: errors.New("db")})
		assertStatus(t, performRequest(http.MethodGet, "/eventos", "", controller.ObtenerEventos, nil), http.StatusInternalServerError)
	})
	t.Run("detalle inexistente", func(t *testing.T) {
		useControllerDB(t, controllerDBResult{columns: eventoColumns()})
		assertStatus(t, performRequest(http.MethodGet, "/eventos/:id", "", controller.ObtenerEventoPorID, nil), http.StatusBadRequest)
	})
	// Las rutas concretas permiten ejercitar las ramas posteriores al parseo del ID.
	t.Run("detalle encontrado", func(t *testing.T) {
		useControllerDB(t, controllerDBResult{columns: eventoColumns(), rows: [][]driver.Value{eventoRow()}})
		assertStatus(t, requestWithParam(http.MethodGet, "/eventos/1", "/eventos/:id", "", controller.ObtenerEventoPorID, nil), http.StatusOK)
	})
	t.Run("detalle no encontrado", func(t *testing.T) {
		useControllerDB(t, controllerDBResult{columns: eventoColumns()})
		assertStatus(t, requestWithParam(http.MethodGet, "/eventos/1", "/eventos/:id", "", controller.ObtenerEventoPorID, nil), http.StatusNotFound)
	})
	t.Run("detalle error interno", func(t *testing.T) {
		useControllerDB(t, controllerDBResult{err: errors.New("fallo detalle")})
		assertStatus(t, requestWithParam(http.MethodGet, "/eventos/1", "/eventos/:id", "", controller.ObtenerEventoPorID, nil), http.StatusInternalServerError)
	})
	t.Run("actualizar id invalido", func(t *testing.T) {
		assertStatus(t, requestWithParam(http.MethodPut, "/eventos/x", "/eventos/:id", eventoJSON(), controller.ActualizarEvento, nil), http.StatusBadRequest)
	})
	t.Run("actualizar json invalido", func(t *testing.T) {
		assertStatus(t, requestWithParam(http.MethodPut, "/eventos/1", "/eventos/:id", "{", controller.ActualizarEvento, nil), http.StatusBadRequest)
	})
	t.Run("actualizar no encontrado", func(t *testing.T) {
		useControllerDB(t, controllerDBResult{columns: eventoColumns()})
		assertStatus(t, requestWithParam(http.MethodPut, "/eventos/1", "/eventos/:id", eventoJSON(), controller.ActualizarEvento, nil), http.StatusNotFound)
	})
	t.Run("actualizar validacion", func(t *testing.T) {
		useControllerDB(t, controllerDBResult{columns: eventoColumns(), rows: [][]driver.Value{eventoRow()}})
		assertStatus(t, requestWithParam(http.MethodPut, "/eventos/1", "/eventos/:id", `{"titulo":"","fecha":"2026-01-01","horario":"20:00","ubicacion":"Cordoba","capacidad":100,"precio":1000}`, controller.ActualizarEvento, nil), http.StatusBadRequest)
	})
	t.Run("actualizar error dao", func(t *testing.T) {
		useControllerDB(t, controllerDBResult{err: errors.New("fallo consulta")})
		assertStatus(t, requestWithParam(http.MethodPut, "/eventos/1", "/eventos/:id", eventoJSON(), controller.ActualizarEvento, nil), http.StatusBadRequest)
	})
	t.Run("actualizar exitoso", func(t *testing.T) {
		useControllerDB(t, controllerDBResult{columns: eventoColumns(), rows: [][]driver.Value{eventoRow()}}, controllerDBResult{})
		assertStatus(t, requestWithParam(http.MethodPut, "/eventos/1", "/eventos/:id", eventoJSON(), controller.ActualizarEvento, nil), http.StatusOK)
	})
	t.Run("eliminar no encontrado", func(t *testing.T) {
		useControllerDB(t, controllerDBResult{columns: eventoColumns()})
		assertStatus(t, requestWithParam(http.MethodDelete, "/eventos/1", "/eventos/:id", "", controller.EliminarEvento, nil), http.StatusNotFound)
	})
	t.Run("eliminar error consulta", func(t *testing.T) {
		useControllerDB(t, controllerDBResult{err: errors.New("fallo consulta")})
		assertStatus(t, requestWithParam(http.MethodDelete, "/eventos/1", "/eventos/:id", "", controller.EliminarEvento, nil), http.StatusBadRequest)
	})
	t.Run("eliminar error dao", func(t *testing.T) {
		useControllerDB(t,
			controllerDBResult{columns: eventoColumns(), rows: [][]driver.Value{eventoRow()}},
			controllerDBResult{columns: []string{"count"}, rows: [][]driver.Value{{int64(0)}}},
			controllerDBResult{columns: []string{"count"}, rows: [][]driver.Value{{int64(0)}}},
			controllerDBResult{err: errors.New("fallo delete")},
		)
		assertStatus(t, requestWithParam(http.MethodDelete, "/eventos/1", "/eventos/:id", "", controller.EliminarEvento, nil), http.StatusBadRequest)
	})
	t.Run("eliminar con actividad", func(t *testing.T) {
		useControllerDB(t,
			controllerDBResult{columns: eventoColumns(), rows: [][]driver.Value{eventoRow()}},
			controllerDBResult{columns: []string{"count"}, rows: [][]driver.Value{{int64(1)}}},
		)
		assertStatus(t, requestWithParam(http.MethodDelete, "/eventos/1", "/eventos/:id", "", controller.EliminarEvento, nil), http.StatusBadRequest)
	})
	t.Run("eliminar exitoso", func(t *testing.T) {
		useControllerDB(t,
			controllerDBResult{columns: eventoColumns(), rows: [][]driver.Value{eventoRow()}},
			controllerDBResult{columns: []string{"count"}, rows: [][]driver.Value{{int64(0)}}},
			controllerDBResult{columns: []string{"count"}, rows: [][]driver.Value{{int64(0)}}},
			controllerDBResult{},
		)
		assertStatus(t, requestWithParam(http.MethodDelete, "/eventos/1", "/eventos/:id", "", controller.EliminarEvento, nil), http.StatusOK)
	})
	t.Run("cambiar estado id invalido", func(t *testing.T) {
		assertStatus(t, requestWithParam(http.MethodPut, "/eventos/x/estado", "/eventos/:id/estado", `{"estado":"CANCELADO"}`, controller.CambiarEstadoEvento, nil), http.StatusBadRequest)
	})
	t.Run("cambiar estado json invalido", func(t *testing.T) {
		assertStatus(t, requestWithParam(http.MethodPut, "/eventos/1/estado", "/eventos/:id/estado", "{", controller.CambiarEstadoEvento, nil), http.StatusBadRequest)
	})
	t.Run("cambiar estado no encontrado", func(t *testing.T) {
		useControllerDB(t, controllerDBResult{columns: eventoColumns()})
		assertStatus(t, requestWithParam(http.MethodPut, "/eventos/1/estado", "/eventos/:id/estado", `{"estado":"CANCELADO"}`, controller.CambiarEstadoEvento, nil), http.StatusNotFound)
	})
	t.Run("cambiar estado invalido", func(t *testing.T) {
		useControllerDB(t, controllerDBResult{columns: eventoColumns(), rows: [][]driver.Value{eventoRow()}})
		assertStatus(t, requestWithParam(http.MethodPut, "/eventos/1/estado", "/eventos/:id/estado", `{"estado":"PAUSADO"}`, controller.CambiarEstadoEvento, nil), http.StatusBadRequest)
	})
	t.Run("cambiar estado exitoso", func(t *testing.T) {
		useControllerDB(t, controllerDBResult{columns: eventoColumns(), rows: [][]driver.Value{eventoRow()}}, controllerDBResult{})
		assertStatus(t, requestWithParam(http.MethodPut, "/eventos/1/estado", "/eventos/:id/estado", `{"estado":"CANCELADO"}`, controller.CambiarEstadoEvento, nil), http.StatusOK)
	})
}

func TestEventoAdminEstadoConMiddlewares(t *testing.T) {
	controller := NewEventoController()

	router := gin.New()
	private := router.Group("/private")
	private.Use(utils.AuthMiddleware())
	admin := private.Group("/admin")
	admin.Use(utils.AdminMiddleware())
	admin.PUT("/eventos/:id/estado", controller.CambiarEstadoEvento)

	t.Run("sin token", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPut, "/private/admin/eventos/1/estado", strings.NewReader(`{"estado":"CANCELADO"}`))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)
		assertStatus(t, response, http.StatusUnauthorized)
	})

	t.Run("cliente sin permisos", func(t *testing.T) {
		token, err := utils.GenerateToken(2, "cliente@example.com", "CLIENTE")
		if err != nil {
			t.Fatalf("no se pudo generar token: %v", err)
		}
		request := httptest.NewRequest(http.MethodPut, "/private/admin/eventos/1/estado", strings.NewReader(`{"estado":"CANCELADO"}`))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Authorization", "Bearer "+token)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)
		assertStatus(t, response, http.StatusForbidden)
	})

	t.Run("admin permitido", func(t *testing.T) {
		token, err := utils.GenerateToken(1, "admin@example.com", "ADMIN")
		if err != nil {
			t.Fatalf("no se pudo generar token: %v", err)
		}
		useControllerDB(t, controllerDBResult{columns: eventoColumns(), rows: [][]driver.Value{eventoRow()}}, controllerDBResult{})
		request := httptest.NewRequest(http.MethodPut, "/private/admin/eventos/1/estado", strings.NewReader(`{"estado":"CANCELADO"}`))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Authorization", "Bearer "+token)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)
		assertStatus(t, response, http.StatusOK)
	})
}

func requestWithParam(method, path, route, body string, handler gin.HandlerFunc, userID any) *httptest.ResponseRecorder {
	router := gin.New()
	if userID != nil {
		router.Use(func(ctx *gin.Context) { ctx.Set("userID", userID); ctx.Next() })
	}
	router.Handle(method, route, handler)
	request := httptest.NewRequest(method, path, strings.NewReader(body))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	return response
}

func TestEntradaController(t *testing.T) {
	controller := NewEntradaController()

	t.Run("comprar json invalido", func(t *testing.T) {
		assertStatus(t, performRequest(http.MethodPost, "/entradas", "{", controller.ComprarEntrada, float64(1)), http.StatusBadRequest)
	})
	t.Run("comprar evento invalido", func(t *testing.T) {
		assertStatus(t, performRequest(http.MethodPost, "/entradas", `{"evento_id":0}`, controller.ComprarEntrada, float64(1)), http.StatusBadRequest)
	})
	t.Run("comprar exitoso", func(t *testing.T) {
		useControllerDB(t,
			controllerDBResult{columns: eventoColumns(), rows: [][]driver.Value{eventoRow()}},
			controllerDBResult{columns: []string{"cantidad"}, rows: [][]driver.Value{{int64(2)}}},
			controllerDBResult{},
		)
		assertStatus(t, performRequest(http.MethodPost, "/entradas", `{"evento_id":2}`, controller.ComprarEntrada, float64(1)), http.StatusCreated)
	})
	t.Run("comprar evento cancelado", func(t *testing.T) {
		row := eventoRow()
		row[11] = "CANCELADO"
		useControllerDB(t, controllerDBResult{columns: eventoColumns(), rows: [][]driver.Value{row}})
		assertStatus(t, performRequest(http.MethodPost, "/entradas", `{"evento_id":2}`, controller.ComprarEntrada, float64(1)), http.StatusBadRequest)
	})
	t.Run("mis entradas", func(t *testing.T) {
		useControllerDB(t, controllerDBResult{
			columns: []string{"id", "evento_id", "titulo", "fecha", "ubicacion", "estado"},
			rows:    [][]driver.Value{{int64(1), int64(2), "Recital", "2026-01-01", "Cordoba", "ACTIVA"}},
		})
		assertStatus(t, performRequest(http.MethodGet, "/entradas", "", controller.ObtenerMisEntradas, float64(1)), http.StatusOK)
	})
	t.Run("mis entradas error dao", func(t *testing.T) {
		useControllerDB(t, controllerDBResult{err: errors.New("fallo entradas")})
		assertStatus(t, performRequest(http.MethodGet, "/entradas", "", controller.ObtenerMisEntradas, float64(1)), http.StatusInternalServerError)
	})
	t.Run("cancelar id invalido", func(t *testing.T) {
		assertStatus(t, requestWithParam(http.MethodPut, "/entradas/x", "/entradas/:id", "", controller.CancelarEntrada, float64(1)), http.StatusBadRequest)
	})
	t.Run("cancelar no encontrada", func(t *testing.T) {
		useControllerDB(t, controllerDBResult{columns: []string{"id", "usuario_id", "evento_id", "estado"}})
		assertStatus(t, requestWithParam(http.MethodPut, "/entradas/1", "/entradas/:id", "", controller.CancelarEntrada, float64(1)), http.StatusBadRequest)
	})
	t.Run("cancelar exitoso", func(t *testing.T) {
		useControllerDB(t,
			controllerDBResult{columns: []string{"id", "usuario_id", "evento_id", "estado"}, rows: [][]driver.Value{{int64(1), int64(1), int64(2), "ACTIVA"}}},
			controllerDBResult{},
		)
		assertStatus(t, requestWithParam(http.MethodPut, "/entradas/1", "/entradas/:id", "", controller.CancelarEntrada, float64(1)), http.StatusOK)
	})
	t.Run("transferir id invalido", func(t *testing.T) {
		assertStatus(t, requestWithParam(http.MethodPut, "/entradas/x", "/entradas/:id", `{"email":"bruno@example.com"}`, controller.TransferirEntrada, float64(1)), http.StatusBadRequest)
	})
	t.Run("transferir json invalido", func(t *testing.T) {
		assertStatus(t, requestWithParam(http.MethodPut, "/entradas/1", "/entradas/:id", "{", controller.TransferirEntrada, float64(1)), http.StatusBadRequest)
	})
	t.Run("transferir email inexistente", func(t *testing.T) {
		useControllerDB(t,
			controllerDBResult{columns: []string{"id", "usuario_id", "evento_id", "estado"}, rows: [][]driver.Value{{int64(1), int64(1), int64(2), "ACTIVA"}}},
			controllerDBResult{columns: []string{"id", "nombre", "email", "password_hash", "rol"}},
		)
		assertStatus(t, requestWithParam(http.MethodPut, "/entradas/1", "/entradas/:id", `{"email":"nadie@example.com"}`, controller.TransferirEntrada, float64(1)), http.StatusBadRequest)
	})
	t.Run("transferir exitoso", func(t *testing.T) {
		useControllerDB(t,
			controllerDBResult{columns: []string{"id", "usuario_id", "evento_id", "estado"}, rows: [][]driver.Value{{int64(1), int64(1), int64(2), "ACTIVA"}}},
			controllerDBResult{columns: []string{"id", "nombre", "email", "password_hash", "rol"}, rows: [][]driver.Value{{int64(3), "Bruno", "bruno@example.com", "hash", "CLIENTE"}}},
			controllerDBResult{},
		)
		assertStatus(t, requestWithParam(http.MethodPut, "/entradas/1", "/entradas/:id", `{"email":"bruno@example.com"}`, controller.TransferirEntrada, float64(1)), http.StatusOK)
	})
}

func TestPuntuacionController(t *testing.T) {
	controller := NewPuntuacionController()

	t.Run("crear json invalido", func(t *testing.T) {
		assertStatus(t, performRequest(http.MethodPost, "/puntuaciones", "{", controller.CrearPuntuacion, float64(1)), http.StatusBadRequest)
	})
	t.Run("sin usuario", func(t *testing.T) {
		assertStatus(t, performRequest(http.MethodPost, "/puntuaciones", `{"evento_id":1,"puntuacion":5}`, controller.CrearPuntuacion, nil), http.StatusUnauthorized)
	})
	t.Run("tipo usuario invalido", func(t *testing.T) {
		assertStatus(t, performRequest(http.MethodPost, "/puntuaciones", `{"evento_id":1,"puntuacion":5}`, controller.CrearPuntuacion, "1"), http.StatusUnauthorized)
	})
	t.Run("validacion service", func(t *testing.T) {
		assertStatus(t, performRequest(http.MethodPost, "/puntuaciones", `{"evento_id":0,"puntuacion":5}`, controller.CrearPuntuacion, float64(1)), http.StatusBadRequest)
	})
	t.Run("crear exitoso", func(t *testing.T) {
		useControllerDB(t, controllerDBResult{columns: eventoColumns(), rows: [][]driver.Value{eventoRow()}}, controllerDBResult{})
		assertStatus(t, performRequest(http.MethodPost, "/puntuaciones", `{"evento_id":1,"puntuacion":5}`, controller.CrearPuntuacion, float64(1)), http.StatusCreated)
	})
	t.Run("crear evento cancelado", func(t *testing.T) {
		row := eventoRow()
		row[11] = "CANCELADO"
		useControllerDB(t, controllerDBResult{columns: eventoColumns(), rows: [][]driver.Value{row}})
		assertStatus(t, performRequest(http.MethodPost, "/puntuaciones", `{"evento_id":1,"puntuacion":5}`, controller.CrearPuntuacion, float64(1)), http.StatusBadRequest)
	})
	t.Run("listar id invalido", func(t *testing.T) {
		assertStatus(t, requestWithParam(http.MethodGet, "/eventos/x/puntuaciones", "/eventos/:id/puntuaciones", "", controller.ObtenerPuntuacionesPorEvento, nil), http.StatusBadRequest)
	})
	t.Run("listar exitoso", func(t *testing.T) {
		useControllerDB(t, controllerDBResult{columns: []string{"id", "usuario_id", "evento_id", "puntuacion"}, rows: [][]driver.Value{{int64(1), int64(2), int64(3), int64(5)}}})
		assertStatus(t, requestWithParam(http.MethodGet, "/eventos/3/puntuaciones", "/eventos/:id/puntuaciones", "", controller.ObtenerPuntuacionesPorEvento, nil), http.StatusOK)
	})
	t.Run("listar error", func(t *testing.T) {
		useControllerDB(t, controllerDBResult{err: errors.New("fallo puntuaciones")})
		assertStatus(t, requestWithParam(http.MethodGet, "/eventos/3/puntuaciones", "/eventos/:id/puntuaciones", "", controller.ObtenerPuntuacionesPorEvento, nil), http.StatusInternalServerError)
	})
	t.Run("promedio id invalido", func(t *testing.T) {
		assertStatus(t, requestWithParam(http.MethodGet, "/eventos/x/promedio", "/eventos/:id/promedio", "", controller.ObtenerPromedioPorEvento, nil), http.StatusBadRequest)
	})
	t.Run("promedio error", func(t *testing.T) {
		useControllerDB(t, controllerDBResult{err: errors.New("fallo promedio")})
		assertStatus(t, requestWithParam(http.MethodGet, "/eventos/3/promedio", "/eventos/:id/promedio", "", controller.ObtenerPromedioPorEvento, nil), http.StatusInternalServerError)
	})
	t.Run("promedio exitoso", func(t *testing.T) {
		useControllerDB(t, controllerDBResult{columns: []string{"promedio"}, rows: [][]driver.Value{{4.5}}})
		assertStatus(t, requestWithParam(http.MethodGet, "/eventos/3/promedio", "/eventos/:id/promedio", "", controller.ObtenerPromedioPorEvento, nil), http.StatusOK)
	})
	t.Run("ranking error", func(t *testing.T) {
		useControllerDB(t, controllerDBResult{err: errors.New("fallo ranking")})
		assertStatus(t, performRequest(http.MethodGet, "/ranking", "", controller.ObtenerRankingEventos, nil), http.StatusInternalServerError)
	})
	t.Run("ranking exitoso", func(t *testing.T) {
		useControllerDB(t, controllerDBResult{columns: []string{"evento_id", "titulo", "promedio"}, rows: [][]driver.Value{{int64(3), "Recital", 4.5}}})
		assertStatus(t, performRequest(http.MethodGet, "/ranking", "", controller.ObtenerRankingEventos, nil), http.StatusOK)
	})
}

func TestUsuarioController(t *testing.T) {
	controller := NewUsuarioController()

	t.Run("obtener usuarios exitoso", func(t *testing.T) {
		useControllerDB(t, controllerDBResult{
			columns: []string{"id", "nombre", "email", "password_hash", "rol"},
			rows:    [][]driver.Value{{int64(1), "Ana", "ana@example.com", "hash", "CLIENTE"}},
		})
		assertStatus(t, performRequest(http.MethodGet, "/usuarios", "", controller.ObtenerUsuarios, nil), http.StatusOK)
	})

	t.Run("obtener usuarios error", func(t *testing.T) {
		useControllerDB(t, controllerDBResult{err: errors.New("fallo usuarios")})
		assertStatus(t, performRequest(http.MethodGet, "/usuarios", "", controller.ObtenerUsuarios, nil), http.StatusInternalServerError)
	})

	t.Run("cambiar rol id invalido", func(t *testing.T) {
		assertStatus(t, requestWithParam(http.MethodPut, "/usuarios/x", "/usuarios/:id", `{"rol":"ADMIN"}`, controller.CambiarRol, nil), http.StatusBadRequest)
	})

	t.Run("cambiar rol json invalido", func(t *testing.T) {
		assertStatus(t, requestWithParam(http.MethodPut, "/usuarios/1", "/usuarios/:id", "{", controller.CambiarRol, nil), http.StatusBadRequest)
	})

	t.Run("cambiar rol invalido", func(t *testing.T) {
		assertStatus(t, requestWithParam(http.MethodPut, "/usuarios/1", "/usuarios/:id", `{"rol":"ROOT"}`, controller.CambiarRol, nil), http.StatusBadRequest)
	})

	t.Run("cambiar rol exitoso", func(t *testing.T) {
		useControllerDB(t, controllerDBResult{})
		assertStatus(t, requestWithParam(http.MethodPut, "/usuarios/1", "/usuarios/:id", `{"rol":"ADMIN"}`, controller.CambiarRol, nil), http.StatusOK)
	})
}

func TestReporteEventoController(t *testing.T) {
	controller := NewEntradaController()

	t.Run("id invalido", func(t *testing.T) {
		assertStatus(t, requestWithParam(http.MethodGet, "/eventos/x/reporte", "/eventos/:id/reporte", "", controller.ObtenerReporteEvento, nil), http.StatusBadRequest)
	})

	t.Run("evento inexistente", func(t *testing.T) {
		useControllerDB(t, controllerDBResult{columns: eventoColumns()})
		assertStatus(t, requestWithParam(http.MethodGet, "/eventos/99/reporte", "/eventos/:id/reporte", "", controller.ObtenerReporteEvento, nil), http.StatusBadRequest)
	})

	t.Run("exitoso con compradores", func(t *testing.T) {
		useControllerDB(t,
			controllerDBResult{columns: eventoColumns(), rows: [][]driver.Value{eventoRow()}},
			controllerDBResult{columns: []string{"cantidad"}, rows: [][]driver.Value{{int64(1)}}},
			controllerDBResult{columns: []string{"id", "nombre", "email"}, rows: [][]driver.Value{{int64(2), "Ana", "ana@example.com"}}},
		)
		response := requestWithParam(http.MethodGet, "/eventos/1/reporte", "/eventos/:id/reporte", "", controller.ObtenerReporteEvento, nil)
		assertStatus(t, response, http.StatusOK)
		if !strings.Contains(response.Body.String(), `"porcentaje_ocupado":1`) {
			t.Fatalf("body de reporte inesperado: %s", response.Body.String())
		}
	})
}
