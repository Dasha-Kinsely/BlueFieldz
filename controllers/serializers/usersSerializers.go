package serializers

import (
	"github.com/Dasha-Kinsely/leaveswears/helpers/processes"
)

// --------------------- Serializers for Signup ----------------------------------

type SignupSuccessSerializer struct {
	Username string
	Email string
	ID uint	
}

type SignupSuccessResponse struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func (serializer *SignupSuccessSerializer) Response() SignupSuccessResponse {
	return SignupSuccessResponse{
		Username: serializer.Username,
		Email: serializer.Email,
		Token: processes.GenerateJWTTokenDefault(serializer.ID),
	}
}

// -----------------------------------------------------------------------------

type SigninSuccessSerializer struct {
	Username string
	Email string
	ID uint
}

type SigninSuccessResponse struct {
	Username string `json:"username"`
	Email string `json:"email"`
	ID uint `json:"id"`
}

func (serializer *SigninSuccessSerializer) Response() SigninSuccessResponse {
	return SigninSuccessResponse{
		Username: serializer.Username,
		Email: serializer.Email,
		ID: serializer.ID,
	}
}