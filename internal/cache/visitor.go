package cache

import (
	"context"
	"fmt"
	"serveAli/internal/database"
	"time"
)

func SeenAd(visitor string, adID uint) bool {
	key := fmt.Sprintf("visitor:%s:seen_ads", visitor)

	exists, err := database.Redis.SIsMember(
		context.Background(),
		key,
		adID,
	).Result()

	if err != nil {
		return false
	}
	return exists
}

func MarkAdSeen(visitor string, adID uint) error {
	key := fmt.Sprintf("visitor:%s:seen_ads", visitor)

	err := database.Redis.SAdd(
		context.Background(),
		key,
		adID,
	).Err()
	if err != nil {
		return err
	}

	return database.Redis.Expire(
		context.Background(),
		key,
		time.Hour,
	).Err()
}
