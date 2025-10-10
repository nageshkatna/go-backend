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

// LoginUser godoc
//
//	@Summary		Login
//	@Description	Login user by email and password.
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.LoginRequest	true	"Email and Password"
//	@Success		200		{object}	dto.LoginResponse	"Token and UserId"
//	@Failure		400		{object}	dto.ErrorResponse	"Bad request"
//	@Failure		404		{object}	dto.ErrorResponse	"User Unauthorized"
//	@Failure		500		{object}	dto.ErrorResponse	"Internal Server Error"
//	@Router			/user/login [post]
func (*UserController) LoginHandler(c *gin.Context) {
	req := &dto.LoginRequest{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}
	fmt.Println("Login request received for user:", req.Email)
	us := services.NewUserService()

	users, err := us.GetUserByEmail(req)

	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Error: err.Error()})
		return
	}

	roleId := (*users[0].UserRoles)[0].RoleId

	token, err := helper.GenerateJWT(req.Password, users[0].Password, roleId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to generate token. User Unauthorized"})
		return
	}

	usersResponse := &dto.LoginResponse{
		Token:  token,
		UserId: users[0].Id.String(),
	}

	c.JSON(http.StatusAccepted, &usersResponse)
}

// RegisterUser godoc
//
//	@Summary		Register
//	@Description	Register user by first name, last name, email and password.
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.RegisterUserRequest	true	"Email and Password"
//	@Success		201		{object}	dto.LoginResponse		"Token and UserId"
//	@Failure		400		{object}	dto.ErrorResponse		"Bad request"
//	@Failure		409		{object}	dto.ErrorResponse		"User Unauthorized"
//	@Failure		500		{object}	dto.ErrorResponse		"Internal Server Error"
//	@Router			/user/register [post]
func (*UserController) RegisterHandler(c *gin.Context) {
	req := &dto.RegisterUserRequest{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	us := services.NewUserService()
	res, errs := us.CreateUser(req, &[]models.UserRole{{RoleId: 1}}) // Default role as 'admin' with RoleId 1
	if errs != nil {
		c.JSON(http.StatusConflict, dto.ErrorResponse{Error: errs.Error()})
		return
	}

	token, err := helper.GenerateJWT(req.Password, res.UserId, 1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Failed to generate token"})
		return
	}
	res.Token = token

	c.JSON(http.StatusCreated, &res)
}
