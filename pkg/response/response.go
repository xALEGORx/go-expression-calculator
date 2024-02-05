package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Code  int         `json:"code"`
	Error interface{} `json:"error"`
}

func Data(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &SuccessResponse{
		Code: http.StatusOK,
		Data: data,
	})
}

func BadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, &ErrorResponse{
		Code:  http.StatusBadRequest,
		Error: message,
	})
}

func InternalServerError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, &ErrorResponse{
		Code:  http.StatusInternalServerError,
		Error: message,
	})
}
