package serializers

import (
	"github.com/Dasha-Kinsely/leaveswears/helpers/processes"
	"github.com/Dasha-Kinsely/leaveswears/models"
	"github.com/gin-gonic/gin"
)

type SignupSuccessSerializer struct {
	c *gin.Context
}

type SignupSuccessResponse struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func (self *SignupSuccessSerializer) Response() SignupSuccessResponse {
	responseModel := self.c.MustGet("signing_up").(models.User)
	return SignupSuccessResponse{
		Username: responseModel.Username,
		Email: responseModel.Email,
		Token: processes.GenerateJWTTokenDefault(responseModel.ID),
	}
}