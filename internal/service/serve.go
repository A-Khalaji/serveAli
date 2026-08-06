package service

import (
	"errors"

	"serveAli/internal/cache"
	"serveAli/internal/models"
)

func Serve(zoneID uint,visitor string,filters []string,) (*models.Ad, error) {
	zone, err := cache.GetZone(zoneID)
	if err != nil {
		return nil, err
	}

	if len(filters) == 0 {

		filters = []string{
			"available_ads",
			"type:" + string(zone.ZoneType),
			"category:" + zone.Metadata.Category,
			"keyword:" + zone.Metadata.Keyword,
		}
	}

	ids, err := cache.GetMatchingAds(visitor,filters)
	if err != nil {
		return nil, err
	}

	for _, id := range ids {

		ad, err := cache.GetAd(id)

		if err != nil {
			continue
		}

		if cache.SeenAd(visitor, ad.ID) {
			continue
		}

		err = cache.MarkAdSeen(visitor,ad.ID,)
		if err != nil {
			return nil, err
		}
		
		return ad, nil
	}
	return nil, errors.New("no matching ad found")
}
