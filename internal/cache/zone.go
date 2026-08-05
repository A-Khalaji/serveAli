package cache

import (
	"encoding/json"
	"errors"
	"fmt"
	"serveAli/internal/database"
	"serveAli/internal/models"
	"strconv"
)

func GetZone(id uint) (*models.Zone, error) {
	key := fmt.Sprintf("zone:%d", id)

	data, err := database.Redis.HGetAll(database.Ctx,key,).Result()

	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, errors.New("zone not found")
	}

	zoneID, _ := strconv.ParseUint(data["id"], 10, 64)
	userID, _ := strconv.ParseUint(data["user_id"], 10, 64)
	siteID, _ := strconv.ParseUint(data["site_id"], 10, 64)

	var metadata models.ZoneMetaData
	err = json.Unmarshal(
		[]byte(data["metadata"]),
		&metadata,
	)

	if err != nil {
		return nil, err
	}

	isActive, _ := strconv.ParseBool(data["is_active"])
	isVerified, _ := strconv.ParseBool(data["is_verified"])

	return &models.Zone{
		ID:         uint(zoneID),
		UserID:     uint(userID),
		SiteID:     uint(siteID),
		Name:       data["name"],
		ZoneType:   models.ZoneType(data["zone_type"]),
		Identifier: data["identifier"],
		Metadata:   metadata,
		IsActive:   isActive,
		IsVerified: isVerified,
	}, nil
}
