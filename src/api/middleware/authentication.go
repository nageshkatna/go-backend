package middleware

import (
	"fmt"
	"go-backend/api/services"
	"go-backend/helper"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthencticateRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")
		token := strings.Split(auth, " ")
		
		isValid, err, claim := helper.VeifyToken(token[1])
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Error in verify token %v", err)})
			ctx.Abort()
			return
		} else if !isValid{
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
			ctx.Abort()
			return 
		}

		ctx.Set("userId", claim.Issuer)
		ctx.Set("roleId", claim.RoleId)
		ctx.Next()
	}
}

func AuthorizeRequest(requiredRoles []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		roleIdRaw, exists := ctx.Get("roleId")
		if !exists {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "userId wasn't provided"})
			ctx.Abort()
			return
		} 
		roleId, ok := roleIdRaw.(uint)

		if ok {
			rs := services.NewRoleService{}
			role, err := rs.GetRoleById(uint(roleId))
			if err != nil {
				ctx.JSON(http.StatusForbidden, gin.H{"error": fmt.Sprintf("Error in finding the role from the provided RoleId %v", err)})
				ctx.Abort()
				return
			}

			authorized := false
			for _, r := range requiredRoles {
				if r == role[0].Name {
					authorized = true
					return
				}
			}

			if !authorized {
				ctx.JSON(http.StatusForbidden, gin.H{"error": "Provided role is not permitted to run this operation"})
				ctx.Abort()
				return
			}

			ctx.Next()
		}
	}
}