package routers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func UsersMaster(router *gin.RouterGroup) {
	router.POST("/signup", UsersSignUp)
	router.POST("/signin", UsersSignIn)
	router.GET("/current", CurrentUser)
}

func UsersSignUp(c *gin.Context) {
	
}