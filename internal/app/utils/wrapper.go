package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, data any, msg string) {
	c.JSON(http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": msg,
			"data":    data,
		},
	)
}

func Error(c *gin.Context, code int) {
	c.JSON(code,
		gin.H{
			"status":  code,
			"message": "Invalid request",
			"data":    nil,
		},
	)
}
