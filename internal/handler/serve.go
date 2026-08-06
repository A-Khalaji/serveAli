package handler

import (
	"net/http"
	"strconv"

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
// @Description  /serve/7?visitor=user123&filter=available_ads&filter=type:BANNER&filter=category:restaurant&filter=keyword:pizza
//
// @Tags         Serve
// @Produce      json
//
// @Param        zone_id  path   int       true  "Zone ID"
// @Param        visitor  query  string    true  "Visitor identifier"
// @Param        filter   query  []string  false "Redis filter sets. Can be repeated."
//
// @Success      200      {object} models.Ad
// @Failure      400      {object} map[string]string
// @Failure      404      {object} map[string]string
//
// @Router       /serve/{zone_id} [get]
func ServeAd(c *gin.Context) {
	zoneID, err := strconv.ParseUint(c.Param("zone_id"),10,64,)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid zone id",
		})
		return
	}

	visitor := c.Query("visitor")
	if visitor == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "visitor is required",
		})
		return
	}

	filters := c.QueryArray("filter")
	ad, err := service.Serve(uint(zoneID),visitor,filters)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ad)
}