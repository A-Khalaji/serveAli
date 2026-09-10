package service

import (
	"errors"

	"github.com/google/uuid"

	"serveAli/internal/cache"
	"serveAli/internal/models"
)

func Serve(zoneID uint, visitor string, filters []string, eventContext models.EventContext) (*models.ServedAd, error) {
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

	ids, err := cache.GetMatchingAds(visitor, filters)
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

		impressionID := uuid.New()

		err = cache.SetImpression(
		    impressionID.String(),
			cache.Impression{
				AdID:           uint64(ad.ID),
				ProgramID:      uint64(ad.ProgramID),
				ZoneID:         uint64(zone.ID),
				SiteID:         uint64(zone.SiteID),
				VisitorID:      visitor,
				DestinationURL: ad.DestinationURL,
			},
		)
		if err != nil {
			return nil, err
		}

		err = cache.MarkAdSeen(visitor, ad.ID)
		if err != nil {
			return nil, err
		}

		err = CreateEvent(
			models.EventTypeView,
			uint64(ad.ID),
			uint64(ad.ProgramID),
			uint64(zone.ID),
			uint64(zone.SiteID),
			eventContext,
			impressionID.String(),
		)
		if err != nil {
			return nil, err
		}

		return &models.ServedAd{
			Ad:           *ad,
			ImpressionID: impressionID.String(),
		}, nil
	}

	return nil, errors.New("no matching ad found")
}