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

func (*UserController) LoginHandler(c *gin.Context) {
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

	roleId := (*users[0].UserRoles)[0].RoleId

	token, err := helper.GenerateJWT(req.Password, users[0].Password, roleId)
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

func (*UserController) RegisterHandler(c *gin.Context) {
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

	token, err := helper.GenerateJWT(req.Password, res.UserId, 1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	res.Token = token

	c.JSON(http.StatusCreated, &res)
}