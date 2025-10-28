package helper

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetIntParams(c *gin.Context, name string, def int) int {
	val, err := strconv.Atoi(c.Param(name))
	if err != nil {
		return def
	}
	return val
}
