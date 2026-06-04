package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	var err error

	DB, err = sql.Open(
		"mysql",
		"root:Popia115a@tcp(localhost:3306)/proyecto_desarrollo_sw",
	)

	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Conectado a MySQL")
}
