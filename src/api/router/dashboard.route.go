package router

import (
	"go-backend/api/controllers"
	"go-backend/api/middleware"

	"github.com/gin-gonic/gin"
)

func DashboardRoutes(router *gin.RouterGroup) {
	c := controllers.NewDashboardController()
	router.Use(middleware.AuthencticateRequest())

	router.POST("/invite-user", middleware.AuthorizeRequest([]string{"Admin","Manager"}), c.InviteUser)
	router.GET("/listAllUsers", c.ListAllUsers)
}