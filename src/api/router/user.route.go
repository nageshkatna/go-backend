package router

import (
	"go-backend/api/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup) {
	c := controllers.NewUserController()
	router.POST("/login", c.LoginHandler)
}