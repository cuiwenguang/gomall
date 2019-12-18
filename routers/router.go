package routers

import (
	"github.com/gin-gonic/gin"
	"gomall/middleware"
	"gomall/pkg/web"
)

import "gomall/api"

func Init() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	/** cors middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1"},
		AllowMethods:     []string{"GET","POST", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://127.0.0.1"
		},
		MaxAge: 12 * time.Hour,
	}))
	*/
	r.POST("/account/register", web.Handler(api.Register))
	r.POST("/account/login", web.Handler(api.Login))

	v1 := r.Group("/api/v1")
	v1.Use(middleware.Authorization())
	{
		// 需要验证权限的路由
	}

	return r

}
