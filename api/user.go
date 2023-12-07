package api

import (
	"diagnofish/model"
	"diagnofish/service"
	"fmt"
	"net/http"
	"net/mail"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type UserAPI interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
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
		fmt.Println(err.Error())
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

func (u *userAPI) Login(c *gin.Context) {
	var user model.UserLogin

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid decode json"))
		return
	}

	var recordUser = model.User{
		Email:    user.Email,
		Password: user.Password,
	}

	token, err := u.userService.Login(&recordUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(err.Error()))
		return
	}

	cookie := &http.Cookie{
		Name:     "session_token",
		Value:    *token,
		Expires:  time.Now().Add(200 * time.Minute),
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}

	if _, err := c.Request.Cookie("session_token"); err == nil {
		http.SetCookie(c.Writer, cookie)
	} else {
		http.SetCookie(c.Writer, cookie)
	}

	var claims model.Claims
	_, err = jwt.ParseWithClaims(*token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(model.JwtKey), nil
	})

	response := gin.H{
		"user_id": claims.Email,
		"message": "login success",
	}

	c.JSON(http.StatusOK, response)
}
