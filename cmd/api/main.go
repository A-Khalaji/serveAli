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
	database.ConnectClickHouse()

	router := gin.Default()

	Routes(router)

	router.Run("0.0.0.0:8001")
}
