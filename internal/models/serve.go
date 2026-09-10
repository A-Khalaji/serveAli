package models

type ServedAd struct {
	Ad           Ad     `json:"ad"`
	ImpressionID string `json:"impression_id"`
}