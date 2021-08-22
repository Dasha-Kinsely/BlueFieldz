package controllers

import (
	"net/http"

	"github.com/Dasha-Kinsely/leaveswears/controllers/middlewares"
	"github.com/Dasha-Kinsely/leaveswears/helpers/errorresponders"
	"github.com/Dasha-Kinsely/leaveswears/helpers/processes"
	"github.com/Dasha-Kinsely/leaveswears/helpers/validators"
	"github.com/Dasha-Kinsely/leaveswears/models"
	"github.com/gin-gonic/gin"
)

func UsersSigninControllers(c *gin.Context){
	if len(c.PostForm("email")) < 1 && len(c.PostForm("username")) != 0 {
		// Checks whether the form's data format matches model (sign in via username)
		currentUser := validators.NewUserUsernameSigninValidator()
		if err := currentUser.BindContext(c); err != nil {
			errorresponders.ContextJSON(c, "form format")
			return
		}
		// Checks if user exists
		fetchedUser, err := models.FindOneUser(&models.User{Username: currentUser.User.Username})
		if err != nil{
			c.JSON(http.StatusForbidden, "user with the given username was not found!")
			return
		}
		// Checks if user's password is correct
		if err:= processes.DecryptPassword(currentUser.User.Password, fetchedUser.PasswordHash); err!= nil{
			c.JSON(http.StatusForbidden, "wrong password!")
			return
		}
		// Update Context and Serialize Result as Response
		middlewares.UpdateContextAuthUser(c, "username", fetchedUser.Username)
	} else if len(c.PostForm("username")) < 1 && len(c.PostForm("email")) != 0 {
		// Checks whether the form's data format matches model (sign in via email)
		currentUser := validators.NewUserEmailSigninValidator()
		if err := currentUser.BindContext(c); err != nil {
			errorresponders.ContextJSON(c, "form format")
			return
		}
		// Checks if user exists
		fetchedUser, err := models.FindOneUser(&models.User{Email: currentUser.User.Email})
		if err != nil{
			c.JSON(http.StatusForbidden, "user with the given username was not found!")
			return
		}
		// Checks if user's password is correct
		if err:= processes.DecryptPassword(currentUser.User.Password, fetchedUser.PasswordHash); err!= nil{
			c.JSON(http.StatusForbidden, "wrong password!")
			return
		}
		// Update Context and Serialize Result as Response
		middlewares.UpdateContextAuthUser(c, "email", fetchedUser.Email)
	} else {
		errorresponders.ContextJSON(c, "illegal query")
		return
	}
}