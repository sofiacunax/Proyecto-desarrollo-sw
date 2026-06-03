package models

type Evento struct {
	ID          int
	Titulo      string
	Descripcion string
	Fecha       string
	Horario     string
	Duracion    int
	Ubicacion   string
	Capacidad   int
	Precio      float64
	Categoria   string
	ImagenURL   string
	Estado      string
}
