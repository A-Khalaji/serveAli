package main

import (
	"serveAli/internal/event"
	"serveAli/internal/handler"
	"serveAli/internal/middleware"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "serveAli/cmd/api/docs"
)

func Routes(router *gin.Engine, geoIP *event.GeoIP) {

	router.Use(middleware.VisitorCookie())

	router.GET("/serve/:zone_id", handler.ServeAd(geoIP))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
