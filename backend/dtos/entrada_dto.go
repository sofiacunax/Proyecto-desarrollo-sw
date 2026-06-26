package dtos

type ComprarEntradaDTO struct {
	EventoID int `json:"evento_id"`
}
type TransferirEntradaDTO struct {
	Email string `json:"email"`
}
