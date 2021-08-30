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
}

func (serializer *UniversalSerializer) SignupSuccessResponse() SignupSuccessResponse {
	res := serializer.C.MustGet("registered_user").(models.User)
	return SignupSuccessResponse{
		Username: res.Username,
		Email: res.Email,
	}
}

// ------------------------ Signin -------------------------------------
type SigninSuccessResponse struct {
	ID uint `json:"id"`
	Username string `json:"username"`
	SigninTime time.Time
	Token string `json:"token"`
}

func (serializer *UniversalSerializer) SigninSuccessResponse() SigninSuccessResponse {
	res := serializer.C.MustGet("auth_user").(models.User)
	return SigninSuccessResponse{
		ID: res.ID,
		Username: res.Username,
		SigninTime: time.Now(),
		Token: middlewares.GenerateJWTTokenDefault(res.ID, res.Username),
	}
}