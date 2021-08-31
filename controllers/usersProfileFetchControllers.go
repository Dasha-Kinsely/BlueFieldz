package controllers

import (
	"log"

	"github.com/Dasha-Kinsely/leaveswears/helpers/errorresponders"
	"github.com/Dasha-Kinsely/leaveswears/helpers/middlewares"
	"github.com/Dasha-Kinsely/leaveswears/models"
	"github.com/gin-gonic/gin"
)

func UsersProfileFetchControllers(c *gin.Context) {
	claim := c.MustGet("claims").(*middlewares.CustomClaims)
	user, err := models.GetProfileByUsername(claim.Username)
	if err != nil {
		errorresponders.ContextJSON(c, "user not found")
		c.Abort()
		return
	}
	log.Println(user)
}