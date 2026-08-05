package cache

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"serveAli/internal/database"
	"serveAli/internal/models"
)

func GetAd(id uint) (*models.Ad, error) {
	key := fmt.Sprintf("ad:%d", id)

	data, err := database.Redis.HGetAll(database.Ctx, key).Result()

	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, errors.New("ad not found")
	}

	adID, _ := strconv.ParseUint(data["id"], 10, 64)
	userID, _ := strconv.ParseUint(data["user_id"], 10, 64)
	programID, _ := strconv.ParseUint(data["program_id"], 10, 64)

	var metadata models.AdMetaData
	err = json.Unmarshal(
		[]byte(data["metadata"]),
		&metadata,
	)

	if err != nil {
		return nil, err
	}

	isActive, _ := strconv.ParseBool(data["is_active"])
	isVerified, _ := strconv.ParseBool(data["is_verified"])

	return &models.Ad{
		ID:         uint(adID),
		UserID:     uint(userID),
		ProgramID:  uint(programID),
		Name:       data["name"],
		AdType:     models.AdType(data["ad_type"]),
		Metadata:   metadata,
		IsActive:   isActive,
		IsVerified: isVerified,
	}, nil
}

func GetAllAds() ([]models.Ad, error) {
	var ads []models.Ad

	keys, err := database.Redis.Keys(
		database.Ctx,
		"ad:*",
	).Result()

	if err != nil {
		return nil, err
	}

	for _, key := range keys {
		id := strings.TrimPrefix(key, "ad:")

		adID, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			continue
		}

		ad, err := GetAd(uint(adID))
		if err != nil {
			continue
		}

		ads = append(ads, *ad)
	}

	return ads, nil
}
