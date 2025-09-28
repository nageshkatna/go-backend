package controllers

import (
	"go-backend/api/dto"
	"go-backend/api/services"
	"go-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DashboardController struct {}

func NewDashboardController() *DashboardController {
	return &DashboardController{}
}
func (*DashboardController) ListAllUsers(c *gin.Context) {
	
	req := &dto.Pagination{}
	err := c.ShouldBindJSON(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	us := services.NewUserService()

	response, errs := us.GetPaginatedUser(req)
	if errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : errs.Error()})
		return
	}

	c.JSON(http.StatusAccepted, response)
}

func (uc *DashboardController) InviteUser(c *gin.Context) {
	req := &dto.InviteUserRequest{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	us := services.NewUserService()

	user := &dto.RegisterUserRequest{
		UserBaseRequest: dto.UserBaseRequest{
			FirstName: req.FirstName,
			LastName: req.LastName,
			Email: req.Email,
		},
		Password: ";ash#2asdf84333as!@9-9/SS",
	}

	role := &[]models.UserRole{{RoleId: req.RoleId}}

	_, errs := us.CreateUser(user, role)
	if(errs != nil) {
		c.JSON(http.StatusConflict, gin.H{"error": errs.Error()})
		return
	}

	//Send Invite email by using Go Routine
	
	c.JSON(http.StatusCreated, gin.H{"message": "Invite Sent"})
	
}