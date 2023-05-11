package server

import (
	"errors"
	"fmt"
	"github.com/NachofraML/bootcamp-go-w8/05-go-testing/clase-03/test-funcionales/vivo/prey"
	"github.com/NachofraML/bootcamp-go-w8/05-go-testing/clase-03/test-funcionales/vivo/shark"
	"github.com/NachofraML/bootcamp-go-w8/05-go-testing/clase-03/test-funcionales/vivo/simulator"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	ErrUnprocessableEntity = errors.New("invalid request")
)

type Handler struct {
	shark shark.Shark
	prey  prey.Prey
}

func NewHandler(shark shark.Shark, prey prey.Prey) *Handler {
	return &Handler{shark: shark, prey: prey}
}

// PUT: /v1/shark

func (h *Handler) ConfigureShark() gin.HandlerFunc {
	type request struct {
		XPosition float64 `json:"x_position"`
		YPosition float64 `json:"y_position"`
		Speed     float64 `json:"speed"`
	}
	type response struct {
		Success bool `json:"success"`
	}

	return func(context *gin.Context) {
		var req request
		if err := context.ShouldBindJSON(&req); err != nil {
			fmt.Println(err)
			context.JSON(http.StatusUnprocessableEntity, gin.H{"message": ErrUnprocessableEntity.Error()})
			return
		}

		// This is ugly, but I have no time
		defaultMaxTimeToCatch := 35.4
		catchSimulator := simulator.NewCatchSimulator(defaultMaxTimeToCatch)
		whiteShark := shark.CreateWhiteShark(catchSimulator)
		h.shark = whiteShark

		position := [2]float64{req.XPosition, req.YPosition}
		h.shark.Configure(position, req.Speed)

		context.JSON(http.StatusOK, response{Success: true})
	}
}

// PUT: /v1/prey

func (h *Handler) ConfigurePrey() gin.HandlerFunc {
	type request struct {
		Speed float64 `json:"speed"`
	}
	type response struct {
		Success bool `json:"success"`
	}

	return func(context *gin.Context) {
		var req request
		if err := context.ShouldBindJSON(&req); err != nil {
			context.JSON(http.StatusUnprocessableEntity, gin.H{"message": ErrUnprocessableEntity.Error()})
			return
		}

		h.prey.SetSpeed(req.Speed)

		context.JSON(http.StatusOK, response{Success: true})
	}
}

// POST: /v1/simulate

func (h *Handler) SimulateHunt() gin.HandlerFunc {
	type response struct {
		Success bool    `json:"success"`
		Message string  `json:"message"`
		Time    float64 `json:"time"`
	}

	return func(context *gin.Context) {
		if h.shark == nil || h.prey == nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": "shark and prey must be configured"})
			return
		}

		err, huntTime := h.shark.Hunt(h.prey)
		if err != nil {
			if errors.Is(err, shark.ErrCouldNotHuntPreyTooFar) {
				context.JSON(http.StatusOK, response{Success: false, Message: "prey is too far, cannot hunt", Time: huntTime})
				return
			}
			if errors.Is(err, shark.ErrCouldNotHuntPreyTooFast) {
				context.JSON(http.StatusOK, response{Success: false, Message: "prey is faster than hunter, cannot hunt", Time: huntTime})
				return
			}
			context.JSON(http.StatusInternalServerError, gin.H{"message": "an error has occurred while hunting"})
			return
		}

		context.JSON(http.StatusOK, response{Success: true, Message: "the prey has been hunted", Time: huntTime})
	}
}
