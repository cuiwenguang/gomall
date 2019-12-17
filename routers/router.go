package routers

import (
	"github.com/gin-gonic/gin"
	"gomall/pkg/web"
)

import "gomall/api"

func Init() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/account/register", web.Handler(api.Register))
	r.POST("/account/login", web.Handler(api.Login))
	return r
}
