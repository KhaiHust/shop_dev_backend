package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AbortErrorHandler(ctx *gin.Context, code int) {
	errResponse := GetErrorResponse(code)
	ctx.JSON(
		errResponse.HTTPCode,
		gin.H{
			"code":    errResponse.ServiceCode,
			"message": errResponse.Message,
			"data":    nil,
		},
	)
}
func AbortErrorHandleCustomMessage(c *gin.Context, code int, message string) {
	errorResponse := GetErrorResponse(code)
	c.JSON(errorResponse.HTTPCode, gin.H{
		"code":    errorResponse.ServiceCode,
		"message": message,
		"data":    nil,
	})
}
func AbortErrorResponseHandle(c *gin.Context, errorResponse *ErrorResponse) {
	c.JSON(errorResponse.HTTPCode, gin.H{
		"code":    errorResponse.ServiceCode,
		"message": errorResponse.Message,
		"data":    nil,
	})
}
func SuccessfulHandle(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": http.StatusText(http.StatusOK),
		"data":    data,
	})
}
