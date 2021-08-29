package middlewares

import (
	"log"
	"time"

	"github.com/Dasha-Kinsely/leaveswears/helpers/errorresponders"
	"github.com/Dasha-Kinsely/leaveswears/helpers/processes"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var SignitureKey string = processes.LoadEnv("getting signiture key", "JWT_SIGNITURE_KEY")

type JwtBase struct {
	JwtSigniture []byte
}

type CustomClaims struct {
	ID uint `json:"id"`
	jwt.StandardClaims
}

func GetSignitureKey() string {
	return SignitureKey
}

func SetSignitureKey(key string) string {
	SignitureKey = key
	return SignitureKey
}

func NewJwtBase() *JwtBase {
	// Simply assign env variable signiture key into JwtBase object as byte array form.
	return &JwtBase{
		[]byte(GetSignitureKey()),
	}
}

func (j *JwtBase) ParseJwtToken(accessToken string) (claims *CustomClaims, err error) {
	// parse accessToken into CustomClaims object and bind using my secret env jwt signiture key
	parsedToken , err := jwt.ParseWithClaims(accessToken, &CustomClaims{}, func(parsedToken *jwt.Token) (interface{}, error) {
		return j.JwtSigniture, nil
	})
	log.Println("Token has raw string of: ", parsedToken.Raw)
	log.Println("Token has encryption method of: ", parsedToken.Method)
	log.Println("Token has signiture from issuer named: ", parsedToken.Signature)
	log.Println("Token is valid: ", parsedToken.Valid)
	log.Println("Token has header that matches: ", parsedToken.Header)
	if err != nil {
		if ve, issue := err.(*jwt.ValidationError); issue && ve.Errors != 0 {
			// If problems occurred while parsing request token, return no claim
			if jwt.ValidationErrorMalformed != 0{
				return nil, errorresponders.JwtError("malformed")
			} else if jwt.ValidationErrorExpired != 0 {
				return nil, errorresponders.JwtError("expired")
			} else {
				return nil, errorresponders.JwtError("default unknown")
			}
		}
	}
	// If claim is valid, bind it to CustomClaims
	if claims, ok := parsedToken.Claims.(*CustomClaims); ok && parsedToken.Valid {
		return claims, nil
	}
	return nil, errorresponders.JwtError("default unknown")
}

func GenerateJWTTokenDefault(uid uint) string {
	j := &JwtBase{
		[]byte(SignitureKey),
	}
	claims := CustomClaims{
		uid,
		jwt.StandardClaims {
			ExpiresAt: int64(time.Now().Unix() + 3600),
			Issuer: SignitureKey,
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString(j.JwtSigniture)
	if err != nil {
		log.Println(err.Error())
		errorresponders.JwtError("generation")
	} else {
		log.Println(token)
	}
	return token 
}

func JWTAuthMiddleware(requireAuth bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.Request.Header.Get("token")
		if accessToken == "" {
			errorresponders.ContextJSON(c, "token not found")
			c.Abort()
			return
		}
		j := NewJwtBase()
		claims, err := j.ParseJwtToken(accessToken)
		if err != nil {
			log.Println(err.Error())
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
}

