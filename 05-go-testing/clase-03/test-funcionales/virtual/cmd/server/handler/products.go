package handler

import (
	"meli-bootcamp/internal/products"

	"github.com/gin-gonic/gin"
)
//creo la estructura de la request
type request struct {
	Name  string  `json:"nombre"`
	Type  string  `json:"tipo"`
	Count int     `json:"cantidad"`
	Price float64 `json:"precio"`
}
//estructura que voy a usar en el servicio
type Product struct {
	service products.Service
}

func NewProduct(p products.Service) *Product {
	return &Product{
		service: p,
	}
}

func (c *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		p, err := c.service.Store(req.Name, req.Type, req.Count, req.Price)
		if err != nil {
			ctx.JSON(401, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, p)
	}
}
