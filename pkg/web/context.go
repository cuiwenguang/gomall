package web

import (
	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
}

type HandlerFunc func(*Context)

func Handler(handler HandlerFunc) gin.HandlerFunc{
	return func(c *gin.Context) {
		context := new(Context)
		context.Context = c
		handler(context)
	}
}