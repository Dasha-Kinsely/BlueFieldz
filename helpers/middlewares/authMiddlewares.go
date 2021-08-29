package middlewares

import (
	"log"
	"net/http"

	"github.com/Dasha-Kinsely/leaveswears/controllers/serializers"
	"github.com/Dasha-Kinsely/leaveswears/models"
	"github.com/gin-gonic/gin"
)

// Simple update on auth_user state of gin.Context
func UpdateContextAuthUser(c *gin.Context, fetchedUser models.User) {
	c.Set("auth_user", fetchedUser)
	log.Println(c.Keys)
	serializer := serializers.UniversalSerializer{c}
	c.JSON(http.StatusAccepted, gin.H{"authenticated_user":serializer.SigninSuccessResponse()})
}