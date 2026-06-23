package db

import (
	"fmt"
	"log"
	"os"
	"proyecto-desarrollo-sw/backend/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		getEnv("DB_USER", "root"),
		getEnv("DB_PASSWORD", "300106"),
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "3306"),
		getEnv("DB_NAME", "proyecto_desarrollo_sw"),
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

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

	if err = DB.Exec("ALTER TABLE eventos MODIFY horario TIME NOT NULL").Error; err != nil {
		log.Fatal(err)
	}

	log.Println("Conectado a MySQL y esquema migrado")
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}
