package main

import (
	"serveAli/internal/handler"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "serveAli/cmd/api/docs"
)

func Routes(router *gin.Engine) {

	router.GET("/serve/:zone_id", handler.ServeAd)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
