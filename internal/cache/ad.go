package cache

import (
	"errors"
	"fmt"
	"strconv"

	"serveAli/internal/database"
	"serveAli/internal/models"

	"github.com/google/uuid"
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

	isActive, _ := strconv.ParseBool(data["is_active"])
	isVerified, _ := strconv.ParseBool(data["is_verified"])

	return &models.Ad{
		ID:        uint(adID),
		UserID:    uint(userID),
		ProgramID: uint(programID),
		Name:      data["name"],
		AdType:    models.AdType(data["ad_type"]),
		Metadata: models.AdMetaData{
			Category: data["category"],
			Keyword:  data["keyword"],
		},
		IsActive:   isActive,
		IsVerified: isVerified,
	}, nil
}

func GetMatchingAds(visitor string, filters []string) ([]uint, error) {

	tempKey := fmt.Sprintf(
		"temp:ads:%s",
		uuid.New().String(),
	)

	err := database.Redis.SInterStore(
		database.Ctx,
		tempKey,
		filters...,
	).Err()

	if err != nil {
		return nil, err
	}

	ids, err := database.Redis.SMembers(
		database.Ctx,
		tempKey,
	).Result()

	if err != nil {
		return nil, err
	}

	if err := database.Redis.Del(
		database.Ctx,
		tempKey,
	).Err(); err != nil {
		return nil, err
	}

	var result []uint
	for _, id := range ids {

		n, err := strconv.ParseUint(id, 10, 64)

		if err != nil {
			continue
		}

		result = append(result, uint(n))
	}

	return result, nil
}
