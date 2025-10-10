package api

import (
	"go-backend/api/router"
	_ "go-backend/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServer() {
	r := gin.Default()

	api := r.Group("/api")
	v1 := api.Group("/v1")
	user := v1.Group("/user")
	dashboard := v1.Group("/dashboard")

	router.UserRoutes(user)
	router.DashboardRoutes(dashboard)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":5000") // listen and serve on 5000
}
