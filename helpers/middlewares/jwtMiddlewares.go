package middlewares

import (
	"log"

	"github.com/Dasha-Kinsely/leaveswears/helpers/processes"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var SignitureKey string = processes.LoadEnv("getting signiture key", "JWT_SIGNITURE_KEY")

type JwtBase struct {
	JwtSigniture []byte
}

type CustomClaims struct {
	IsAdmin int8 `json:role`
	jwt.StandardClaims
}

func GetSignitureKey() string {
	return SignitureKey
}

func SetSignitureKey(key string) string {
	SignitureKey = key
	return SignitureKey
}

func JWTAuthMiddleware(requireAuth bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println(requireAuth)
	}
}