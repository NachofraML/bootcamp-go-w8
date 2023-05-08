package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"vivo/internal/products"
)

func main() {
	databaseConfig := mysql.Config{
		User:   "root",
		Passwd: "",
		Addr:   "127.0.0.1:3306",
		DBName: "my_db",
	}

	database, err := sql.Open("mysql", databaseConfig.FormatDSN())
	if err != nil {
		panic(err)
	}
	defer database.Close()

	if err = database.Ping(); err != nil {
		panic(err)
	}

	repository := products.NewMySqlRepositoty(database)

	product := products.Product{
		Name:        "Manzana Roja",
		Quantity:    1,
		CodeValue:   "1234",
		IsPublished: true,
		Expiration:  "2023-05-04",
		Price:       100,
	}

	err = repository.Create(&product)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = repository.Get(product.ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	product.Name = "Manzana Verde"
	err = repository.Update(&product)
	if err != nil {
		fmt.Println(err)
		return
	}

	searchedProduct2, err := repository.Get(product.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(searchedProduct2)

	err = repository.Delete(product.ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	searchedProduct3, err := repository.Get(product.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(searchedProduct3)
}
