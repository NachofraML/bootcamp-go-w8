package handler

import (
	"errors"
	"github.com/NachofraML/bootcamp-go-w8/web/clase-02/live-estructuramos-nuestra-api/internal/domain"
	"github.com/NachofraML/bootcamp-go-w8/web/clase-02/live-estructuramos-nuestra-api/internal/product"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func NewProductController(sv product.Service) *ProductController {
	return &ProductController{sv: sv}
}

type ProductController struct {
	sv product.Service
}

func (pr *ProductController) Save() gin.HandlerFunc {
	type request struct {
		Name        string  `json:"name"`
		Quantity    int     `json:"quantity"`
		CodeValue   string  `json:"code_value"`
		IsPublished bool    `json:"is_published"`
		Expiration  string  `json:"expiration"`
		Price       float64 `json:"price"`
	}
	return func(ctx *gin.Context) {
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
			return
		}
		producto := &domain.Producto{
			Name:        req.Name,
			Quantity:    req.Quantity,
			CodeValue:   req.CodeValue,
			IsPublished: req.IsPublished,
			Expiration:  req.Expiration,
			Price:       req.Price,
		}

		err := pr.sv.Save(producto)
		if err != nil {
			if errors.Is(err, product.ErrServiceInvalid) {
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid product"})
				return
			}
			if errors.Is(err, product.ErrServiceNotUnique) {
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "code value already exists, try with another"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"message": "success", "data": producto})
	}
}

func (pr *ProductController) GetId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
			return
		}
		productoEncontrado, err := pr.sv.GetByID(id)
		if err != nil {
			if errors.Is(err, product.ErrServiceNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{"message": "movie not found"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": productoEncontrado})
	}

}
