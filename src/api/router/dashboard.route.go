package router

import (
	"go-backend/api/controllers"
	"go-backend/api/middleware"

	"github.com/gin-gonic/gin"
)

func DashboardRoutes(router *gin.RouterGroup) {
	c := controllers.NewDashboardController()
	router.Use(middleware.AuthencticateRequest())

	router.POST("/invite-user", middleware.AuthorizeRequest([]string{"admin", "manager"}), c.InviteUser)
	router.POST("/listAllUsers", c.ListAllUsers)
	router.PATCH("/updateUser", middleware.AuthorizeRequest([]string{"admin", "manager"}), c.UpdateUser)
	router.DELETE("/deleteUser", middleware.AuthorizeRequest([]string{"admin", "manager"}), c.DeleteUser)
}
