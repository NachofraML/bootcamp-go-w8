package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

func TokenMiddleware(c *gin.Context) {
	token := c.GetHeader("token")
	if token != "12345" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
		return
	}
	c.Next()
}

type middlewareResponse struct {
	Code          string        `json:"code"`
	StartDateTime time.Time     `json:"start_date_time"`
	EndDateTime   time.Time     `json:"end_date_time"`
	Endpoint      string        `json:"endpoint"`
	Size          []byte        `json:"size"`
	Duration      time.Duration `json:"duration"`
}

func LogResponseMiddleware(ctx *gin.Context) {
	startTime := time.Now()

	ctx.Next()

	responseSize := strconv.Itoa(ctx.Writer.Size())

	endTime := time.Now()

	middlewareResponse := middlewareResponse{
		Code:          ctx.Request.Method,
		StartDateTime: startTime,
		EndDateTime:   endTime,
		Endpoint:      ctx.Request.URL.Path,
		Size:          []byte(responseSize),
		Duration:      endTime.Sub(startTime),
	}

	log.Printf("END TIME: [%s], METHOD: %s, URL: %s, SIZE: %d, DURATION: %s\n",
		middlewareResponse.EndDateTime,
		middlewareResponse.Code,
		middlewareResponse.Endpoint,
		middlewareResponse.Size,
		middlewareResponse.Duration,
	)
}
