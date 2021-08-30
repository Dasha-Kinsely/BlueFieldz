package routers

import "github.com/gin-gonic/gin"

func InitializeRoutes(r *gin.Engine) {
	base := r.Group("/api")
	UsersRoutes(base.Group("/users"))
}