package routes

import (
	"codebase-golang/internal/app/controller"

	"github.com/gin-gonic/gin"
)

func CommonRoute(r *gin.Engine) {
	r.GET("/", controller.HealthCheck)
}
