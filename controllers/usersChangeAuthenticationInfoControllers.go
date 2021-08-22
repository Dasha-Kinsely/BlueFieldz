package controllers

import (
	"log"

	"github.com/Dasha-Kinsely/leaveswears/helpers/errorresponders"
	"github.com/gin-gonic/gin"
)

func UsersChangeAuthenticationInfoControllers(c *gin.Context){
	/*parsedIdentifier := c.Param("identifier")
	if parsedIdentifier == "" {
		errorresponders.ContextJSON(c, "form format")
		return
	}*/
	parsedField := c.Param("field")
	log.Println(parsedField)
	if parsedField == "" {
		errorresponders.ContextJSON(c, "form format")
		return
	}
	fetchedUser, exists := c.Get("auth_user")
	log.Println(fetchedUser)
	if exists != true {
		errorresponders.ContextJSON(c, "unauthorized")
		return
	}
	return
}

