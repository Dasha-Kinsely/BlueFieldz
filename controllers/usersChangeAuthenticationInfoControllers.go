package controllers

import (
	"net/http"

	"github.com/Dasha-Kinsely/leaveswears/helpers/dao"
	"github.com/Dasha-Kinsely/leaveswears/helpers/errorresponders"
	"github.com/Dasha-Kinsely/leaveswears/helpers/processes"
	"github.com/gin-gonic/gin"
)

func HandleNonHashedUpdate(c *gin.Context, condition []byte) {
	user, err := dao.GetUserByClaim(c)
	if err != nil {
		errorresponders.ContextJSON(c, "user not found")
		c.Abort()
		return
	}
	match := processes.CustomTypeStruct(condition)
	err2 := user.UpdateOneUser(match)
	if err2 != nil{
		errorresponders.ContextJSON(c, "database saving")
		c.Abort()
		return
	}
	// Note: successful changes on authinfo will automatically sign you out by invalidating jwt
	c.Set("claims", nil)
	c.Redirect(http.StatusTemporaryRedirect, "/api/users/signin")
}

func UsersChangeAuthenticationInfoControllers(c *gin.Context) {
	// Get from form the fields you wish to change
	parsedField := c.Param("field")
	parsedQuery := c.Query("value")
	if parsedQuery == "" {
		errorresponders.ContextJSON(c, "form format")
		return
	}
	var buildCondition []byte
	switch parsedField {
		case "username":
			buildCondition = []byte("{\"Username\": \""+parsedQuery+"\"}")
			HandleNonHashedUpdate(c, buildCondition)
			return
		case "email":
			buildCondition = []byte("{\"Email\": \""+parsedQuery+"\"}")
			HandleNonHashedUpdate(c, buildCondition)
			return
		case "password":
			newHash := processes.EncryptPassword(parsedQuery)
			buildCondition = []byte("{\"PasswordHash\": \""+newHash+"\"}")
			HandleNonHashedUpdate(c, buildCondition)
			return
		default:
			errorresponders.ContextJSON(c, "form format")
			return
	}
}

