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
	fetchedClaim, yes := c.Get("claims")
	log.Println("Fetched this guy: ", fetchedUser)
	log.Println("with claims: ", fetchedClaim)
	if yes != true {
		errorresponders.ContextJSON(c, "unauthorized")
	}
	if exists != true {
		errorresponders.ContextJSON(c, "unauthorized")
		return
	}
	return
}

