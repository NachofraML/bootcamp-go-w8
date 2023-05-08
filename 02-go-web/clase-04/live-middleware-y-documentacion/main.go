package main

import (
	"github.com/NachofraML/bootcamp-go-w8/web/clase-04/live-middleware-y-documentacion/cmd/handler"
	"github.com/NachofraML/bootcamp-go-w8/web/clase-04/live-middleware-y-documentacion/cmd/middleware"
	"github.com/NachofraML/bootcamp-go-w8/web/clase-04/live-middleware-y-documentacion/docs"
	"github.com/NachofraML/bootcamp-go-w8/web/clase-04/live-middleware-y-documentacion/internal/product"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"os"
)

func main() {
	err := godotenv.Load("./web/clase-04/live-middleware-y-documentacion/.env")
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

	sv := gin.New()
	sv.Use(gin.Recovery())

	//Docs route
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	sv.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	p := sv.Group("/products", middleware.TokenMiddleware, middleware.LogResponseMiddleware)
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
