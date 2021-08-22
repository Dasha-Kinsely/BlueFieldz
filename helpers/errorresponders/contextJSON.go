package errorresponders

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ContextJSON(c *gin.Context, msg string) {
	switch msg{
	case "form format":
		c.JSON(http.StatusUnprocessableEntity, "Form binding issues occurred due to invalid form format for this action!")
	case "database saving":
		c.JSON(http.StatusUnprocessableEntity, "Failed to save into database due to violation of constraints!")
	case "user not found":
		c.JSON(http.StatusForbidden, "User does not exist!")
	case "incorrect password":
		c.JSON(http.StatusForbidden, "Invalid password!")
	case "illegal query":
		c.JSON(http.StatusBadRequest, "Must use either email or username to authenticate!")
	case "unauthorized":
		c.JSON(http.StatusUnauthorized, "you must first login in order to perform this action!")
	default:
		c.JSON(http.StatusNotImplemented, "Unknown error occurred at server side!")
	}
}