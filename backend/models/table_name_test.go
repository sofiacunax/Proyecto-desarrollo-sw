package models

import "testing"

func TestTableNames(t *testing.T) {
	tests := []struct {
		name string
		got  string
		want string
	}{
		{name: "usuario", got: (Usuario{}).TableName(), want: "usuarios"},
		{name: "evento", got: (Evento{}).TableName(), want: "eventos"},
		{name: "entrada", got: (Entrada{}).TableName(), want: "entradas"},
		{name: "puntuacion", got: (Puntuacion{}).TableName(), want: "puntuaciones"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.got != test.want {
				t.Fatalf("tabla esperada %q, obtenida %q", test.want, test.got)
			}
		})
	}
}
