package middleware

import (
	"diagnofish/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		cookie, err := c.Request.Cookie("session_token")
		if err != nil {
			if c.GetHeader("Content-Type") == "application/json" {
				c.JSON(http.StatusUnauthorized, model.NewErrorResponse("unauthorized"))
			} else {
				c.JSON(http.StatusUnauthorized, model.NewErrorResponse("unauthorized"))
				c.Redirect(http.StatusSeeOther, "/user/login")
			}
		}

		tokenStr := cookie.Value
		claims := &model.Claims{}

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return model.JwtKey, nil
		})
		if err != nil {
			// error when parsing token
			if err == jwt.ErrSignatureInvalid {
				c.JSON(http.StatusUnauthorized, model.NewErrorResponse("invalid token"))
			} else {
				c.JSON(http.StatusBadRequest, model.NewErrorResponse("failed to parse"))
			}

			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*model.Claims); ok && token.Valid {
			c.Set("email", claims.Email)
		} else {
			c.JSON(http.StatusUnauthorized, model.NewErrorResponse("unauthorized"))

			c.Abort()
			return
		}

		c.Next()
	})
}
