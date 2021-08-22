package middlewares

import (
	//"log"
	"net/http"

	"github.com/Dasha-Kinsely/leaveswears/controllers/serializers"
	"github.com/Dasha-Kinsely/leaveswears/helpers/errorresponders"
	"github.com/Dasha-Kinsely/leaveswears/models"
	"github.com/gin-gonic/gin"
)

/*func RequireAuthentication(redir401 bool) gin.HandlerFunc{
	return func(c *gin.Context){

	}
}*/

// Simple update on auth_user state of gin.Context
func UpdateContextAuthUser(c *gin.Context, choiceString string, currentUser string) {
	if choiceString == "username"{
		currentUser, err := models.FindOneUser(&models.User{Username:currentUser})
		if err != nil {
			errorresponders.ContextJSON(c, "user not found")
			return
		}
		c.Set("auth_user", currentUser)
		serializer := serializers.SigninSuccessSerializer{
			Username: currentUser.Username,
			Email: currentUser.Email,
			ID: currentUser.ID,
		}
		//log.Println(c.Keys["auth_user"])
		c.JSON(http.StatusAccepted, gin.H{"user":serializer.Response()})		
	} else if choiceString == "email"{
		currentUser, err := models.FindOneUser(&models.User{Email:currentUser})
		if err != nil {
			errorresponders.ContextJSON(c, "user not found")
			return
		}
		c.Set("auth_user", currentUser)
		serializer := serializers.SigninSuccessSerializer{
			Username: currentUser.Username,
			Email: currentUser.Email,
			ID: currentUser.ID,
		}
		c.JSON(http.StatusAccepted, gin.H{"user":serializer.Response()})	
	} else {
		errorresponders.ContextJSON(c, "illegal query")
		return
	}
}