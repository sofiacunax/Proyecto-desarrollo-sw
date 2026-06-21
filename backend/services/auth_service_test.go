package services

import (
	"database/sql/driver"
	"errors"
	"strings"
	"testing"

	"golang.org/x/crypto/bcrypt"

	"proyecto-desarrollo-sw/backend/dtos"
)

func TestAuthServiceRegister(t *testing.T) {
	useDatabaseStub(t, databaseResult{})
	service := &AuthService{}
	err := service.Register(dtos.RegisterRequest{
		Nombre: "Ana", Email: "ana@example.com", Password: "secreto",
	})
	if err != nil {
		t.Fatalf("Register devolvio un error inesperado: %v", err)
	}
}

func TestAuthServiceRegisterPropagaErrorDePersistencia(t *testing.T) {
	useDatabaseStub(t, databaseResult{err: errors.New("fallo insert")})
	service := &AuthService{}
	if err := service.Register(dtos.RegisterRequest{Password: "secreto"}); err == nil {
		t.Fatal("se esperaba el error de persistencia")
	}
}

func TestAuthServiceLogin(t *testing.T) {
	hash, err := bcrypt.GenerateFromPassword([]byte("secreto"), bcrypt.MinCost)
	if err != nil {
		t.Fatal(err)
	}
	useDatabaseStub(t, databaseResult{
		columns: []string{"id", "nombre", "email", "password_hash", "rol"},
		rows:    [][]driver.Value{{int64(7), "Ana", "ana@example.com", string(hash), "CLIENTE"}},
	})
	service := &AuthService{}
	token, err := service.Login(dtos.LoginRequest{Email: "ana@example.com", Password: "secreto"})
	if err != nil || token == "" {
		t.Fatalf("se esperaba un token valido, token=%q err=%v", token, err)
	}
}

func TestAuthServiceLoginRechazaCredenciales(t *testing.T) {
	t.Run("usuario inexistente", func(t *testing.T) {
		useDatabaseStub(t, databaseResult{err: errors.New("sin filas")})
		_, err := (&AuthService{}).Login(dtos.LoginRequest{Email: "nadie@example.com"})
		if err == nil || !strings.Contains(err.Error(), "credenciales") {
			t.Fatalf("se esperaba error de credenciales, se obtuvo %v", err)
		}
	})

	t.Run("password incorrecto", func(t *testing.T) {
		hash, _ := bcrypt.GenerateFromPassword([]byte("correcto"), bcrypt.MinCost)
		useDatabaseStub(t, databaseResult{
			columns: []string{"id", "nombre", "email", "password_hash", "rol"},
			rows:    [][]driver.Value{{int64(1), "Ana", "ana@example.com", string(hash), "CLIENTE"}},
		})
		_, err := (&AuthService{}).Login(dtos.LoginRequest{Email: "ana@example.com", Password: "incorrecto"})
		if err == nil {
			t.Fatal("se esperaba error para password incorrecto")
		}
	})
}
