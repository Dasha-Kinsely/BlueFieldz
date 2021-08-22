package controllers

import (
	"net/http"

	"github.com/Dasha-Kinsely/leaveswears/controllers/serializers"
	"github.com/Dasha-Kinsely/leaveswears/helpers/errorresponders"
	"github.com/Dasha-Kinsely/leaveswears/helpers/validators"
	"github.com/Dasha-Kinsely/leaveswears/models"
	"github.com/gin-gonic/gin"
)	

func UsersSignUpControllers(c *gin.Context){
	// Validation process to ensure the data being saved is clean and secure
	newUser := validators.NewUserSignUpValidator()
	if err := newUser.BindContext(c); err != nil {
		errorresponders.ContextJSON(c, "form format")
		return
	}
	// Validation process to ensure nothing happened during the actual dao processes
	if err := models.SaveOneUser(&newUser.ComparedTo); err != nil {
		errorresponders.ContextJSON(c, "database saving")
		return
	}
	serializer := serializers.SignupSuccessSerializer{
		Username: newUser.ComparedTo.Username,
		Email: newUser.ComparedTo.Email,
		ID: newUser.ComparedTo.ID,
	}
	c.Set("registered_user", serializer)
	c.JSON(http.StatusAccepted, gin.H{"registered": serializer.Response()})
}