package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestObtenerEventoPorIDInvalido(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	controller := NewEventoController()

	router.GET("/eventos/:id", controller.ObtenerEventoPorID)

	req, _ := http.NewRequest("GET", "/eventos/abc", nil)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("se esperaba %d y se obtuvo %d",
			http.StatusBadRequest,
			w.Code,
		)
	}
}

func TestEliminarEventoIDInvalido(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	controller := NewEventoController()

	router.DELETE("/eventos/:id", controller.EliminarEvento)

	req, _ := http.NewRequest("DELETE", "/eventos/abc", nil)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("se esperaba %d y se obtuvo %d",
			http.StatusBadRequest,
			w.Code,
		)
	}
}

func TestCrearEventoJSONInvalido(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	controller := NewEventoController()

	router.POST("/eventos", controller.CrearEvento)

	req, _ := http.NewRequest(
		"POST",
		"/eventos",
		strings.NewReader("json-invalido"),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("se esperaba %d y se obtuvo %d",
			http.StatusBadRequest,
			w.Code,
		)
	}
}
