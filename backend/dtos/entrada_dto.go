package dtos

type ComprarEntradaDTO struct {
	EventoID int `json:"evento_id"`
}

type TransferirEntradaDTO struct {
	NuevoUsuarioID int `json:"nuevo_usuario_id"`
}