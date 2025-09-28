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

type FetchUserWithRole struct {
	UserBaseRequest
	RoleName string `json:"roleName" binding:"required"`
}

type FetchUserRoleWithPaginatedResponse struct {
	Users []FetchUserWithRole `json:"users"`
	PaginatedResponse
}

type Pagination struct {
	Page int `json:"page" binding:"required"`
	PageSize int `json:"pageSize" binding:"required"`
}

type PaginatedResponse struct {
	Pagination
	TotalRecords int64 `json:"totalRecords" binding:"required"`
	TotalPages uint `json:"totalPages" binding:"required"`
}