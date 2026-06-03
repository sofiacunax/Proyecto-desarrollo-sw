package models

type Usuario struct {
	ID           uint   `json:"id"`
	Nombre       string `json:"nombre"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
	Rol          string `json:"rol"`
}
