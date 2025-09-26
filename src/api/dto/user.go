package dto

type LoginRequest struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
	UserId string `json:"user_id"`
}

type UserBaseRequest struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName string `json:"lastName" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type InviteUserRequest struct {
	UserBaseRequest
	RoleId uint `json:"roleId" binding:"required"`
}

type RegisterUserRequest struct {
	UserBaseRequest
	Password string `json:"password" binding:"required"`
}