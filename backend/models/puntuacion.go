package models

type Puntuacion struct {
	ID         int     `json:"id" gorm:"primaryKey;autoIncrement;type:int"`
	UsuarioID  int     `json:"usuario_id" gorm:"column:usuario_id;type:int;not null;uniqueIndex:idx_usuario_evento"`
	EventoID   int     `json:"evento_id" gorm:"column:evento_id;type:int;not null;uniqueIndex:idx_usuario_evento"`
	Puntuacion int     `json:"puntuacion" gorm:"not null"`
	Usuario    Usuario `json:"-" gorm:"foreignKey:UsuarioID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Evento     Evento  `json:"-" gorm:"foreignKey:EventoID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (Puntuacion) TableName() string { return "puntuaciones" }
