package models

import "time"

type Event struct {
	Timestamp  time.Time `json:"timestamp"`
	EventType  string    `json:"event_type"`

	AdID       uint64 `json:"ad_id"`
	ProgramID  uint64 `json:"program_id"`
	ZoneID     uint64 `json:"zone_id"`
	SiteID     uint64 `json:"site_id"`

	VisitorID  string `json:"visitor_id"`

	IP         string `json:"ip"`
	Country    string `json:"country"`
	City       string `json:"city"`

	DeviceType string `json:"device_type"`
	OS         string `json:"os"`
	Browser    string `json:"browser"`
}

const (
	EventTypeView  = "view"
	EventTypeClick = "click"
)

type EventContext struct {
	VisitorID  string
	IP         string
	Country    string
	City       string
	DeviceType string
	OS         string
	Browser    string
}