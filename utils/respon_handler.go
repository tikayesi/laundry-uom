package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseWrapper struct {
	Message      string
	ResponseCode int
	Result       interface{}
}

func HandleSuccess(context *gin.Context, data interface{}, message string) {
	response := ResponseWrapper{
		Message:      message,
		ResponseCode: 200,
		Result:       data,
	}
	context.JSON(http.StatusOK, response)
}

func HandleSuccessCreated(context *gin.Context, data interface{}, message string) {
	response := ResponseWrapper{
		Message:      message,
		ResponseCode: 201,
		Result:       data,
	}
	context.JSON(http.StatusCreated, response)
}

func HandleNotFound(context *gin.Context, message string) {
	response := ResponseWrapper{
		Message:      message,
		ResponseCode: 404,
		Result:       nil,
	}
	context.JSON(http.StatusNotFound, response)
}

func HandleInternalServerError(context *gin.Context, message string) {
	response := ResponseWrapper{
		Message:      message,
		ResponseCode: 500,
		Result:       nil,
	}
	context.JSON(http.StatusInternalServerError, response)
}

func HandleBadRequest(context *gin.Context, message string) {
	response := ResponseWrapper{
		Message:      message,
		ResponseCode: 400,
		Result:       nil,
	}
	context.JSON(http.StatusBadRequest, response)
}
