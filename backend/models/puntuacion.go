package models

type Puntuacion struct {
	ID         int `json:"id"`
	UsuarioID  int `json:"usuario_id"`
	EventoID   int `json:"evento_id"`
	Puntuacion int `json:"puntuacion"`
}
