package db

import (
	"log"
	"proyecto-desarrollo-sw/backend/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error

	DB, err = gorm.Open(mysql.Open(
		"root:SasukeKira2006.@tcp(localhost:3306)/proyecto_desarrollo_sw?charset=utf8mb4",
	), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	if err = DB.AutoMigrate(
		&models.Usuario{},
		&models.Evento{},
		&models.Entrada{},
		&models.Puntuacion{},
	); err != nil {
		log.Fatal(err)
	}

	log.Println("Conectado a MySQL y esquema migrado")
}
