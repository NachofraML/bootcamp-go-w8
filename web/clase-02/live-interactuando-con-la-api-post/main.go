package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func main() {
	sv := gin.Default()

	p := sv.Group("/products")

	p.GET("/:id", GetProductById())
	p.POST("", SaveProduct())

	err := sv.Run(":8080")
	if err != nil {
		panic(err)
	}
}

type Producto struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func SaveProduct() gin.HandlerFunc {
	type request struct {
		Name        string  `json:"name" binding:"required"`
		Quantity    int     `json:"quantity" binding:"required"`
		CodeValue   string  `json:"code_value" binding:"required"`
		IsPublished bool    `json:"is_published"`
		Expiration  string  `json:"expiration" binding:"required"`
		Price       float64 `json:"price" binding:"required"`
	}
	return func(c *gin.Context) {
		var req request
		err := c.ShouldBindJSON(&req)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid request", "data": nil})
			return
		}
		err = validExpirationDate(req.Expiration)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid expiration date", "data": nil})
			return
		}
		err = validCodeValue(req.CodeValue)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid code value, unique violation", "data": nil})
			return
		}
		producto.ID = lastID + 1
		producto.Name = req.Name
		producto.Quantity = req.Quantity
		producto.CodeValue = req.CodeValue
		producto.IsPublished = req.IsPublished
		producto.Expiration = req.Expiration
		producto.Price = req.Price
		lastID++
		*productos = append(*productos, *producto)
		c.JSON(http.StatusCreated, gin.H{"message": "success", "data": producto})
	}
}

func GetProductById() gin.HandlerFunc {
	type response struct {
		ID          int     `json:"id"`
		Name        string  `json:"name"`
		Quantity    int     `json:"quantity"`
		CodeValue   string  `json:"code_value"`
		IsPublished bool    `json:"is_published"`
		Expiration  string  `json:"expiration"`
		Price       float64 `json:"price"`
	}
	return func(c *gin.Context) {
		stringRequestId := c.Param("id")
		if stringRequestId == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "an id is required to search products", "data": nil})
			return
		}
		intRequestId, err := strconv.Atoi(stringRequestId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "cannot parse id to int", "data": nil})
			return
		}
		for _, p := range *productos {
			if p.ID == intRequestId {
				var res response
				res.ID = p.ID
				res.Name = p.Name
				res.Quantity = p.Quantity
				res.CodeValue = p.CodeValue
				res.IsPublished = p.IsPublished
				res.Expiration = p.Expiration
				res.Price = p.Price
				c.JSON(http.StatusOK, gin.H{"message": "success", "data": res})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "product not found", "data": nil})
	}

}

func validExpirationDate(expirationDate string) error {
	const YYYYMMDD string = "02/01/2006"
	_, err := time.Parse(YYYYMMDD, expirationDate)
	if err != nil {
		return errors.New("invalid expiration date format")
	}
	return nil
}

func validCodeValue(actualCode string) error {
	for _, p := range *productos {
		if p.CodeValue == actualCode {
			return errors.New("cannot create code_value, unique violation")
		}
	}
	return nil
}

var producto *Producto = &Producto{}
var productos *[]Producto = &[]Producto{}
var lastID int = 0
