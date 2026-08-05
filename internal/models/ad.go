package models

type Ad struct {
	ID        uint       `json:"id"`
	UserID    uint       `json:"user_id"`
	ProgramID uint       `json:"program_id"`
	Name   string `json:"name"`
	AdType AdType `json:"ad_type"`
	Metadata AdMetaData `json:"metadata"`
	IsActive   bool `json:"is_active"`
	IsVerified bool `json:"is_verified"`
}

type AdType string

const (
	AdTypeVideo  AdType = "VIDEO"
	AdTypeBanner AdType = "BANNER"
	AdTypeNative AdType = "NATIVE"
)

type AdMetaData struct {
	Category string `json:"category"`
	Keyword  string `json:"keyword"`
}