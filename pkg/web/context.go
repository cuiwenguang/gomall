package web

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"gomall/pkg/tenant"
)

type Context struct {
	*gin.Context
	db *gorm.DB
}

type HandlerFunc func(*Context)

func Handler(handler HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		context := new(Context)
		context.Context = c
		tid := tenant.ResolverOrigin(c)
		if db, exist := tenant.DBMaps[tid]; exist {
			context.db = db
		}
		handler(context)
	}
}
