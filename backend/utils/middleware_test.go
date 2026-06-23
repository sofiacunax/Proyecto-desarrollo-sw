package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestAuthMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("token inexistente", func(t *testing.T) {
		response := runProtectedRequest("")
		if response.Code != http.StatusUnauthorized {
			t.Fatalf("status esperado %d, obtenido %d: %s", http.StatusUnauthorized, response.Code, response.Body.String())
		}
	})

	t.Run("token invalido", func(t *testing.T) {
		response := runProtectedRequest("Bearer token-roto")
		if response.Code != http.StatusUnauthorized {
			t.Fatalf("status esperado %d, obtenido %d: %s", http.StatusUnauthorized, response.Code, response.Body.String())
		}
	})

	t.Run("token valido", func(t *testing.T) {
		token, err := GenerateToken(7, "admin@example.com", "ADMIN")
		if err != nil {
			t.Fatalf("no se pudo generar token: %v", err)
		}

		response := runProtectedRequest("Bearer " + token)
		if response.Code != http.StatusOK {
			t.Fatalf("status esperado %d, obtenido %d: %s", http.StatusOK, response.Code, response.Body.String())
		}
		if response.Body.String() != `{"email":"admin@example.com","rol":"ADMIN","userID":7}` {
			t.Fatalf("body inesperado: %s", response.Body.String())
		}
	})
}

func TestAdminMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("cliente denegado", func(t *testing.T) {
		response := runAdminRequest("CLIENTE")
		if response.Code != http.StatusForbidden {
			t.Fatalf("status esperado %d, obtenido %d: %s", http.StatusForbidden, response.Code, response.Body.String())
		}
	})

	t.Run("admin permitido", func(t *testing.T) {
		response := runAdminRequest("ADMIN")
		if response.Code != http.StatusOK {
			t.Fatalf("status esperado %d, obtenido %d: %s", http.StatusOK, response.Code, response.Body.String())
		}
	})
}

func runProtectedRequest(authorization string) *httptest.ResponseRecorder {
	router := gin.New()
	router.Use(AuthMiddleware())
	router.GET("/private", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"userID": ctx.GetFloat64("userID"),
			"email":  ctx.GetString("email"),
			"rol":    ctx.GetString("rol"),
		})
	})

	request := httptest.NewRequest(http.MethodGet, "/private", nil)
	if authorization != "" {
		request.Header.Set("Authorization", authorization)
	}

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	return response
}

func runAdminRequest(rol string) *httptest.ResponseRecorder {
	router := gin.New()
	router.Use(func(ctx *gin.Context) {
		ctx.Set("rol", rol)
		ctx.Next()
	})
	router.Use(AdminMiddleware())
	router.GET("/admin", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	response := httptest.NewRecorder()
	router.ServeHTTP(response, httptest.NewRequest(http.MethodGet, "/admin", nil))
	return response
}
