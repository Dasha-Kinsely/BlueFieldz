package routers

import (
	//"log"

	"github.com/Dasha-Kinsely/leaveswears/controllers"
	"github.com/gin-gonic/gin"
)

func UsersMaster(router *gin.RouterGroup) {
	router.POST("/signup", UsersSignup)
	router.POST("/signin", UsersSignin)
	/*router.GET("/current", CurrentUser)*/
}

func UsersSignup(c *gin.Context) {
	controllers.UsersSignUpControllers(c)
}

func UsersSignin(c *gin.Context) {
	controllers.UsersSigninControllers(c)
}