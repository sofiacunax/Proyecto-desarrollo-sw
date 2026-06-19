package models

type Evento struct {
	ID           int          `gorm:"primaryKey;autoIncrement;type:int"`
	Titulo       string       `gorm:"size:255;not null"`
	Descripcion  string       `gorm:"type:text"`
	Fecha        string       `gorm:"type:date;not null"`
	Horario      string       `gorm:"type:time;not null"`
	Duracion     int          `gorm:"not null"`
	Ubicacion    string       `gorm:"size:255;not null"`
	Capacidad    int          `gorm:"not null"`
	Precio       float64      `gorm:"type:decimal(10,2);not null"`
	Categoria    string       `gorm:"size:100"`
	ImagenURL    string       `gorm:"column:imagen_url;size:255"`
	Estado       string       `gorm:"size:20;not null"`
	Entradas     []Entrada    `json:"-" gorm:"foreignKey:EventoID;references:ID"`
	Puntuaciones []Puntuacion `json:"-" gorm:"foreignKey:EventoID;references:ID"`
}

func (Evento) TableName() string { return "eventos" }
