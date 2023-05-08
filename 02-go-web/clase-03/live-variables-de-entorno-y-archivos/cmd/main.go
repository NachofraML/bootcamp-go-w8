package main

import (
	"github.com/NachofraML/bootcamp-go-w8/web/clase-03/live-variables-de-entorno-y-archivos/cmd/handler"
	"github.com/NachofraML/bootcamp-go-w8/web/clase-03/live-variables-de-entorno-y-archivos/internal/product"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./web/clase-03/live-variables-de-entorno-y-archivos/.env")
	if err != nil {
		panic(err)
	}
	//rp := product.NewInMemoryRepository()
	rp := product.NewJsonStorageRepository()
	if err != nil {
		panic(err)
	}
	s := product.NewDefaultService(rp)
	cr := handler.NewProductController(s)

	sv := gin.Default()

	p := sv.Group("/products")
	{
		p.GET("/:id", cr.GetId())
		p.POST("", cr.Save())
		p.PUT("/:id", cr.Update())
		p.PATCH("/:id", cr.UpdatePartial())
		p.DELETE("/:id", cr.Delete())
	}
	err = sv.Run(":8080")
	if err != nil {
		panic(err)
	}
}
