package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

//Debemos crear un repositorio en github.com para poder subir nuestros avances. Este repositorio es el que vamos a
//utilizar para llevar lo que realicemos durante las distintas prácticas de Go Web.
//Primero debemos clonar el repositorio creado, luego iniciar nuestro proyecto de go con con el comando go mod init.
//El siguiente paso será crear un archivo main.go donde deberán cargar en una slice, desde un archivo JSON, los datos
//de productos. Esta slice se debe cargar cada vez que se inicie la API para realizar las distintas consultas.
//El archivo para trabajar es el siguiente:

type Product struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

var Products, _ = ParseJSON()

func ParseJSON() (products []Product, err error) {

	file, err := os.ReadFile("products.json")
	if err != nil {
		return
	}
	err = json.Unmarshal(file, &products)
	if err != nil {
		return
	}
	return
}

func getAllProducts(c *gin.Context) {
	c.JSON(http.StatusOK, Products)
}

func getProductByID(c *gin.Context) {
	ID, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	for _, p := range Products {
		if ID == p.ID {
			c.JSON(http.StatusOK, p)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "Product not found",
	})
}

func GetProductByPrice(c *gin.Context) {
	var findedProducts []Product
	price, err := strconv.ParseFloat(c.Query("priceGt"), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal error",
		})
	}

	for _, p := range Products {
		if p.Price > price {
			findedProducts = append(findedProducts, p)
		}
	}

	if len(findedProducts) <= 0 {
		c.JSON(http.StatusNoContent, gin.H{
			"error": "Products more expensive than price entered dont exist",
		})
	} else {
		c.JSON(http.StatusOK, findedProducts)
	}
}

func main() {
	server := gin.Default()

	serverProducts := server.Group("/products")
	serverProducts.GET("", getAllProducts)
	serverProducts.GET("/:id", getProductByID)
	serverProducts.GET("/search", GetProductByPrice)

	err := server.Run(":8080")
	if err != nil {
		panic(err)
	}
}
