package db

import "testing"

func TestGetEnv(t *testing.T) {
	t.Run("usa default si no existe", func(t *testing.T) {
		t.Setenv("EVENTIA_TEST_ENV", "")
		if got := getEnv("EVENTIA_TEST_ENV", "default"); got != "default" {
			t.Fatalf("valor esperado %q, obtenido %q", "default", got)
		}
	})

	t.Run("usa variable configurada", func(t *testing.T) {
		t.Setenv("EVENTIA_TEST_ENV", "custom")
		if got := getEnv("EVENTIA_TEST_ENV", "default"); got != "custom" {
			t.Fatalf("valor esperado %q, obtenido %q", "custom", got)
		}
	})
}
