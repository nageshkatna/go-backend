package controllers

import (
	"fmt"
	"go-backend/api/dto"
	"go-backend/api/services"
	"go-backend/helper"
	"go-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (h *UserController) LoginHandler(c *gin.Context) {
	req := &dto.LoginRequest{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("Login request received for user:", req.Email)
	us := services.NewUserService()

	users, err := us.GetUserByEmail(req)
	if( err != nil){
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	token, err := helper.GenerateJWT(req.Password, users[0].Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	usersResponse := &dto.LoginResponse{
		Token:  token,
		UserId: users[0].Id.String(),
	}

	c.JSON(http.StatusAccepted, &usersResponse)
}

func (h *UserController) RegisterHandler(c *gin.Context) {
	req := &dto.RegisterUserRequest{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	us := services.NewUserService()
	res, errs := us.CreateUser(req, &[]models.UserRole{{RoleId: 1}}) // Default role as 'admin' with RoleId 1
	if errs != nil {
		c.JSON(http.StatusConflict, gin.H{"error": errs.Error()})
		return
	}

	token, err := helper.GenerateJWT(req.Password, res.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	res.Token = token

	c.JSON(http.StatusCreated, &res)
}

func (uc *UserController) InviteUser(c *gin.Context) {
	req := &dto.InviteUserRequest{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
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