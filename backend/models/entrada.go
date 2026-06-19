package models

type Entrada struct {
	ID        int     `gorm:"primaryKey;autoIncrement;type:int"`
	UsuarioID int     `gorm:"column:usuario_id;type:int;not null;index"`
	EventoID  int     `gorm:"column:evento_id;type:int;not null;index"`
	Estado    string  `gorm:"size:20;not null"`
	Usuario   Usuario `json:"-" gorm:"foreignKey:UsuarioID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Evento    Evento  `json:"-" gorm:"foreignKey:EventoID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (Entrada) TableName() string { return "entradas" }
