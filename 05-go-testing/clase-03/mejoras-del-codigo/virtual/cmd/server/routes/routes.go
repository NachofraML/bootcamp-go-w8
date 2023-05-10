package routes

import (
	"database/sql"
	"github.com/NachofraML/bootcamp-go-w8/05-go-testing/clase-03/mejoras-del-codigo/virtual/cmd/server/handler"
	"github.com/NachofraML/bootcamp-go-w8/05-go-testing/clase-03/mejoras-del-codigo/virtual/internal/seller"

	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	r  *gin.Engine
	rg *gin.RouterGroup
	db *sql.DB
}

func NewRouter(r *gin.Engine, db *sql.DB) Router {
	return &router{r: r, db: db}
}

func (r *router) MapRoutes() {
	r.setGroup()

	r.buildSellerRoutes()
}

func (r *router) setGroup() {
	r.rg = r.r.Group("/api/v1")
}

func (r *router) buildSellerRoutes() {
	repo := seller.NewRepository(r.db)
	service := seller.NewService(repo)
	sellerHandler := handler.NewSeller(service)

	r.rg.POST("/sellers", sellerHandler.Create())
	r.rg.GET("/sellers", sellerHandler.GetAll())
	r.rg.GET("/sellers/:id", sellerHandler.Get())
}
