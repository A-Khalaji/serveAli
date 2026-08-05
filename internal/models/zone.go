package models

type Zone struct {
	ID         uint         `json:"id"`
	UserID     uint         `json:"user_id"`
	SiteID     uint         `json:"site_id"`
	Name       string       `json:"name"`
	ZoneType   ZoneType     `json:"zone_type"`
	Identifier string       `json:"identifier"`
	Metadata   ZoneMetaData `json:"metadata"`
}

type ZoneType string

const (
	ZoneTypeVideo  ZoneType = "VIDEO"
	ZoneTypeBanner ZoneType = "BANNER"
	ZoneTypeNative ZoneType = "NATIVE"
)

type ZoneMetaData struct {
	Category string `json:"category"`
	Keyword  string `json:"keyword"`
}