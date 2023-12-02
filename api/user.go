package api

import (
	"diagnofish/model"
	"diagnofish/service"
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
)

type UserAPI interface {
	Register(c *gin.Context)
}

type userAPI struct {
	userService service.UserService
}

func NewUserAPI(userService service.UserService) *userAPI {
	return &userAPI{userService}
}

func (u *userAPI) Register(c *gin.Context) {
	var user model.UserRegister

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid decode json"))
		return
	}

	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid email"))
		return
	}

	if user.Password != user.RepeatPassword {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("password not match"))
		return
	}

	var recordUser = model.User{
		Email:    user.Email,
		Password: user.Password,
	}

	recordUser, err = u.userService.Register(&recordUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, model.NewSuccessResponse("register success"))
}
