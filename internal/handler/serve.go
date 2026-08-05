package handler

import (
	"net/http"
	"serveAli/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ServeAd godoc
//
//	@Summary		Serve an ad
//	@Description	Returns the first matching ad for the specified zone.
//	@Tags			Serve
//	@Produce		json
//	@Param			zone_id	path		int	true	"Zone ID"
//	@Success		200		{object}	models.Ad
//	@Failure		400		{object}	map[string]string
//	@Failure		404		{object}	map[string]string
//	@Router			/serve/{zone_id} [get]
func ServeAd(c *gin.Context) {
	zoneID, err := strconv.ParseUint(c.Param("zone_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid zone id",
		})
		return
	}

	ad, err := service.Serve(uint(zoneID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ad)
}
