package dtos

type RankingDTO struct {
	EventoID int     `json:"evento_id"`
	Titulo   string  `json:"titulo"`
	Promedio float64 `json:"promedio"`
}
