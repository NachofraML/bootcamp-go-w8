package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/NachofraML/bootcamp-go-w8/web/clase-03/live-variables-de-entorno-y-archivos/internal/domain"
	"github.com/NachofraML/bootcamp-go-w8/web/clase-03/live-variables-de-entorno-y-archivos/internal/product"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
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
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
			return
		}

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
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
			return
		}
		productoEncontrado, err := pr.sv.GetByID(id)
		if err != nil {
			if errors.Is(err, product.ErrServiceNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": productoEncontrado})
	}

}

func (pr *ProductController) Update() gin.HandlerFunc {
	type request struct {
		Name        string  `json:"name" binding:"required"`
		Quantity    int     `json:"quantity" binding:"required"`
		CodeValue   string  `json:"code_value" binding:"required"`
		IsPublished bool    `json:"is_published" binding:"required"`
		Expiration  string  `json:"expiration" binding:"required"`
		Price       float64 `json:"price" binding:"required"`
	}
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
			return
		}
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
			return
		}
		producto := &domain.Producto{
			ID:          id,
			Name:        req.Name,
			Quantity:    req.Quantity,
			CodeValue:   req.CodeValue,
			IsPublished: req.IsPublished,
			Expiration:  req.Expiration,
			Price:       req.Price,
		}

		err = pr.sv.Update(producto.ID, producto)
		if err != nil {
			if errors.Is(err, product.ErrServiceInvalid) {
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid product"})
				return
			}
			if errors.Is(err, product.ErrServiceNotUnique) {
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "entered code value already exists in another product, try yo update with another"})
				return
			}
			if errors.Is(err, product.ErrServiceNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
				return
			}
			fmt.Printf("MAIN: %e", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": producto})
	}
}

func (pr *ProductController) UpdatePartial() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
			return
		}
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
			return
		}

		producto, err := pr.sv.GetByID(id)
		if err != nil {
			if errors.Is(err, product.ErrServiceNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
			return
		}

		if err := json.NewDecoder(ctx.Request.Body).Decode(&producto); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
			return
		}

		err = pr.sv.Update(id, producto)
		if err != nil {
			if errors.Is(err, product.ErrServiceInvalid) {
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid product"})
				return
			}
			if errors.Is(err, product.ErrServiceNotUnique) {
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "entered code value already exists in another product, try yo update with another"})
				return
			}
			if errors.Is(err, product.ErrServiceNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": producto})
	}
}

func (pr *ProductController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
			return
		}
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
			return
		}
		err = pr.sv.Delete(id)
		if err != nil {
			if errors.Is(err, product.ErrServiceNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
			return
		}

		ctx.Header("Location", fmt.Sprintf("/products/%d", id))
		ctx.JSON(http.StatusNoContent, nil)
	}
}
