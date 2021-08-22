package routers

import (
	"log"

	"github.com/Dasha-Kinsely/leaveswears/controllers"
	"github.com/gin-gonic/gin"
)

func UsersMaster(router *gin.RouterGroup) {
	router.POST("/signup", UsersSignup)
	router.POST("/signin", UsersSignin)
	router.PUT("/change/:field", UpdateAuthInfo)
}

func UsersSignup(c *gin.Context) {
	controllers.UsersSignUpControllers(c)
}

func UsersSignin(c *gin.Context) {
	log.Println("registered: ", c.Keys)
	controllers.UsersSigninControllers(c)
}

func UpdateAuthInfo(c *gin.Context){
	log.Println("current authed: ", c.Keys)
	controllers.UsersChangeAuthenticationInfoControllers(c)
}