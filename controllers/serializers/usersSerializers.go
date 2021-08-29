package serializers

import (
	"time"

	"github.com/Dasha-Kinsely/leaveswears/helpers/middlewares"
	"github.com/Dasha-Kinsely/leaveswears/models"
)

// --------------------- Signup ----------------------------------

type SignupSuccessResponse struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func (serializer *UniversalSerializer) SignupSuccessResponse() SignupSuccessResponse {
	res := serializer.C.MustGet("registered_user").(models.User)
	return SignupSuccessResponse{
		Username: res.Username,
		Email: res.Email,
		Token: middlewares.GenerateJWTTokenDefault(res.ID),
	}
}

// ------------------------ Signin -------------------------------------
type SigninSuccessResponse struct {
	ID uint `json:"id"`
	SigninTime time.Time
	Token string `json:"token"`
}

func (serializer *UniversalSerializer) SigninSuccessResponse() SigninSuccessResponse {
	res := serializer.C.MustGet("auth_user").(models.User)
	return SigninSuccessResponse{
		ID: res.ID,
		SigninTime: time.Now(),
		Token: middlewares.GenerateJWTTokenDefault(res.ID),
	}
}