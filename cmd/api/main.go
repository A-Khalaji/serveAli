package main

import (
	"flag"
	"serveAli/internal/database"

	"github.com/gin-gonic/gin"
)

//	@title			ServeAli API
//	@version		1.0
//	@description	Simple ad serving service.
//	@BasePath		/

func main() {
	flag.Parse()

	database.ConnectRedis()

	router := gin.Default()

	Routes(router)

	router.Run("localhost:8001")
}
