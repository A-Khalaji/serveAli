package main

import (
	"flag"
	"log"

	"serveAli/internal/database"
	"serveAli/internal/event"

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

	geoIP, err := event.NewGeoIP("data/GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal("failed to initialize GeoIP:", err)
	}
	defer geoIP.Close()

	router := gin.Default()

	Routes(router, geoIP)

	router.Run("0.0.0.0:8001")
}
