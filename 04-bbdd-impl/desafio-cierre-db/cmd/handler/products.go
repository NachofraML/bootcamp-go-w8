package handler

import (
	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
	"github.com/bootcamp-go/desafio-cierre-db.git/internal/products"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Products struct {
	s products.Service
}

func NewHandlerProducts(s products.Service) *Products {
	return &Products{s}
}

func (p *Products) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := p.s.ReadAll()
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, products)
	}
}

func (p *Products) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products := domain.Products{}
		err := ctx.ShouldBindJSON(&products)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		err = p.s.Create(&products)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(201, gin.H{"data": products})
	}
}

func (p *Products) LoadFromJson() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := p.s.LoadFromJson()
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, gin.H{"data": "successfully migrated products json to database"})
	}
}

func (p *Products) GetTopProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		limit, _ := ctx.GetQuery("limit")
		parsedLimit, err := strconv.Atoi(limit)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "limit must be a number"})
			return
		}

		topProducts, err := p.s.GetTopProducts(parsedLimit)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, gin.H{"data": topProducts})
	}
}
