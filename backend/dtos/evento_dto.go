package dtos

type EventoDTO struct {
	Titulo      string  `json:"titulo"`
	Descripcion string  `json:"descripcion"`
	Fecha       string  `json:"fecha"`
	Horario     string  `json:"horario"`
	Duracion    int     `json:"duracion"`
	Ubicacion   string  `json:"ubicacion"`
	Capacidad   int     `json:"capacidad"`
	Precio      float64 `json:"precio"`
	Categoria   string  `json:"categoria"`
	ImagenURL   string  `json:"imagen_url"`
	Estado      string  `json:"estado"`
}
