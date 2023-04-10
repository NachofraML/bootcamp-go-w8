package main

import (
	"github.com/NachofraML/bootcamp-go-w8/web/clase-02/live-estructuramos-nuestra-api/cmd/handler"
	"github.com/NachofraML/bootcamp-go-w8/web/clase-02/live-estructuramos-nuestra-api/internal/product"
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
	}
	err := sv.Run(":8080")
	if err != nil {
		panic(err)
	}
}
