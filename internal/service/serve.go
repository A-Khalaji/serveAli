package service

import (
	"errors"
	"serveAli/internal/cache"
	"serveAli/internal/models"
)

func Serve(zoneID uint) (*models.Ad, error) {

	zone, err := cache.GetZone(zoneID)
	if err != nil {
		return nil, err
	}

	ads, err := cache.GetAllAds()
	if err != nil {
		return nil, err
	}

	for _, ad := range ads {

		if Match(zone, &ad) {
			return &ad, nil
		}

	}

	return nil, errors.New("no matching ad found")
}

func Match(zone *models.Zone, ad *models.Ad) bool {
	if !ad.IsActive || !ad.IsVerified {
		return false
	}

	if zone.Metadata.Category != ad.Metadata.Category {
		return false
	}

	if zone.Metadata.Keyword != ad.Metadata.Keyword {
		return false
	}

	if ad.AdType != models.AdType(zone.ZoneType) {
		return false
	}

	return true
}
