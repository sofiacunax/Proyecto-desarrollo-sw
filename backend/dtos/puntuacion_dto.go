package dtos

type PuntuacionDTO struct {
	UsuarioID  int `json:"usuario_id"`
	EventoID   int `json:"evento_id"`
	Puntuacion int `json:"puntuacion"`
}
