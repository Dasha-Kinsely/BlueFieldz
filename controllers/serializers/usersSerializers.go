package serializers

import (
	"github.com/Dasha-Kinsely/leaveswears/helpers/processes"
	"github.com/gin-gonic/gin"
)

type SignupSuccessResponse struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func (c *gin.Context) Response() SignupSuccessResponse {
	responseModel := c.MustGet("signing_up")
	return SignupSuccessResponse{
		Username: responseModel.Username,
		Email: responseModel.Email,
		Token: processes.GenerateJWTTokenDefault(responseModel.ID),
	}
}