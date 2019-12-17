package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gomall/pkg/e"
	"gomall/pkg/token"
)

// JWT  middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		tokenStr := c.Query("token")
		if tokenStr == "" {
			code = e.BAD_REQUEST
		} else {
			_, err := token.Parse(tokenStr)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = e.FORBIDDEN
				default:
					code = e.UNAUTHORIZED
				}
			}
		}

		if code != e.SUCCESS {
			c.JSON(e.UNAUTHORIZED, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}