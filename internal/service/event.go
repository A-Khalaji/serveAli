package service

import (
	"time"

	"serveAli/internal/database"
	"serveAli/internal/models"
)

func CreateEvent(
	eventType string,
	adID uint64,
	programID uint64,
	zoneID uint64,
	siteID uint64,
	context models.EventContext,
) error {
	event := models.Event{
		Timestamp: time.Now(),
		EventType: eventType,

		AdID:       adID,
		ProgramID: programID,
		ZoneID:     zoneID,
		SiteID:     siteID,

		VisitorID: context.VisitorID,

		IP:         context.IP,
		Country:    context.Country,
		City:       context.City,

		DeviceType: context.DeviceType,
		OS:         context.OS,
		Browser:    context.Browser,
	}

	return database.ClickHouseDB.Create(&event).Error
}