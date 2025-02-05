package routes

import (
	"codebase-golang/internal/app/config"
	"codebase-golang/internal/app/controller"
	"codebase-golang/internal/app/utils/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.Engine) {
	username := config.GetEnv("BASIC_AUTH_USERNAME")
	password := config.GetEnv("BASIC_AUTH_PASSWORD")

	users := r.Group("/v1/users", middleware.BasicAuth(username, password))
	{
		users.GET("/get/:id", controller.Get)
		users.GET("/get", controller.GetAll)
		users.POST("/add", controller.Add)
		users.PATCH("/update/:id", controller.Update)
		users.DELETE("/delete/:id", controller.Delete)
	}
}
