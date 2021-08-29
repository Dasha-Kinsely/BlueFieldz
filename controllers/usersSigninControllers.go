package controllers

import (
	"net/http"

	"github.com/Dasha-Kinsely/leaveswears/controllers/serializers"
	"github.com/Dasha-Kinsely/leaveswears/helpers/errorresponders"
	"github.com/Dasha-Kinsely/leaveswears/helpers/middlewares"
	"github.com/Dasha-Kinsely/leaveswears/helpers/processes"
	"github.com/Dasha-Kinsely/leaveswears/models"
	"github.com/gin-gonic/gin"
)

func UsersSigninControllers(c *gin.Context){
	// Checks whether this is a signin using email or using username
	var identifier, option string
	if len(c.PostForm("email")) < 1 && len(c.PostForm("username")) != 0 {
		identifier = c.PostForm("username")
		option = "username"
	} else if len(c.PostForm("username")) < 1 && len(c.PostForm("email")) != 0 {
		identifier = c.PostForm("email")
		option = "email"
	} else {
		errorresponders.ContextJSON(c, "illegal query")
		return
	}
	// Checks if user exists, condition creates a temporary custom type struct
	fetchCondition := []byte("{\""+option+"\":\""+identifier+"\"}")
	condition := processes.CustomTypeStruct(fetchCondition)
	fetchedUser, err := models.FindOneUser(*condition)
	if err != nil{
		errorresponders.ContextJSON(c, "user not found")
		return
	}
	// Checks if user's password is correct
	if err := processes.DecryptPassword(c.PostForm("password"), fetchedUser.PasswordHash); err!= nil{
		errorresponders.ContextJSON(c, "wrong password")
		return
	}
	// Update *gin.Context and serialize result as response
	middlewares.UpdateContextAuthUser(c, fetchedUser)
	serializer := serializers.UniversalSerializer{C: c}
	c.JSON(http.StatusAccepted, gin.H{"authenticated_user":serializer.SigninSuccessResponse()})
}