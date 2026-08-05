package cache

import (
	"serveAli/internal/database"
	"serveAli/internal/models"
	"encoding/json"
	"fmt"
	"strconv"
	)

func SetZone(zone models.Zone) error {
	key := fmt.Sprintf("zone:%d", zone.ID)

	metadata, err := json.Marshal(zone.Metadata)
	if err != nil {
		return err
	}

	data := map[string]string{
		"id":         strconv.FormatUint(uint64(zone.ID), 10),
		"user_id":    strconv.FormatUint(uint64(zone.UserID), 10),
		"site_id":    strconv.FormatUint(uint64(zone.SiteID), 10),
		"name":       zone.Name,
		"zone_type":  string(zone.ZoneType),
		"identifier": zone.Identifier,
		"metadata":   string(metadata),
	}

	return database.Redis.HSet(
		database.Ctx,
		key,
		data,
	).Err()
}