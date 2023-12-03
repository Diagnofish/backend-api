package middleware

import (
	"diagnofish/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		cookie, err := ctx.Request.Cookie("session_token")
		if err != nil {
			if ctx.GetHeader("Content-Type") == "application/json" {
				ctx.JSON(http.StatusUnauthorized, model.NewErrorResponse("unauthorized"))
			} else {
				ctx.Redirect(http.StatusSeeOther, "/user/login")
			}
		}

		tokenStr := cookie.Value
		claims := &model.Claims{}

		token, err := jwt.ParseWithClaims(tokenStr, claims,
			func(token *jwt.Token) (interface{}, error) {
				return model.JwtKey, nil
			})
		if err != nil {
			// error when parsing token
			if err == jwt.ErrSignatureInvalid {
				ctx.JSON(http.StatusUnauthorized, model.NewErrorResponse("invalid token"))
			} else {
				ctx.JSON(http.StatusBadRequest, model.NewErrorResponse(err.Error()))
			}

			ctx.Abort()
			return
		}

		if claims, ok := token.Claims.(*model.Claims); ok && token.Valid {
			ctx.Set("email", claims.Email)
		} else {
			ctx.JSON(http.StatusUnauthorized, model.NewErrorResponse("unauthorized"))

			ctx.Abort()
			return
		}

		ctx.Next()
	})
}
