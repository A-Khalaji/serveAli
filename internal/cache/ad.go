package cache

import (
	"encoding/json"
	"fmt"
	"serveAli/internal/database"
	"serveAli/internal/models"
	"strconv"
)

func SetAd(ad models.Ad) error {
	key := fmt.Sprintf("ad:%d", ad.ID)

	metadata, err := json.Marshal(ad.Metadata)
	if err != nil {
		return err
	}

	data := map[string]string{
		"id":         strconv.FormatUint(uint64(ad.ID), 10),
		"user_id":    strconv.FormatUint(uint64(ad.UserID), 10),
		"program_id": strconv.FormatUint(uint64(ad.ProgramID), 10),
		"name":       ad.Name,
		"ad_type":    string(ad.AdType),
		"metadata":   string(metadata),
	}

	return database.Redis.HSet(
		database.Ctx,
		key,
		data,
	).Err()
}
