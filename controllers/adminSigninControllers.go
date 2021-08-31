package controllers

import (
	"github.com/Dasha-Kinsely/leaveswears/helpers/errorresponders"
	"github.com/Dasha-Kinsely/leaveswears/helpers/processes"
	"github.com/gin-gonic/gin"
)

func AdminSigninControllers(c *gin.Context) {
	// Note: the only admins were already assigned by env variables during initial migration step, no new admin may be introduced
	id := c.PostForm("id")
	password := c.PostForm("password")
	if err := processes.DecryptPassword(password, ); err != nil {
		errorresponders.ContextJSON(c, "wrong password")
		return
	}
	
}