package router

import (
	"go-backend/api/controllers"
	"go-backend/api/middleware"

	"github.com/gin-gonic/gin"
)

func DashboardRoutes(router *gin.RouterGroup) {
	c := controllers.NewDashboardController()
	router.Use(middleware.AuthencticateRequest())

	router.POST("/User", middleware.AuthorizeRequest([]string{"admin", "manager"}), c.InviteUser)
	router.POST("/inviteBulkUser", middleware.AuthorizeRequest([]string{"admin", "manager"}), c.InviteBulkUsers)
	router.GET("/User/:page/:pageSize", c.ListAllUsers)
	router.PATCH("/User", middleware.AuthorizeRequest([]string{"admin", "manager"}), c.UpdateUser)
	router.DELETE("/User", middleware.AuthorizeRequest([]string{"admin", "manager"}), c.DeleteUser)
}
