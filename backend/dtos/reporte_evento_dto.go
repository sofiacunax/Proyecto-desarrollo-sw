package dtos

type CompradorDTO struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
}

type ReporteEventoDTO struct {
	EventoID          int            `json:"evento_id"`
	Titulo            string         `json:"titulo"`
	Capacidad         int            `json:"capacidad"`
	EntradasVendidas  int            `json:"entradas_vendidas"`
	PorcentajeOcupado float64        `json:"porcentaje_ocupado"`
	Compradores       []CompradorDTO `json:"compradores"`
}
