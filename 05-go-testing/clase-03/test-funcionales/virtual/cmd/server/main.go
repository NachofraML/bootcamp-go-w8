package main

import (
	"github.com/NachofraML/bootcamp-go-w8/05-go-testing/clase-03/test-funcionales/virtual/cmd/server/handler"
	"github.com/NachofraML/bootcamp-go-w8/05-go-testing/clase-03/test-funcionales/virtual/internal/products"
	"github.com/NachofraML/bootcamp-go-w8/05-go-testing/clase-03/test-funcionales/virtual/pkg/store"
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
