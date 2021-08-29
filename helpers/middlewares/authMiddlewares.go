package middlewares

import (
	"github.com/Dasha-Kinsely/leaveswears/models"
	"github.com/gin-gonic/gin"
)

// Simple update on auth_user state of gin.Context
func UpdateContextAuthUser(c *gin.Context, fetchedUser models.User) {
	c.Set("auth_user", fetchedUser)
}