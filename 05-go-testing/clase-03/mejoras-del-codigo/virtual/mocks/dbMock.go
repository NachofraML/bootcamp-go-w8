package mocks

import (
	"database/sql"
	"log"

	"github.com/DATA-DOG/go-txdb"
)

func InitDb() (*sql.DB, error) {
	dbMock, err := sql.Open("txdb", "identifier")
	if err != nil {
		log.Fatal("Error al crear el mock")
	}

	return dbMock, nil
}

func init() {
	txdb.Register("txdb", "mysql", "root:root@/melisprint")
}
