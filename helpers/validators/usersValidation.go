package validators

import (
	"github.com/Dasha-Kinsely/leaveswears/helpers/processes"
	"github.com/Dasha-Kinsely/leaveswears/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// Sign Up Validation
type UserSignupValidator struct {
	User struct {
		Username string `form:"username" json:"username" binding:"required,alphanum,max=255"`
		Email string `form:"email" json:"email" binding:"required,email"`
		Password string `form:"password" json:"password" binding:"required,min=8,max=255"`
	} `json:"user"`
	ValidatedNewUser models.User `json:"-"`
}

func NewUserSignUpValidator() UserSignupValidator {
	newUser := UserSignupValidator{}
	return newUser
}

func (m *UserSignupValidator) BindContext(c *gin.Context) error {
	// Intelligent way to select binding engine and to complete binding
	headerMethod := c.Request.Method
	headContentType := c.ContentType()
	if err:=c.ShouldBindWith(m, binding.Default(headerMethod, headContentType)); err != nil {
		return err
	}
	// Set user based on received
	m.ValidatedNewUser.Username = m.User.Username
	m.ValidatedNewUser.Email = m.User.Email
	m.ValidatedNewUser.PasswordHash = processes.EncryptPassword(m.User.Password)
	m.ValidatedNewUser.IsAdmin = false
	return nil
}
//---------------------------------------------------------------------------------------------------------

//---------------------------------------------------------------------------------------------------------