package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Data interface{} `json:"data"`
}

func SendErrorResponse(ctx *gin.Context, status int, code string, message string) {
	errResponse := ErrorResponse{
		Status:  status,
		Code:    code,
		Message: message,
	}
	ctx.JSON(status, errResponse)
}

func SendResponse(ctx *gin.Context, data interface{}) {
	response := Response{
		Data: data,
	}
	ctx.JSON(http.StatusOK, response)
}
