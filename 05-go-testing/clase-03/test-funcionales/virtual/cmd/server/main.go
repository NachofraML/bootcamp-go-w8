package main

import (
	"meli-bootcamp/cmd/server/handler"
	"meli-bootcamp/internal/products"
	"meli-bootcamp/pkg/store"

	"github.com/gin-gonic/gin"
)

func main() {
	db := store.New(store.FileType, "./products.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()
	pr := r.Group("/products")
	pr.POST("/", p.Store())
	
	r.Run()
}
