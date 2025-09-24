package api

import (
	"go-backend/api/router"

	"github.com/gin-gonic/gin"
)

func InitServer() {
	r := gin.Default()

	api := r.Group("/api")
	user := api.Group("/user")
	router.UserRoutes(user)

	r.Run(":5000") // listen and serve on 5000
}