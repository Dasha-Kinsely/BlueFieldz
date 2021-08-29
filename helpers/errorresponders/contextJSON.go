package errorresponders

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ContextJSON(c *gin.Context, msg string) {
	switch msg{
	case "form format":
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": -1,
			"msg": "Form binding issues occurred due to invalid form format for this action!",
		})
	case "database saving":
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": -1,
			"msg": "Failed to save into database due to violation of constraints!",
		})
	case "user not found":
		c.JSON(http.StatusForbidden, gin.H{
			"status": -1,
			"msg": "User does not exist!",
		})
	case "incorrect password":
		c.JSON(http.StatusForbidden, gin.H{
			"status": -1,
			"msg": "Invalid password!",
		})
	case "illegal query":
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"msg": "Must use either email or username to authenticate!",
		})
	case "unauthorized":
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": -1,
			"msg": "you must first login in order to perform this action!", 
		})
	case "token not found":
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": -1,
			"msg": "token not found, request to this content is invalid",
		})
	default:
		c.JSON(http.StatusNotImplemented, gin.H{
			"status": -1,
			"msg": "Unknown error occurred at server side!",
		})
	}
}