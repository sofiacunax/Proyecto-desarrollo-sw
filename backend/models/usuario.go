package models

type Usuario struct {
	ID           int          `json:"id" gorm:"primaryKey;autoIncrement;type:int"`
	Nombre       string       `json:"nombre" gorm:"size:100;not null"`
	Email        string       `json:"email" gorm:"size:100;not null;uniqueIndex"`
	PasswordHash string       `json:"-" gorm:"column:password_hash;size:255;not null"`
	Rol          string       `json:"rol" gorm:"size:20;not null"`
	Entradas     []Entrada    `json:"-" gorm:"foreignKey:UsuarioID;references:ID"`
	Puntuaciones []Puntuacion `json:"-" gorm:"foreignKey:UsuarioID;references:ID"`
}

func (Usuario) TableName() string { return "usuarios" }
