package controllers

import (
	"fmt"
	"go-backend/api/dto"
	"go-backend/api/services"
	"go-backend/helper"
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

	users, err := us.AuthenticateUser(req)
	if( err != nil){
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}
	
	token, err := helper.GenerateJWT(req.Password, users)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	usersResponse := &dto.LoginResponse{
		Token:  token,
		UserId: users,
	}

	c.JSON(http.StatusAccepted, &usersResponse)
}