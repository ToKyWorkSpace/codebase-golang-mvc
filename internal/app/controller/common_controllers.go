package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": nil, "status": 200, "msg": "Your service work properly"})
}
