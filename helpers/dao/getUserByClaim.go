package dao

import (
	"github.com/Dasha-Kinsely/leaveswears/helpers/middlewares"
	"github.com/Dasha-Kinsely/leaveswears/models"
	"github.com/gin-gonic/gin"
)

func GetUserByClaim(c *gin.Context) (*models.User, error) {
	// Ensure jwt claims is valid
	claim := c.MustGet("claims").(*middlewares.CustomClaims)
	u := claim.Username
	fetchedUser, err := models.FindOneUser(&models.User{Username: u})
	if err != nil {
		return nil, err
	}
	return &fetchedUser, nil
}