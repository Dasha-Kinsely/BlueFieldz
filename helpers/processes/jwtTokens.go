package processes

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWTTokenDefault(uid uint) string {
	newToken := jwt.New(jwt.GetSigningMethod("HS256"))
	newToken.Claim= jwt.MapClaims{
		"id": uid,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}
	token, _ := newToken.SignedString([]byte(loadEnv("jwt secret step", "PASSWORD_SECRET")))
	return token
}