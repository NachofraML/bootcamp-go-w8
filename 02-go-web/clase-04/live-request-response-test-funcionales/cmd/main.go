package main

import (
	"github.com/NachofraML/bootcamp-go-w8/web/clase-04/live-request-response-test-funcionales/cmd/handler"
	"github.com/NachofraML/bootcamp-go-w8/web/clase-04/live-request-response-test-funcionales/internal/product"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	err := godotenv.Load("./web/clase-04/live-request-response-test-funcionales/.env")
	if err != nil {
		panic(err)
	}
	jsonStoragePath := os.Getenv("JSON_STORAGE_PATH_DEV")
	rp := product.NewJsonStorageRepository(jsonStoragePath)
	if err != nil {
		panic(err)
	}
	s := product.NewDefaultService(rp)
	cr := handler.NewProductController(s)

	sv := gin.Default()

	p := sv.Group("/products")
	{
		p.GET("", cr.GetAll())
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
