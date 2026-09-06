package handler

import (
	"net/http"
	"strconv"

	"serveAli/internal/event"
	"serveAli/internal/models"
	"serveAli/internal/service"

	"github.com/gin-gonic/gin"
)

// ServeAd godoc
//
// @Summary      Serve an ad
// @Description  Returns a matching ad for the specified zone.
// @Description
// @Description  Filters can be provided multiple times.
// @Description  Example:
// @Description  /serve/7?filter=available_ads&filter=type:BANNER&filter=category:restaurant&filter=keyword:pizza
//
// @Tags         Serve
// @Produce      json
//
// @Param        zone_id  path   int       true  "Zone ID"
// @Param        filter   query  []string  false "Redis filter sets. Can be repeated."
//
// @Success      200      {object} models.Ad
// @Failure      400      {object} map[string]string
// @Failure      404      {object} map[string]string
//
// @Router       /serve/{zone_id} [get]
func ServeAd(geoIP *event.GeoIP) gin.HandlerFunc {
	return func(c *gin.Context) {
		zoneID, err := strconv.ParseUint(c.Param("zone_id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid zone id",
			})
			return
		}

		visitor, err := c.Cookie("visitor")
		if err != nil || visitor == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "visitor cookie is required",
			})
			return
		}

		filters := c.QueryArray("filter")

		ip := c.ClientIP()

		userAgent := c.GetHeader("User-Agent")
		uaInfo := event.ParseUserAgent(userAgent)

		location := geoIP.Lookup(ip)

		eventContext := models.EventContext{
			VisitorID:  visitor,
			IP:         ip,
			Country:    location.Country,
			City:       location.City,
			DeviceType: uaInfo.DeviceType,
			OS:         uaInfo.OS,
			Browser:    uaInfo.Browser,
		}

		ad, err := service.Serve(
			uint(zoneID),
			visitor,
			filters,
			eventContext,
		)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, ad)
	}
}