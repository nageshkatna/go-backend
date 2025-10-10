package controllers

import (
	"go-backend/api/dto"
	"go-backend/api/services"
	"go-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DashboardController struct{}

func NewDashboardController() *DashboardController {
	return &DashboardController{}
}

//	@BasePath		/dashboard
//
// ListUser godoc
//
//	@Summary		List All Users
//	@Description	Get All Users
//	@Tags			Dashboard
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.Pagination							true	"Pagination info"
//	@Success		202		{object}	dto.FetchUserRoleWithPaginatedResponse	"List of All Users with roles"
//	@Failure		400		{object}	dto.ErrorResponse						"Bad request"
//	@Security		BearerAuth
//	@Router			/dashboard/listAllUsers [post]
func (*DashboardController) ListAllUsers(c *gin.Context) {

	req := &dto.Pagination{}
	err := c.ShouldBindJSON(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	us := services.NewUserService()

	response, errs := us.GetPaginatedUser(req)
	if errs != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: errs.Error()})
		return
	}

	c.JSON(http.StatusAccepted, response)
}

// ListUser godoc
//
//	@Summary		Invite a user
//	@Description	Invite a user or Create a user by Admin or Manager.
//	@Tags			Dashboard
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.InviteUserRequest	true	"User Basic Info with role id"
//	@Success		200		{object}	dto.MessageResponse		"Generic response with message"
//	@Failure		400		{object}	dto.ErrorResponse		"Bad request"
//	@Security		BearerAuth
//	@Router			/dashboard/invite-user [post]
func (*DashboardController) InviteUser(c *gin.Context) {
	req := &dto.InviteUserRequest{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	us := services.NewUserService()

	user := &dto.RegisterUserRequest{
		UserBaseRequest: dto.UserBaseRequest{
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Email:     req.Email,
		},
		Password: ";ash#2asdf84333as!@9-9/SS",
	}

	role := &[]models.UserRole{{RoleId: req.RoleId}}

	_, errs := us.CreateUser(user, role)
	if errs != nil {
		c.JSON(http.StatusConflict, dto.ErrorResponse{Error: errs.Error()})
		return
	}

	//Send Invite email by using Go Routine

	c.JSON(http.StatusCreated, dto.MessageResponse{Message: "Invite Sent"})
}

// UpdateUser godoc
//
//	@Summary		Update a user by Id
//	@Description	Updates First Names, Last Name, Email and Role only. To update Role, need to provide role id. Authorized user like admin and manager can only delete the user.
//	@Tags			Dashboard
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.UpdateRequest	true	"User Basic Info parameter to update"
//	@Success		200		{object}	dto.MessageResponse	"Generic response with message"
//	@Failure		400		{object}	dto.ErrorResponse	"Bad request"
//	@Security		BearerAuth
//	@Router			/dashboard/updateUser [patch]
func (*DashboardController) UpdateUser(c *gin.Context) {
	req := dto.UpdateRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
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

	us := services.NewUserService()
	_, err := us.UpdateUser(req.UserId, user)

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.MessageResponse{Message: "User Updated Successfully!"})
}

// DeleteUser godoc
//
//	@Summary		Delete a user by Id
//	@Description	Delete a user by Id. Authorized user like admin and manager can only delete the user.
//	@Tags			Dashboard
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.RequestWithUserId	true	"User Id"
//	@Success		200		{object}	dto.MessageResponse		"Generic response with message"
//	@Failure		400		{object}	dto.ErrorResponse		"Bad request"
//	@Security		BearerAuth
//	@Router			/dashboard/deleteUser [delete]
func (*DashboardController) DeleteUser(c *gin.Context) {
	req := dto.RequestWithUserId{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	us := services.NewUserService()

	if err := us.DeleteUser(req.UserId); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, dto.MessageResponse{Message: "User Deleted Successfully"})
}
