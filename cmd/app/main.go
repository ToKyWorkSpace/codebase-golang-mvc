package main

import (
	"codebase-golang/internal/app/config"
	"codebase-golang/internal/app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	port := config.GetEnv("PORT")

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	routes.CommonRoute(r)
	routes.UserRoute(r)

	r.Run(":" + port)
}
