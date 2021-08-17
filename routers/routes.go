package routers

import (
	"github.com/gin-gonic/gin"
	//"productRoutes"
	"userRoutes"
)

func initializeRoutes(r *gin.Engine) {
	base := r.Group("/api")
	userRoutes.UsersMaster(base.Group("/users"))
}
