package routers

import (
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
	base := r.Group("/api")
	UsersMaster(base.Group("/users"))
}
