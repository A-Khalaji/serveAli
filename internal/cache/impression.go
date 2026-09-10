package cache

import (
	"encoding/json"
	"fmt"
	"serveAli/internal/database"
	"time"
)

type Impression struct {
	AdID           uint64 `json:"ad_id"`
	ProgramID      uint64 `json:"program_id"`
	ZoneID         uint64 `json:"zone_id"`
	SiteID         uint64 `json:"site_id"`
	VisitorID      string `json:"visitor_id"`
	DestinationURL string `json:"destination_url"`
}

func SetImpression(impressionID string, impression Impression) error {
	key := fmt.Sprintf("impression:%s", impressionID)

	data, err := json.Marshal(impression)
	if err != nil {
		return err
	}

	return database.Redis.Set(
		database.Ctx,
		key,
		data,
		24*time.Hour,
	).Err()
}

func GetImpression(impressionID string) (*Impression, error) {
	key := fmt.Sprintf("impression:%s", impressionID)

	data, err := database.Redis.Get(
		database.Ctx,
		key,
	).Result()
	if err != nil {
		return nil, err
	}

	var impression Impression

	if err := json.Unmarshal([]byte(data), &impression); err != nil {
		return nil, err
	}

	return &impression, nil
}
