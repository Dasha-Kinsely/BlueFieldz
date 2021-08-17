package validators

import (
	"log"

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
	ComparedTo models.User `json:"-"`
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
		log.Println(err)
		return err
	}
	m.ComparedTo.Username = m.User.Username
	m.ComparedTo.Email = m.User.Email
	m.ComparedTo.PasswordHash = processes.EncryptPassword(m.User.Password)
	return nil
}
//---------------------------------------------------------------------------------------------------------
// Sign In Validation
type UserEmailSigninValidator struct {
	User struct {
		Email string `form:"email" json:"email"`
		Password string `form:"password" json:"password" binding:"required,min=8,max=255"`
	} `json:"user"`
	fetchModel models.User `json:"-"`
}

type UserUsernameSigninValidator struct {
	User struct {
		Username string `form:"username" json:"username"`
		Password string `form:"password" json:"password" binding:"required,min=8,max=255"`
	} `json:"user"`
	fetchModel models.User `json:"-"`
}

func NewUserEmailSigninValidator() UserEmailSigninValidator {
	newSignin := UserEmailSigninValidator{}
	return newSignin
}

func NewUserUsernameSigninValidator() UserUsernameSigninValidator {
	newSignin := UserUsernameSigninValidator{}
	return newSignin
}

func (m *UserEmailSigninValidator) BindContext(c *gin.Context) error {
	// Intelligent way to select binding engine and to complete binding
	headerMethod := c.Request.Method
	headContentType := c.ContentType()
	if err:=c.ShouldBindWith(m, binding.Default(headerMethod, headContentType)); err != nil {
		log.Println(err)
		return err
	}
	m.fetchModel.Email = m.User.Email
	return nil
}

func (m *UserUsernameSigninValidator) BindContext(c *gin.Context) error {
	// Intelligent way to select binding engine and to complete binding
	headerMethod := c.Request.Method
	headContentType := c.ContentType()
	if err:=c.ShouldBindWith(m, binding.Default(headerMethod, headContentType)); err != nil {
		log.Println(err)
		return err
	}
	m.fetchModel.Username = m.User.Username
	return nil
}

//---------------------------------------------------------------------------------------------------------