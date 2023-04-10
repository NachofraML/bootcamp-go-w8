package main

import (
	"github.com/NachofraML/bootcamp-go-w8/web/clase-03/live-interactuando-con-la-api-otros-metodos/cmd/handler"
	"github.com/NachofraML/bootcamp-go-w8/web/clase-03/live-interactuando-con-la-api-otros-metodos/internal/product"
	"github.com/gin-gonic/gin"
)

func main() {
	rp := product.NewInMemoryRepository()
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
	err := sv.Run(":8080")
	if err != nil {
		panic(err)
	}
}
