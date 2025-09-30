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

func (*DashboardController) InviteUser(c *gin.Context) {
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

func (*DashboardController) UpdateUser(c *gin.Context) {
	req := dto.UpdateRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
	}
	
	var user = dto.UpdateUserObj{}

	if req.FirstName != nil {
		user.FirstName = req.FirstName
	}

	if req.LastName != nil {
		user.LastName = req.LastName
	}

	if req.Email != nil {
		user.Email = req.Email
	}

	if req.RoleId != nil {
		user.RoleId = req.RoleId
	}

	// fmt.Printf("%+v", user)

	us := services.NewUserService()
	_, err := us.UpdateUser(req.UserId, user)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User Updated Successfully!"})
}

func (*DashboardController) DeleteUser(c *gin.Context) {
	req := dto.RequestWithUserId{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	us := services.NewUserService()

	if err := us.DeleteUser(req.UserId); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "User Deleted Successfully"})
}