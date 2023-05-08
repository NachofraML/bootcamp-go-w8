package router

import (
	"github.com/bootcamp-go/desafio-go-web/cmd/handler"
	"github.com/bootcamp-go/desafio-go-web/internal/domain"
	"github.com/bootcamp-go/desafio-go-web/internal/tickets"
	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine, db []domain.Ticket) *RouteHandler {
	rp := tickets.NewRepository(db)
	sv := tickets.NewService(rp)
	cr := handler.NewTicketsController(sv)

	return &RouteHandler{cr: *cr, e: r}
}

type RouteHandler struct {
	cr handler.TicketsController
	e  *gin.Engine
}

func (r *RouteHandler) MapRoutes() {
	t := r.e.Group("/ticket")

	t.GET("/getByCountry/:dest", r.cr.GetByCountry())
	t.GET("/getAverage/:dest", r.cr.GetAverage())
}
