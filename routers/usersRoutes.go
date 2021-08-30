package routers

import (
	"github.com/Dasha-Kinsely/leaveswears/controllers"
	"github.com/Dasha-Kinsely/leaveswears/helpers/middlewares"
	"github.com/gin-gonic/gin"
)
func UsersSignup(c *gin.Context) {
	controllers.UsersSignUpControllers(c)
}

func UsersSignin(c *gin.Context) {
	controllers.UsersSigninControllers(c)
}

func UpdateAuthInfo(c *gin.Context){
	controllers.UsersChangeAuthenticationInfoControllers(c)
}

func GetProfileInfo(c *gin.Context){
	controllers.UsersProfileFetchControllers(c)
}

func UpdateProfileInfo(c *gin.Context){
	controllers.UsersProfileChangeControllers(c)
}

func UsersRoutes(router *gin.RouterGroup) {
	router.POST("/signup", UsersSignup)
	router.POST("/signin", UsersSignin)
	router.PUT("/signin", UsersSignin)
	// Auth required routes
	router.Use(middlewares.JWTAuthMiddleware(true))
	{
		router.PUT("/change/:field", UpdateAuthInfo)
		router.GET("/profile", GetProfileInfo)
		router.POST("/profile/:field", UpdateProfileInfo)
	}
}