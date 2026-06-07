package dtos

type MisEntradaDTO struct {
	ID        int    `json:"id"`
	EventoID  int    `json:"evento_id"`
	Titulo    string `json:"titulo"`
	Fecha     string `json:"fecha"`
	Ubicacion string `json:"ubicacion"`
	Estado    string `json:"estado"`
}