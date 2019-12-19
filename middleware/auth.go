package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gomall/pkg/e"
	"gomall/pkg/token"
)

// Authorization  middleware
func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			code = e.BAD_REQUEST
		} else {
			// 如果需要根据用户角色来确定可以访问的资源,此处解析claims获取用户信息后来判断是否有权限访问
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
