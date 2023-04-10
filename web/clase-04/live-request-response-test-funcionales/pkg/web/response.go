package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type errorResponse struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type response struct {
	Data interface{} `json:"data"`
}

func SendErrorResponse(ctx *gin.Context, status int, code string, message string) {
	errResponse := errorResponse{
		Status:  status,
		Code:    code,
		Message: message,
	}
	ctx.JSON(status, errResponse)
}

func SendResponse(ctx *gin.Context, data interface{}) {
	response := response{
		Data: data,
	}
	ctx.JSON(http.StatusOK, response)
}
