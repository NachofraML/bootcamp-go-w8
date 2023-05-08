package mocks

import (
	"database/sql"
	"github.com/DATA-DOG/go-txdb"
	"github.com/go-sql-driver/mysql"
	"log"
)

func InitDb() (*sql.DB, error) {
	dbMock, err := sql.Open("txdb", "my_db")
	if err != nil {
		log.Fatal("Error while creating mock")
	}
	return dbMock, nil
}

func init() {
	databaseConfig := mysql.Config{
		User:   "root",
		Passwd: "",
		Addr:   "127.0.0.1:3306",
		DBName: "my_db",
	}

	txdb.Register("txdb", "mysql", databaseConfig.FormatDSN())
}
