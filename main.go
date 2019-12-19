package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gomall/datasource"
	"gomall/pkg/logging"
	"gomall/pkg/settings"
	"gomall/pkg/token"
	"gomall/routers"
	"log"
	"net/http"
)

func init() {
	settings.Setup()
	datasource.Setup()
	logging.Setup()
	token.Setup()
}

func main() {
	gin.SetMode(settings.AppConfig.Server.RunMode)
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", settings.AppConfig.Server.HTTPPort),
		Handler:        routers.Init(),
		ReadTimeout:    settings.AppConfig.Server.ReadTimeout,
		WriteTimeout:   settings.AppConfig.Server.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("[info] start http server listening: %s", server.Addr)
	server.ListenAndServe()
}
