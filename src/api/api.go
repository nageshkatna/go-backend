package api

import (
	"go-backend/api/router"

	"github.com/gin-gonic/gin"
)

func InitServer() {
	r := gin.Default()

	v1 := r.Group("/v1")
	user := v1.Group("/user")
	dashboard := v1.Group("/dashboard")

	router.UserRoutes(user)
	router.DashboardRoutes(dashboard)

	r.Run(":5000") // listen and serve on 5000
}