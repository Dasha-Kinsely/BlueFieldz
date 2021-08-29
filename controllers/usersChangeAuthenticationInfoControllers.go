package controllers

import (
	"log"

	"github.com/Dasha-Kinsely/leaveswears/helpers/errorresponders"
	"github.com/gin-gonic/gin"
)

func UsersChangeAuthenticationInfoControllers(c *gin.Context){
	// Get from form the fields you wish to change
	parsedField := c.Param("field")
	if parsedField == "" {
		errorresponders.ContextJSON(c, "form format")
		return
	}
	fetchedUser, exists := c.Get("auth_user")
	log.Println("Fetched this guy:", fetchedUser)
	if exists != true {
		errorresponders.ContextJSON(c, "unauthorized")
		return
	}
	return
}

