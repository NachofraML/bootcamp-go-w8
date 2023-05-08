package handler

import (
	"errors"
	"github.com/bootcamp-go/desafio-go-web/internal/tickets"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewTicketsController(sv tickets.Service) *TicketsController {
	return &TicketsController{sv: sv}
}

type TicketsController struct {
	sv tickets.Service
}

func (tc *TicketsController) GetByCountry() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		destination := ctx.Param("dest")
		quantity, err := tc.sv.GetTotalTickets(ctx, destination)
		if err != nil {
			if errors.Is(err, tickets.ErrServiceTicketsNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{"message": "no tickets found"})
				return
			}
			if errors.Is(err, tickets.ErrServiceTicketsDbEmpty) {
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "empty database"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "success", "quantity": quantity})
	}
}

func (tc *TicketsController) GetAverage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		destination := ctx.Param("dest")
		average, err := tc.sv.AverageDestination(ctx, destination)
		if err != nil {
			if errors.Is(err, tickets.ErrServiceTicketsNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{"message": "no tickets found"})
				return
			}
			if errors.Is(err, tickets.ErrServiceTicketsDbEmpty) {
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "empty database"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "success", "average": average})
	}

}
